package turbo

import "net/http"

//FilterFunc :
type FilterFunc func(http.Handler) http.Handler

//Filter :
func (filter FilterFunc) Filter(handler http.Handler) http.Handler {
	return filter(handler)
}

// AddFilter : Making the Filter Chain in the order of filters being added
// if f1, f2, f3, finalHandler handlers are added to the filter chain then the order of execution remains
// f1 -> f2 -> f3 -> finalHandler
func (route *Route) AddFilter(filter ...FilterFunc) {
	newMiddlewares := make([]FilterFunc, 0, len(route.middlewares)+len(filter))
	newMiddlewares = append(newMiddlewares, route.middlewares...)
	newMiddlewares = append(newMiddlewares, filter...)
	route.middlewares = newMiddlewares
}

//middle handler function to be defined by the dev explicitly
//r.AddAuthenticator() //pass the middleware handler
//r.SetLogger() //pass the logger reference

/*
Authenticator Filters
1. Basic Auth
more to be added
 */

func (route *Route) AddAuthenticator(auth http.Handler) {
	route.isAuthenticated = true

}

func (route *Route) SetLogger() {

}