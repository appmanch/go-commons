package turbo

import (
	"errors"
	"fmt"
	"go.appmanch.org/commons/textutils"
	"net/http"
	"strings"
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

// Router router stuct that holds the router configuration
type Router struct {
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

//NewRouter Creates a new Router
func NewRouter() *Router {
	return &Router{
		unManagedRouteHandler:    nil,
		unsupportedMethodHandler: nil,
		topLevelRoutes:           make(map[string]*Route),
	}
}

//Get route : Add a turbo handler for get method
func (r *Router) Get(path string, handler http.Handler) (*Route, error) {
	return r.Add(path, handler, GET)
}

//Post route : Add a turbo handler for post method
func (r *Router) Post(path string, handler http.Handler) (*Route, error) {
	return r.Add(path, handler, POST)
}

//Put route : Add a turbo handler for put method
func (r *Router) Put(path string, handler http.Handler) (*Route, error) {
	return r.Add(path, handler, PUT)
}

//Delete route : Add a turbo handler for delete method
func (r *Router) Delete(path string, handler http.Handler) (*Route, error) {
	return r.Add(path, handler, DELETE)
}

//Add route : Add a turbo handler for one or more HTTP methods.
func (r *Router) Add(path string, handler http.Handler, methods ...string) (*Route, error) {
	//Check if the methods provided are valid if not return error straight away
	for _, method := range methods {
		if _, contains := Methods[method]; !contains {
			return nil, fmt.Errorf("Invalid/Unsupported Http method  %s provided", method)
		}
	}
	//TODO add path check for any query variables specified.

	pathValue := strings.TrimSpace(path)

	pathValues := strings.Split(pathValue, PathSeparator)[1:]
	length := len(pathValues)
	var route *Route = nil
	if length > 0 && pathValues[0] != textutils.EmptyStr {
		isPathVar := false
		name := ""

		for i, pathValue := range pathValues {
			isPathVar = pathValue[0] == ':'
			if isPathVar {
				name = pathValue[1:]
			} else {
				name = pathValue
			}

			currentRoute := &Route{
				path:         name,
				isPathVar:    isPathVar,
				childVarName: "",
				handlers:     make(map[string]http.Handler),
				subRoutes:    make(map[string]*Route),
				queryParams:  make(map[string]*QueryParam),
			}

			if route == nil {
				if v, ok := r.topLevelRoutes[name]; ok {
					route = v
				} else {
					//No Parent present add the current route as route and continue
					if currentRoute.isPathVar {
						return nil, errors.New("the framework does not support path variables at root context")
					}
					r.topLevelRoutes[name] = currentRoute
					route = currentRoute
				}
			} else {
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
					currentRoute.handlers[method] = handler
				}
			}

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
			currentRoute.handlers[method] = handler
		}
		//Root route will not have any path value
		r.topLevelRoutes[""] = currentRoute

	}
	return route, nil

}

func (r *Route) DebugPrintRoute() {
	fmt.Println("path: ", r.path, ", isPathVar: ", r.isPathVar, ", childVarName: ", r.childVarName)
	for k, v := range r.subRoutes {
		fmt.Println("Printing Info of sub route ", k)
		v.DebugPrintRoute()
	}

}

func (r *Router) DebugPrint() {
	for k, v := range r.topLevelRoutes {
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
