package http

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrNotFound = errors.New("requested route not found in the registered routes")
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
func (turboEngine *TurboEngine) RegisterTurboRoute(path string, f func(w http.ResponseWriter, r *http.Request)) *TurboRoute {
	log.Printf("Registering Route : %s\n", path)
	te := turboEngine.PreWork(path)
	return te.HandlerFunc(f)
}

// PreWork : serves as a function to perform the necessary prework onto the routes if required
// It serves as a middleware function which can be extended to multiple functionalities
func (turboEngine *TurboEngine) PreWork(path string) *TurboRoute {
	log.Printf("Performing Prework : %s\n", path)
	// Add registered routes to a central store
	te := turboEngine.AddPaths(path)
	// Add more functions in the prework as the need and purpose arises
	return te
}

// GetRoutes : returns the list of registered routes
func (turboEngine *TurboEngine) GetRoutes() []*TurboRoute{
	return turboEngine.routes
}


func (turboEngine *TurboEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if turboEngine.Match(r, &match) {
		log.Println("isMatch")
		handler = match.Handler
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

func (turboEngine *TurboEngine) Match(r *http.Request, match *MatchedTurboRoute) bool {
	for _, val := range turboEngine.routes {
		log.Printf("checking registered path : %s with incoming path : %s\n", val.path, r.URL.Path)
		if val.path == r.URL.Path {
			return true
		}
	}
	match.Err = ErrNotFound
	return false
}