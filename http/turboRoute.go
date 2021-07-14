package http

import (
	"net/http"
)

type TurboRoute struct {

	turboHandler http.Handler

	name string

	path string

	err error //not needed

	// supportedMethods : `,` separated methods can be registered for a single route i.e. "GET,POST,DELETE"
	//supportedMethods string

	routeMethod string

	//registeredRoutes map[string]*TurboRoute

	//scheme string // not needed
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
	turboEngine.routes = append(turboEngine.routes, route)
	return route
}

// TurboMethod : Function stores the respective supported methods required for the API
/*func (turboRoute *TurboRoute) TurboMethod(methods... string) *TurboRoute {
	methodString := strings.Join(methods, ",")
	turboRoute.supportedMethods = strings.ToUpper(methodString)
	return turboRoute
}*/