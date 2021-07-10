package http

import (
	"log"
	"net/http"
	"path"
	"strings"
)

// refinePath
// Borrowed from the golang's net/http package
func refinePath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	rp := path.Clean(p)
	if p[len(p)-1] == '/' && rp != "/" {
		rp += "/"
	}
	return rp
}

// endpointNotFound :
func endpointNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Endpoint Not Found : " + r.URL.Path + "\n"))
	if err != nil {
		return 
	}
}

// endpointNotFoundHandler : when a requested endpoint is not found in the registered route's this handler is invoked
func endpointNotFoundHandler() http.Handler {
	return http.HandlerFunc(endpointNotFound)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := w.Write([]byte("Requested Method : " + r.Method + " not supported for Endpoint : " + r.URL.Path + "\n"))
	if err != nil {
		return 
	}
}

// methodNotAllowedHandler : when a requested method is not allowed in the registered route's method list this handler is invoked
func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(methodNotAllowed)
}

// contains : checks if the requested method is present in the supported methods of the route
func contains(supportedMethods string, method string) bool {
	log.Println("contains call " + supportedMethods)
	supMethods := strings.Split(supportedMethods, ",")
	for _, val := range supMethods {
		if val == method {
			return true
		}
	}
	return false
}

