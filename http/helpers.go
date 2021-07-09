package http

import (
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
	w.Write([]byte("Endpoint Not Found : " + r.URL.Path + "\n"))
}

// endpointNotFoundHandler :
func endpointNotFoundHandler() http.Handler {
	return http.HandlerFunc(endpointNotFound)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Requested Method : " + r.Method + " not supported for Endpoint : " + r.URL.Path + "\n"))
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(methodNotAllowed)
}

func contains(supportedMethods string, method string) bool {
	supMethods := strings.Split(supportedMethods, ",")
	for _, val := range supMethods {
		if val == method {
			return true
		}
	}
	return false
}

