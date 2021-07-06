package http

import "net/http"

type TurboRouter struct {

	NotFoundHandler http.Handler

	MethodNotAllowedHandler http.Handler

	routes []*TurboRoute

	namedRoutes map[string]*TurboRoute
}

// RegisterTurbo returns a new router instance
func RegisterTurbo() *TurboRouter {
	return &TurboRouter{
		namedRoutes:	make(map[string]*TurboRoute),
	}
}

// NewTurboRoute registers a blank route
func (turboRouter *TurboRouter) NewTurboRoute() *TurboRoute {
	route := &TurboRoute{}
	turboRouter.routes = append(turboRouter.routes, route)
	return route
}

// RegisterRoute helps in registering the endpoint paths for the API
func (turboRouter *TurboRouter) RegisterRoute(path string, f func(http.ResponseWriter, *http.Request)) *TurboRoute {
	return turboRouter.NewTurboRoute().HandlerFunc(f)
}