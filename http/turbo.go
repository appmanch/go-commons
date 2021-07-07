package http

import (
	"net/http"
)

type TurboRouter struct {

	//NotFoundHandler http.Handler

	//MethodNotAllowedHandler http.Handler

	routes []*TurboRoute

	//registeredRoutes map[string]*TurboRoute

	//TurboRouteConfig

	operation string

}

// TurboRouteConfig configs that can be passed while creating the Turbo instance
type TurboRouteConfig struct {
	isPathUrlEncoded bool
	isRegex bool
}

// RegisterTurbo returns a new router instance
func RegisterTurbo() *TurboRouter {
	return &TurboRouter{
		operation:	"GET",
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
	tr := turboRouter.PreWork(path)
	return tr.HandlerFunc(f)
}

// PreWork serves as a middleware function which can be extended to multiple functionalities
func (turboRouter *TurboRouter) PreWork(path string) *TurboRoute {
	return turboRouter.AddPaths(path)
}

// GetRoutes provides a list of all the routes that have been registered
func (turboRouter *TurboRouter) GetRoutes() []*TurboRoute {
	// should be sent as a map object in deserialized form
	return turboRouter.routes
}

