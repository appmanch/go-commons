package turbo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"go.appmanch.org/commons/textutils"
)

// Router router struct that holds the router configuration
type Router struct {
	lock sync.RWMutex
	//Handler for any route that is not defined
	unManagedRouteHandler http.Handler
	//Handler for any methods that are not supported
	unsupportedMethodHandler http.Handler
	//Routes Managed by this router
	topLevelRoutes map[string]*Route
}

//Route : TODO Documentation
type Route struct {
	//name of the route fragment if this is a path variable the name of the variable will be used here.
	path      string
	paramType string
	//Checks if this is a variable. only one path variable at this level will be supported.
	isPathVar bool
	//childVarName varName
	childVarName string
	//handlers for HTTP Methods <method>|<Handler>
	handlers map[string]http.Handler
	//Sub Routes from this path
	subRoutes map[string]*Route
	//Query Parameters that may be used.
	queryParams map[string]*QueryParam
}

//QueryParam for the Route configuration
type QueryParam struct {
	//required flag : fail upfront if a required query param not present
	required bool
	//name of the query parameter
	name string
	// TODO add mechanism for creating a typed query parameter to do auto type conversion in the framework.
}

// NewRouter : registers the new instance of the Turbo Framework
func NewRouter() *Router {
	logger.InfoF("Initiating Turbo")
	return &Router{
		lock:                     sync.RWMutex{},
		unManagedRouteHandler:    endpointNotFoundHandler(),
		unsupportedMethodHandler: methodNotAllowedHandler(),
		topLevelRoutes:           make(map[string]*Route),
	}
}

//Get route : Add a turbo handler for GET method
func (router *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) *Route {
	return router.Add(path, f, GET)
}

//Post route : Add a turbo handler for POST method
func (router *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) *Route {
	return router.Add(path, f, POST)
}

//Put route : Add a turbo handler for PUT method
func (router *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) *Route {
	return router.Add(path, f, PUT)
}

//Delete route : Add a turbo handler for DELETE method
func (router *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *Route {
	return router.Add(path, f, DELETE)
}

//Add route : Add a turbo handler for one or more HTTP methods.
func (router *Router) Add(path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) *Route {
	router.lock.Lock()
	defer router.lock.Unlock()
	var route *Route = nil
	//Check if the methods provided are valid if not return error straight away
	for _, method := range methods {
		if _, contains := Methods[method]; !contains {
			panic(fmt.Sprintf("Invalid/Unsupported Http method  %s provided", method))
		}
	}
	logger.InfoF("Registering New Route: %s\n", path)
	log.Printf("Registering New Route: %s\n", path)
	//TODO add path check for any query variables specified.
	pathValue := strings.TrimSpace(path)
	pathValues := strings.Split(pathValue, PathSeparator)[1:]
	length := len(pathValues)
	if length > 0 && pathValues[0] != textutils.EmptyStr {
		isPathVar := false
		name := textutils.EmptyStr
		for i, pathValue := range pathValues {
			isPathVar = pathValue[0] == textutils.ColonChar
			if isPathVar {
				name = pathValue[1:]
			} else {
				name = pathValue
			}
			log.Println(name)
			currentRoute := &Route{
				path:         name,
				isPathVar:    isPathVar,
				childVarName: textutils.EmptyStr,
				handlers:     make(map[string]http.Handler),
				subRoutes:    make(map[string]*Route),
				queryParams:  make(map[string]*QueryParam),
			}
			if route == nil {
				if v, ok := router.topLevelRoutes[name]; ok {
					route = v
				} else {
					//No Parent present add the current route as route and continue
					if currentRoute.isPathVar {
						panic("the framework does not support path variables at root context")
					}
					router.topLevelRoutes[name] = currentRoute
					route = currentRoute
				}
			} else {
				if v, ok := route.subRoutes[name]; ok {
					if v.isPathVar && isPathVar && v.path != name {
						panic("one path cannot have multiple names")
					}
					route = v
				} else {
					route.subRoutes[name] = currentRoute
					if isPathVar {
						route.childVarName = name
					}
					route = currentRoute
				}
			}
			//At Last index add the method(s) to the map.
			if i == len(pathValues)-1 {
				for _, method := range methods {
					currentRoute.handlers[method] = http.HandlerFunc(f)
				}
			}
		}
	} else {
		//TODO Handle the Root context path
		currentRoute := &Route{
			path:         textutils.EmptyStr,
			isPathVar:    false,
			childVarName: textutils.EmptyStr,
			handlers:     make(map[string]http.Handler),
			subRoutes:    make(map[string]*Route),
			queryParams:  make(map[string]*QueryParam),
		}
		for _, method := range methods {
			currentRoute.handlers[method] = prepareHandler(method, http.HandlerFunc(f))
		}
		//Root route will not have any path value
		router.topLevelRoutes[textutils.EmptyStr] = currentRoute
	}
	return route
}

//Any default features like logging, auth etc will be injected here
func prepareHandler(method string, handler http.Handler) http.Handler {
	return handler
}

func (route *Route) DebugPrintRoute() {
	logger.InfoF("path: %s , isPathVar: %t , childVarName: %s", route.path, route.isPathVar, route.childVarName)
	for k, v := range route.subRoutes {
		logger.InfoF("Printing Info of sub route %s", k)
		v.DebugPrintRoute()
	}
}

func (router *Router) DebugPrint() {
	for k, v := range router.topLevelRoutes {
		logger.InfoF("Printing Info of Top route %s", k)
		v.DebugPrintRoute()
	}
}

func (route *Route) addQueryVar(name string, required bool) *Route {
	//TODO add name validation.
	queryParams := &QueryParam{
		required: required,
		name:     name,
	}
	//TODO Check if this name can be url encoded and save decoding per request,
	route.queryParams[name] = queryParams
	return route
}

// ServeHTTP :
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var handler http.Handler
	// perform the path checks before, set the 301 status even before further computation
	// these checks need not to be performed once the PreWork is refined and up to the mark
	if p := refinePath(path); p != path {
		url := *r.URL
		url.Path = p
		p = url.String()
		w.Header().Set("Location", p)
		w.WriteHeader(http.StatusMovedPermanently)
		_, err := w.Write([]byte("Path Moved : " + p + "\n"))
		if err != nil {
			logger.Error(err)
		}
		return
	}
	// start by checking where the method of the Request is same as that of the registered method
	match, ctx := router.findRoute(r)
	if match != nil {
		handler = match.handlers[r.Method]
	} else {
		handler = router.unManagedRouteHandler
	}
	if handler == nil {
		handler = router.unsupportedMethodHandler
	}
	handler.ServeHTTP(w, r.WithContext(ctx))
}

