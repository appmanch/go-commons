package http

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestRoutesHandler(t *testing.T) {
	tests := []struct {
		name string
		want *TurboEngine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterTurboEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoutesHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction(t *testing.T) {
	router := RegisterTurboEngine()
	router.RegisterTurboRoute("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from turbo"))
	} ).StoreTurboMethod("get", "Post")
	router.RegisterTurboRoute("/api/v2/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from turbo"))
	} )
	router.RegisterTurboRoute("/api/v3/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from turbo"))
	} )

	log.Println(router.GetRoutes())

}