package http

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrNotFound = errors.New("requested route not found in the registered routes")
	MethodNotFound = errors.New("request method not registered for the route")
)

// TurboEngine : code of the http framework
type TurboEngine struct {
	routes []*TurboRoute
	operation string
	isRegex bool
	isPathUrlEncoded bool
}

// RegisterTurboEngine : registers the new instance of the Turbo Framework
func RegisterTurboEngine() *TurboEngine {
	log.Println("Registering Turbo")
	return &TurboEngine{}
}

// RegisterTurboRoute : registers the new route in the HTTP Server for the API
func (turboEngine *TurboEngine) RegisterTurboRoute(methods string, path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	log.Printf("Registering Route : %s\n", path)
	te := turboEngine.PreWork(methods, path)
	return te.HandlerFunc(f)
}

// PreWork : serves as a function to perform the necessary prework onto the routes if required
// It serves as a middleware function which can be extended to multiple functionalities
func (turboEngine *TurboEngine) PreWork(methods string, path string) *TurboRoute {
	log.Printf("Performing Prework : %s\n", path)
	// Add registered routes to a central store
	te := turboEngine.StoreTurboRoutes(methods, path)
	// Add more functions in the prework as the need and purpose arises
	return te
}

// GetRoutes : returns the list of registered routes
func (turboEngine *TurboEngine) GetRoutes() []*TurboRoute{
	return turboEngine.routes
}

// ServeHTTP :
func (turboEngine *TurboEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var match MatchedTurboRoute
	var handler http.Handler
	method := r.Method

	// perform the path checks before, set the 301 status even before further computation
	if p := refinePath(path); p != path {
		url := *r.URL
		url.Path = p
		p = url.String()
		w.Header().Set("Location", p)
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write([]byte("Path Moved : " + p + "\n"))
		return
	}

	// start by checking where the method of the Request is same as that of the registered method
	if turboEngine.checkForMethod(method, &match) {
		log.Printf("ServeHTTP : %s\n", path)

		/*
			possible middle check
			Cautious handlers should read the Request.Body first, and then reply.
		*/

		if turboEngine.Match(r, &match) {
			handler = match.Handler
		}
	}

	if handler == nil && match.Err == MethodNotFound {
		handler = methodNotAllowedHandler()
	}

	if handler == nil && match.Err == ErrNotFound {
		handler = endpointNotFoundHandler()
	}

	if handler == nil {
		log.Printf("handler is nil")
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}

// MatchedTurboRoute :
type MatchedTurboRoute struct {
	TurboRoute *TurboRoute
	Handler http.Handler
	Vars map[string]string
	Err error
}

// Match :
func (turboEngine *TurboEngine) Match(r *http.Request, match *MatchedTurboRoute) bool {
	for _, val := range turboEngine.routes {
		log.Printf("checking registered path : %s with incoming path : %s\n", val.path, r.URL.Path)
		if val.path == r.URL.Path {
			match.Handler = val.turboHandler
			return true
		}
	}
	match.Err = ErrNotFound
	return false
}

//checkForMethod :
func (turboEngine *TurboEngine) checkForMethod(method string, match *MatchedTurboRoute) bool {
	for _, val := range turboEngine.routes {
		if contains(val.supportedMethods, method) {
			return true
		}
	}
	match.Err = MethodNotFound
	return false
}