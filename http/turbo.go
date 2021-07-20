package http

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

var (
	ErrNotFound = errors.New("requested route not found in the registered routes")
	MethodNotFound = errors.New("request method not registered for the route")
)

// TurboEngine : core of the http framework
type TurboEngine struct {
	routes []*TurboRoute
	isRegex bool
	isPathUrlEncoded bool
	RouteNotFoundHandler http.Handler
	MethodNotAllowedHandler http.Handler
	matchedRoutes map[string]*TurboRoute
}

// RegisterTurboEngine : registers the new instance of the Turbo Framework
func RegisterTurboEngine() *TurboEngine {
	log.Println("Registering Turbo")
	return &TurboEngine{
		RouteNotFoundHandler: endpointNotFoundHandler(),
		MethodNotAllowedHandler: methodNotAllowedHandler(),
		matchedRoutes: make(map[string]*TurboRoute),
	}
}

// Get : Function exposed the GET Http Method
func (turboEngine *TurboEngine) Get(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	return turboEngine.RegisterTurboRoute(path, f, "GET")
}

// Post : Function exposed the POST Http Method
func (turboEngine *TurboEngine) Post(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	return turboEngine.RegisterTurboRoute(path, f, "POST")
}

// Put : Function exposed the PUT Http Method
func (turboEngine *TurboEngine) Put(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	return turboEngine.RegisterTurboRoute(path, f, "PUT")
}

// Delete : Function exposed the DELETE Http Method
func (turboEngine *TurboEngine) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	return turboEngine.RegisterTurboRoute(path, f, "DELETE")
}

// RegisterTurboRoute : registers the new route in the HTTP Server for the API
func (turboEngine *TurboEngine) RegisterTurboRoute(path string, f func(w http.ResponseWriter, r *http.Request), method string) *TurboRoute {
	log.Printf("Registering Route : %s\n", path)
	//turboEngine.rootPath = path
	turboRoute := turboEngine.PreWork(path, method)
	// register a default GET method for each route, further methods can be overwritten using the StoreTurboMethod
	//te = te.TurboMethod(method)
	return turboRoute.HandlerFunc(f)
}

// PreWork : serves as a function to perform the necessary prework onto the routes if required
// It serves as a middleware function which can be extended to multiple functionalities
func (turboEngine *TurboEngine) PreWork(path string, method string) *TurboRoute {
	log.Printf("Performing Prework : %s\n", path)
	// Add registered routes to a central store
	turboRoute := turboEngine.StoreTurboRoutes(path, method)
	//createMapping(path)
	// Add more functions in the prework as the need and purpose arises
	return turboRoute
}

func (turboEngine *TurboEngine) Group(path string) *TurboRoute {
	turboRoute := turboEngine.StoreTurboRoutes(path, "")
	return turboRoute
}

// GetRoutes : returns the list of registered routes
func (turboEngine *TurboEngine) GetRoutes() []*TurboRoute{
	return turboEngine.routes
}

// ServeHTTP :
func (turboEngine *TurboEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var match MatchedTurboRoute
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
	if turboEngine.Match(r, &match) {
		handler = match.Handler
	}

	if handler == nil && match.Err == MethodNotFound {
		handler = turboEngine.MethodNotAllowedHandler
	}

	if handler == nil && match.Err == ErrNotFound {
		handler = turboEngine.RouteNotFoundHandler
	}

	if handler == nil {
		log.Printf("handler is nil")
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}

// MatchedTurboRoute :
type MatchedTurboRoute struct {
	TurboRoute *TurboRoute
	Handler http.Handler
	Vars map[string]string
	Err error
}

// Match : the function checks for the incoming request path whether it matches with the registered route's path or not
func (turboEngine *TurboEngine) Match(r *http.Request, match *MatchedTurboRoute) bool {
	log.Printf("matchedRoutes %v\n\n", turboEngine.matchedRoutes)
	log.Println(r.URL.Path)
	endpoints := strings.Split(r.URL.Path, "/")
	log.Println(len(endpoints))
	returnFlag := false
	url := ""
	for i:= 1; i < len(endpoints); i++ {
		url = url + "/" + endpoints[i]
		log.Printf("endpoint arr : %s\n", endpoints[i])
		log.Printf("url: %s\n", url)
		route, isMatch := turboEngine.matchedRoutes[url]
		if isMatch {
			// add a check to check further subroutes, logic to be implemented
			log.Println(isMatch)
			log.Println(route.isSubRoutePresent)
			if route.isSubRoutePresent {
				log.Println(url)
				subRoutePath := "/" + strings.Join(endpoints[i+1:], "/")
				log.Println(subRoutePath)
				subRoute, isSubMatch := route.matchedSubRoute[subRoutePath]
				log.Printf("issubmatch: %t\n", isSubMatch)
				if isSubMatch {
					if subRoute.routeMethod == r.Method {
						match.Handler = subRoute.turboHandler
						returnFlag = true
						break
					} else {
						match.Err = MethodNotFound
						returnFlag = false
						break
					}
				} else {
					match.Err = ErrNotFound
					returnFlag = false
					break
				}
			} else {
				if route.routeMethod == r.Method {
					match.Handler = route.turboHandler
					returnFlag = true
					break
				} else {
					match.Err = MethodNotFound
					returnFlag = false
					break
				}
			}
		} else {
			match.Err = ErrNotFound
			returnFlag = false
		}
	}
	match.Err = ErrNotFound
	return returnFlag
}