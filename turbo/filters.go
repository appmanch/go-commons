package turbo

import "net/http"

//FilterFunc :
type FilterFunc func(http.Handler) http.Handler

// AddFilter : Making the Filter Chain in the order of filters being added
// if f1, f2, f3, finalHandler handlers are added to the filter chain then the order of execution remains
// f1 -> f2 -> f3 -> finalHandler
func (route *Route) AddFilter(filter ...FilterFunc) *Route {
	newMiddlewares := make([]FilterFunc, 0, len(route.middlewares)+len(filter))
	newMiddlewares = append(newMiddlewares, route.middlewares...)
	newMiddlewares = append(newMiddlewares, filter...)
	route.middlewares = newMiddlewares
	return route
}

//middle handler function to be defined by the dev explicitly
//r.AddAuthenticator() //pass the middleware handler
//r.SetLogger() //pass the logger reference

/*
Authenticator Filters
1. Basic Auth
more to be added
*/

func (route *Route) AddAuthenticator(auth FilterFunc) *Route {
	route.isAuthenticated = auth
	return route
}

func (route *Route) SetLogger() {

}
