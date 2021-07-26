package turbo

import (
	"errors"
	"fmt"
	"go.appmanch.org/commons/textutils"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	ErrNotFound    = errors.New("requested route not found in the registered routes")
	MethodNotFound = errors.New("request method not registered for the route")
)

//PathSeparator constant that holds the path separator for the URIs
const (
	PathSeparator = "/"
	GET           = "GET"
	HEAD          = "HEAD"
	POST          = "POST"
	PUT           = "PUT"
	DELETE        = "DELETE"
	OPTIONS       = "OPTIONS"
	TRACE         = "TRACE"
	PATCH         = "PATCH"
)

var Methods = map[string]string{
	GET:     GET,
	HEAD:    HEAD,
	POST:    POST,
	PUT:     PUT,
	DELETE:  DELETE,
	OPTIONS: OPTIONS,
	TRACE:   TRACE,
	PATCH:   PATCH,
}

var routers = make(map[string]*Router)
var lock = sync.RWMutex{}

// Router router struct that holds the router configuration
type Router struct {
	//group of the router
	group string
	//Handler for any route that is not defined
	unManagedRouteHandler http.Handler
	//Handler for any methods that are not supported
	unsupportedMethodHandler http.Handler
	//Routes Managed by this router
	topLevelRoutes map[string]*Route
}

//Route : TODO Documentation
type Route struct {
	Err error
	//name of the route fragment if this is a path variable the name of the variable will be used here.
	path string
	//Checks if this is a variable. only one path variable at this level will be supported.
	isPathVar bool
	//childVarName varName
	childVarName string
	//Handlers for HTTP Methods <method>|<Handler>
	handlers map[string]http.Handler
	//Sub Routes from this path
	subRoutes map[string]*Route
	//Query Parameters that may be used.
	queryParams map[string]*QueryParam
}

//QueryParam for the Route configuration
type QueryParam struct {
	//required flag
	required bool
	//name of the query parameter
	name string
	// TODO add mechanism for creating a typed query parameter to do auto type conversion in the framework.
}

// New : registers the new instance of the Turbo Framework
func New() *Router {
	return NewGroup(textutils.EmptyStr)
}

// NewGroup : registers the new instance of the Turbo Framework
func NewGroup(name string) *Router {
	lock.Lock()
	defer lock.Unlock()
	logger.Info("Registering New Turbo Instance")
	if v, ok := routers[name]; ok {
		return v
	} else {
		v := &Router{
			group:                    name,
			unManagedRouteHandler:    endpointNotFoundHandler(),
			unsupportedMethodHandler: methodNotAllowedHandler(),
			topLevelRoutes:           make(map[string]*Route),
		}
		routers[name] = v
		return v
	}
}

//Get route : Add a turbo handler for get method
func (router *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) (*Route, error) {
	return router.Add(path, f, GET)
}

//Post route : Add a turbo handler for post method
func (router *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) (*Route, error) {
	return router.Add(path, f, POST)
}

//Put route : Add a turbo handler for put method
func (router *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) (*Route, error) {
	return router.Add(path, f, PUT)
}

//Delete route : Add a turbo handler for delete method
func (router *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) (*Route, error) {
	return router.Add(path, f, DELETE)
}

