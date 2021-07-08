package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"path"
)

var (
	ErrNotFound = errors.New("requested route not found in the registered routes")
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
// there should be an option to initialize turbo with some configuration
func RegisterTurbo() *TurboRouter {
	log.Println("Registering Turbo")
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
	log.Printf("Registering Route : %s\n", path)
	tr := turboRouter.PreWork(path)
	return tr.HandlerFunc(f)
}

// PreWork serves as a middleware function which can be extended to multiple functionalities
func (turboRouter *TurboRouter) PreWork(path string) *TurboRoute {
	log.Printf("Performing Prework : %s\n", path)
	return turboRouter.AddPaths(path)
}

// GetRoutes provides a list of all the routes that have been registered
func (turboRouter *TurboRouter) GetRoutes() []*TurboRoute {
	// should be sent as a map object in deserialized form
	return turboRouter.routes
}

func (turboRouter *TurboRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("ServeHTTP : %s\n", path)
	if p := refinePath(path); p != path {
		log.Printf("refinedPath : %s\n", p)
		url := *r.URL
		url.Path = p
		p = url.String()
		log.Printf("urlString() : %s\n", p)
		w.Header().Set("Location", p)
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}

	var match MatchedTurboRoute
	var handler http.Handler

	if turboRouter.Match(r, &match) {
		log.Println("isMatch")
		handler = match.Handler
		r = requestWithVars(r, match.Vars)
		r = requestWithRoute(r, match.TurboRoute)
	}

	/*if handler == nil && match.MatchErr == ErrMethodsMismatch {
		handler = methodNotAllowedHandler()
	}*/

	if handler == nil {
		log.Printf("handler is nil")
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}

type MatchedTurboRoute struct {
	TurboRoute *TurboRoute
	Handler http.Handler
	Vars map[string]string
	Err error
}

type contextKey int

const (
	varsKey contextKey = iota
	routeKey
)

func (turboRouter *TurboRouter) Match(r *http.Request, match *MatchedTurboRoute) bool {
	for _, val := range turboRouter.routes {
		log.Printf("checking registered path : %s with incoming path : %s\n", val.path, r.URL.Path)
		if val.path == r.URL.Path {
			return true
		}
	}
	match.Err = ErrNotFound
	return false
}

func requestWithVars(r *http.Request, vars map[string]string) *http.Request {
	ctx := context.WithValue(r.Context(), varsKey, vars)
	return r.WithContext(ctx)
}

func requestWithRoute(r *http.Request, route *TurboRoute) *http.Request {
	ctx := context.WithValue(r.Context(), routeKey, route)
	return r.WithContext(ctx)
}

// refinePath
// Borrowed from the golang's net/http package
func refinePath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}