// findRoute : the function checks for the incoming request path whether it matches with the registered route's path or not
func (router *Router) findRoute(req *http.Request) (*Route, context.Context) {
	inReq := strings.Split(req.URL.Path, PathSeparator)[1:]
	var route *Route
	ctx := req.Context()
	for _, val := range inReq {
		if route == nil {
			route = router.topLevelRoutes[val]
			continue
		} else {
			if route.childVarName != textutils.EmptyStr {
				route = route.subRoutes[route.childVarName]
			} else {
				if r, ok := route.subRoutes[val]; ok {
					route = r
				} else {
					return nil, ctx
				}
			}
			if route.isPathVar {
				if val == "" {
					logger.ErrorF("Route Registered with a Path Param : %s", route.path)
					return nil, ctx
				}
				ctx = context.WithValue(ctx, route.path, val)
			}
		}
	}
	return route, ctx
}

func (router *Router) GetPathParams(id string, r *http.Request) string {
	val, ok := r.Context().Value(id).(string)
	if !ok {
		logger.ErrorF("Error Fetching Path Param %s", id)
	}
	return val
}

func (router *Router) GetIntPathParams(id string, r *http.Request) int {
	val, ok := r.Context().Value(id).(int)
	if !ok {
		logger.ErrorF("Error Fetching Path Param %s", id)
	}
	return val
}

func (router *Router) GetFloatPathParams(id string, r *http.Request) float64 {
	val, ok := r.Context().Value(id).(float64)
	if !ok {
		logger.ErrorF("Error Fetching Path Param %s", id)
	}
	return val
}

func (router *Router) GetBoolPathParams(id string, r *http.Request) bool {
	val, ok := r.Context().Value(id).(bool)
	if !ok {
		logger.ErrorF("Error Fetching Path Param %s", id)
	}
	return val
}

func (router *Router) GetQueryParams(id string, r *http.Request) string {
	val := r.URL.Query().Get(id)
	return val
}

func (router *Router) GetIntQueryParams(id string, r *http.Request) int {
	val, ok := strconv.Atoi(r.URL.Query().Get(id))
	if ok != nil {
		logger.ErrorF("Error Fetching Query Parameter %s", id)
	}
	return val
}

func (router *Router) GetFloatQueryParams(id string, r *http.Request) float64 {
	val, ok := strconv.ParseFloat(r.URL.Query().Get(id), 64)
	if ok != nil {
		logger.ErrorF("Error Fetching Query Parameter %s", id)
	}
	return val
}

func (router *Router) GetBoolQueryParams(id string, r *http.Request) bool {
	val, ok := strconv.ParseBool(r.URL.Query().Get(id))
	if ok != nil {
		logger.ErrorF("Error Fetching Query Parameter %s", id)
	}
	return val
}
