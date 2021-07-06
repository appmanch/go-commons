package http

import "net/http"

type TurboRoute struct {

	turboHandler http.Handler

	name string

	err error

	namedRoutes map[string]*TurboRoute
}

func (turboRoute *TurboRoute) Handler(handler http.Handler) *TurboRoute {
	if turboRoute.err == nil {
		turboRoute.turboHandler = handler
	}
	return turboRoute
}

func (turboRoute *TurboRoute) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *TurboRoute {
	return turboRoute.Handler(http.HandlerFunc(f))
}
