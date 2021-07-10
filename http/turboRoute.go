package http

import (
	"net/http"
	"strings"
)

type TurboRoute struct {

	turboHandler http.Handler

	name string

	path string

	err error

	// supportedMethods : `,` separated methods can be registered for a single route i.e. "GET,POST,DELETE"
	supportedMethods string

	//registeredRoutes map[string]*TurboRoute

	scheme string
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

// StoreTurboRoutes :
func (turboEngine *TurboEngine) StoreTurboRoutes(path string) *TurboRoute {
	route := &TurboRoute{path: path}
	turboEngine.routes = append(turboEngine.routes, route)
	return route
}

func (turboRoute *TurboRoute) StoreTurboMethod(methods... string) *TurboRoute {
	methodString := strings.Join(methods, ",")
	turboRoute.supportedMethods = methodString
	return turboRoute
}