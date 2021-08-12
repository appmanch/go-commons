package turbo

import (
	"net/http"

	"go.appmanch.org/commons/logging"
	"go.appmanch.org/commons/turbo/auth"
)

//FilterFunc :
type FilterFunc func(http.Handler) http.Handler

// AddFilter : Making the Filter Chain in the order of filters being added
// if f1, f2, f3, finalHandler handlers are added to the filter chain then the order of execution remains
// f1 -> f2 -> f3 -> finalHandler
func (route *Route) AddFilter(filter ...FilterFunc) *Route {
	newFilters := make([]FilterFunc, 0, len(route.filters)+len(filter))
	newFilters = append(newFilters, route.filters...)
	newFilters = append(newFilters, filter...)
	route.filters = newFilters
	return route
}

//middle handler function to be defined by the dev explicitly
//r.AddAuthenticator() //pass the middleware handler
//r.SetLogger() //pass the logger reference

func (route *Route) AddAuthenticator(auth auth.Authenticator) *Route {
	route.authFilter = auth
	return route
}

func (route *Route) SetLogger(logger *logging.Logger) *Route {
	route.logger = logger
	return route
}
