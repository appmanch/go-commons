package turbo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

var router = NewRouter()

/*func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *Router
	}{
		// TODO: Add test cases.
		{
			name: "InitTest",
			want: router,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestRouter_findRoute(t *testing.T) {
	var tlr = make(map[string]*Route)
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		req *http.Request
	}
	route := &Route{
		path:         "abc",
		isPathVar:    false,
		childVarName: "",
		handlers:     make(map[string]http.Handler),
		subRoutes:    make(map[string]*Route),
		queryParams:  nil,
	}
	tlr["abc"]=route
	testUrl, _ := url.Parse("/abc")
	a := args{req: &http.Request{
		Method:           "",
		URL:              testUrl,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}}
	tests := []struct {
		name   string
		fields fields
		args args
		want   *Route
		want1  context.Context
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           tlr,
			},
			args: a,
			want: route,
			want1: context.Background(),
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			got, got1 := router.findRoute(tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("findRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRouter_GetPathParams(t *testing.T) {
	req := &http.Request{
		Method:           "",
		URL:              nil,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}
	type fields struct {
		lock                     sync.RWMutex
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		val
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			fields: fields{
				lock:                     sync.RWMutex{},
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "key",
				val: struct{
					valType string
				}{
					valType: "string",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				lock:                     tt.fields.lock,
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			ctx := context.WithValue(context.Background(), tt.args.id, tt.args.val)
			if got := router.GetPathParams(tt.args.id, tt.args.r.WithContext(ctx)); got != tt.want {
				t.Errorf("GetPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}