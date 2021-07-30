package turbo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var router = NewRouter()

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *Router
	}{
		// TODO: Add test cases.
		{
			name: "InitTest",
			want: NewRouter(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	req := &http.Request{}
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		val string
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test1",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "key",
				val: "value",
				r: req,
			},
			want: "value",
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id:  "key2",
				val: "123",
				r:   req,
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			ctx := context.WithValue(context.Background(), tt.args.id, tt.args.val)
			if got := router.GetPathParams(tt.args.id, tt.args.r.WithContext(ctx)); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_GetIntPathParams(t *testing.T) {
	req := &http.Request{}
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		val int
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Test1",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "key",
				val: 2134,
				r: req,
			},
			want: 2134,
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id:  "key2",
				val: 123124124123123,
				r:   req,
			},
			want: 123124124123123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			ctx := context.WithValue(context.Background(), tt.args.id, tt.args.val)
			if got := router.GetIntPathParams(tt.args.id, tt.args.r.WithContext(ctx)); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}


