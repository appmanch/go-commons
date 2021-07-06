package http

import "net/http"

type TurboRoute struct {
	turboHandler http.Handler

	name string

	err error

	namedRoutes map[string]*TurboRoute
}

type TurboRouter struct {

	NotFoundHandler http.Handler

	MethodNotAllowedHandler http.Handler

	routes []*TurboRoute

	namedRoutes map[string]*TurboRoute
}

// RegisterTurbo returns a new router instance
func RegisterTurbo() *TurboRouter {
	return &TurboRouter{namedRoutes: make(map[string]*TurboRoute)}
}

func (turbo *TurboRoute) Handler(handler http.Handler) *TurboRoute {
	if turbo.err == nil {
		turbo.turboHandler = handler
	}
	return turbo
}

func (turbo *TurboRoute) RegisterRoute(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	return turbo.Handler(http.HandlerFunc(f))
}