//Add route : Add a turbo handler for one or more HTTP methods.
func (router *Router) Add(path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) (*Route, error) {
	lock.Lock()
	defer lock.Unlock()
	//Check if the methods provided are valid if not return error straight away
	for _, method := range methods {
		if _, contains := Methods[method]; !contains {
			return nil, fmt.Errorf("Invalid/Unsupported Http method  %s provided", method)
		}
	}
	//TODO add path check for any query variables specified.

	pathValue := strings.TrimSpace(path)

	pathValues := strings.Split(pathValue, PathSeparator)[1:]
	log.Println(pathValues)
	length := len(pathValues)
	var route *Route = nil
	if length > 0 && pathValues[0] != textutils.EmptyStr {
		isPathVar := false
		name := ""
		for i, pathValue := range pathValues {
			log.Println(pathValue)
			isPathVar = pathValue[0] == ':'
			log.Println(isPathVar)
			if isPathVar {
				name = pathValue[1:]
			} else {
				name = pathValue
			}
			log.Printf("Name: %s\n", name)
			currentRoute := &Route{
				path:         name,
				isPathVar:    isPathVar,
				childVarName: "",
				handlers:     make(map[string]http.Handler),
				subRoutes:    make(map[string]*Route),
				queryParams:  make(map[string]*QueryParam),
			}

			if route == nil {
				log.Println("Route is nil")
				if v, ok := router.topLevelRoutes[name]; ok {
					route = v
				} else {
					log.Println(currentRoute.isPathVar)
					//No Parent present add the current route as route and continue
					if currentRoute.isPathVar {
						return nil, errors.New("the framework does not support path variables at root context")
					}
					router.topLevelRoutes[name] = currentRoute
					route = currentRoute
				}
			} else {
				log.Println("Route is not nil")
				if v, ok := route.subRoutes[name]; ok {
					if v.isPathVar && isPathVar && v.path != name {
						return nil, errors.New("one path cannot have multiple names")
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
			log.Printf("CurrentRoute : %v\n", currentRoute)
			log.Printf("Handlers: %v\n", currentRoute.handlers)
		}
	} else {
		//TODO Handle the Root context path
		currentRoute := &Route{
			path:         "",
			isPathVar:    false,
			childVarName: "",
			handlers:     make(map[string]http.Handler),
			subRoutes:    make(map[string]*Route),
			queryParams:  make(map[string]*QueryParam),
		}
		for _, method := range methods {
			currentRoute.handlers[method] = prepareHandler(method, http.HandlerFunc(f))
		}
		//Root route will not have any path value
		router.topLevelRoutes[""] = currentRoute
	}
	log.Printf("Registered Route: %v\n", route)
	return route, nil
}

//Any default features like logging, auth etc will be injected here
func prepareHandler(method string, handler http.Handler) http.Handler {
	return handler
}

func (r *Route) DebugPrintRoute() {
	fmt.Println("path: ", r.path, ", isPathVar: ", r.isPathVar, ", childVarName: ", r.childVarName)
	for k, v := range r.subRoutes {
		fmt.Println("Printing Info of sub route ", k)
		v.DebugPrintRoute()
	}

}

func (router *Router) DebugPrint() {
	for k, v := range router.topLevelRoutes {
		fmt.Println("Printing Info of Top route ", k)
		v.DebugPrintRoute()
	}
}

func (r *Route) addQueryVar(name string, required bool) *Route {
	//TODO add name validation.
	queryParams := &QueryParam{
		required: required,
		name:     name,
	}
	//TODO Check if this name can be url encoded and save decoding per request,
	r.queryParams[name] = queryParams
	return r
}

//
//func (turboEngine *TurboEngine) group(path string) *TurboRoute {
//	turboRoute := turboEngine.StoreTurboRoutes(path, "")
//	return turboRoute
//}
//
//// GetRoutes : returns the list of registered routes
//func (r *Route) GetRoutes() []*TurboRoute {
//	return r.routes
//}

// ServeHTTP :
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var match Route
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
			return
		}
		return
	}

	// start by checking where the method of the Request is same as that of the registered method
	if router.Match(r, &match) {
		handler = match.handlers[r.Method]
	}

	if handler == nil && match.Err == MethodNotFound {
		handler = router.unsupportedMethodHandler
	} else if handler == nil && match.Err == ErrNotFound {
		handler = router.unManagedRouteHandler
	}

	if handler == nil {
		logger.InfoF("handler is nil")
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}


// Match : the function checks for the incoming request path whether it matches with the registered route's path or not
func (router *Router) Match(r *http.Request, route *Route) bool {

	return true
}

//// Match : the function checks for the incoming request path whether it matches with the registered route's path or not
//func (turboEngine *TurboEngine) Match(r *http.Request, match *MatchedTurboRoute) bool {
//	logger.InfoF("matchedRoutes %v", turboEngine.matchedRoutes)
//	logger.Info(r.URL.Path)
//	endpoints := strings.Split(r.URL.Path, "/")
//	logger.Info(len(endpoints))
//	returnFlag := false
//	url := ""
//	for i := 1; i < len(endpoints); i++ {
//		url = url + "/" + endpoints[i]
//		logger.InfoF("endpoint arr : %s", endpoints[i])
//		logger.InfoF("url: %s", url)
//		route, isMatch := turboEngine.matchedRoutes[url]
//		if isMatch {
//			// add a check to check further subroutes, logic to be implemented
//			logger.Info(isMatch)
//			logger.Info(route.isSubRoutePresent)
//			if route.isSubRoutePresent {
//				logger.Info(url)
//				subRoutePath := "/" + strings.Join(endpoints[i+1:], "/")
//				logger.Info(subRoutePath)
//				subRoute, isSubMatch := route.matchedSubRoute[subRoutePath]
//				logger.InfoF("issubmatch: %t", isSubMatch)
//				if isSubMatch {
//					if subRoute.routeMethod == r.Method {
//						match.Handler = subRoute.turboHandler
//						return true
//					} else {
//						match.Err = MethodNotFound
//						return false
//					}
//				} else {
//					match.Err = ErrNotFound
//					returnFlag = false
//				}
//			} else {
//				if route.routeMethod == r.Method {
//					match.Handler = route.turboHandler
//					return true
//				} else {
//					match.Err = MethodNotFound
//					return false
//				}
//			}
//		} else {
//			match.Err = ErrNotFound
//			returnFlag = false
//		}
//	}
//	match.Err = ErrNotFound
//	return returnFlag
//}
