package http

import (
	"log"
	"net/http"
)

type TurboRoute struct {
	turboHandler http.Handler
	path string
	routeMethod string
	isSubRoutePresent bool
	err error //not needed
	turboEngine *TurboEngine
	matchedSubRoute map[string]*TurboRoute
}

// Handler :
func (turboRoute *TurboRoute) Handler(handler http.Handler) *TurboRoute {
	if turboRoute.err == nil {
		turboRoute.turboHandler = handler
	}
	return turboRoute
}

// HandlerFunc :
func (turboRoute *TurboRoute) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *TurboRoute {
	return turboRoute.Handler(http.HandlerFunc(f))
}

// StoreTurboRoutes : Function stores all the registered Routes
func (turboEngine *TurboEngine) StoreTurboRoutes(path string, method string) *TurboRoute {
	var route *TurboRoute
	if method == "" {
		route = &TurboRoute{
			path: path,
			turboEngine: turboEngine,
		}
	} else {
		route = &TurboRoute{path: path, routeMethod: method, turboEngine: turboEngine}
	}
	turboEngine.matchedRoutes[path] = route
	log.Printf("routeinfo :%v\n\n", turboEngine.matchedRoutes[path])
	//turboEngine.routes = append(turboEngine.routes, route)
	return route
}

func (turboRoute *TurboRoute) RegisterSubRoute(path string, method string) *TurboRoute {
	route := &TurboRoute{path: path, routeMethod: method}
	turboRoute.matchedSubRoute[path] = route
	return route
}

// SubRoute : Initialize a blank SubRoute
func (turboRoute *TurboRoute) SubRoute() *TurboRoute {
	turboRoute.matchedSubRoute = make(map[string]*TurboRoute)
	turboRoute.isSubRoutePresent = true
	return turboRoute
}

func (turboRoute *TurboRoute) Get(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	tr := turboRoute.RegisterSubRoute(path, "GET")
	return tr.HandlerFunc(f)
}

func (turboRoute *TurboRoute) Post(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	tr := turboRoute.RegisterSubRoute(path, "POST")
	return tr.HandlerFunc(f)
}

func (turboRoute *TurboRoute) Put(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	tr := turboRoute.RegisterSubRoute(path, "PUT")
	return tr.HandlerFunc(f)
}

func (turboRoute *TurboRoute) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	tr := turboRoute.RegisterSubRoute(path, "DELETE")
	return tr.HandlerFunc(f)
}
