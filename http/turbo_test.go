package http

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

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

// create mock server route tests

func TestRegisterTurboEngine(t *testing.T) {
	tests := []struct {
		name string
		want *TurboEngine
	}{
		// TODO: Add test cases.
		{
			name: "TurboEngine",
			want: RegisterTurboEngine(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterTurboEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterTurboEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurboEngine_RegisterTurboRoute(t *testing.T) {
	type fields struct {
		routes           []*TurboRoute
		operation        string
		isRegex          bool
		isPathUrlEncoded bool
	}
	type args struct {
		path string
		f    func(w http.ResponseWriter, r *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TurboRoute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			turboEngine := &TurboEngine{
				routes:           tt.fields.routes,
				operation:        tt.fields.operation,
				isRegex:          tt.fields.isRegex,
				isPathUrlEncoded: tt.fields.isPathUrlEncoded,
			}
			if got := turboEngine.RegisterTurboRoute(tt.args.path, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterTurboRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}