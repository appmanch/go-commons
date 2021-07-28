package turbo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var router = New()

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
			want1: nil,
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