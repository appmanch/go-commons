package http

import (
	"log"
	"net/http"
)

type TurboRoute struct {

	turboHandler http.Handler

	name string

	path string

	err error //not needed

	isSubRoutePresent bool

	// supportedMethods : `,` separated methods can be registered for a single route i.e. "GET,POST,DELETE"
	//supportedMethods string

	routeMethod string

	//registeredRoutes map[string]*TurboRoute

	subRoute []*TurboRoute

	//scheme string // not needed
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
	route := &TurboRoute{path: path, routeMethod: method}
	turboEngine.matchedRoutes[path] = route
	log.Printf("routeinfo :%v\n\n", turboEngine.matchedRoutes[path])
	//turboEngine.routes = append(turboEngine.routes, route)
	return route
}

// SubRoute : Initialize a blank SubRoute
func (turboRoute *TurboRoute) SubRoute() *TurboRoute {
	return &TurboRoute{
		isSubRoutePresent: true,
		matchedSubRoute: make(map[string]*TurboRoute),
	}
}

