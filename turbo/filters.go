package turbo

import "net/http"

// AddFilter :
func (route *Route) AddFilter(filter  http.Handler) {
	route.middlewares = append(route.middlewares, filter)
}

/*
Authenticator Filters
1. Basic Auth
more to be added
 */

func (route *Route) AddAuthenticator(auth http.Handler) {

}

func (route *Route) SetLogger() {

}