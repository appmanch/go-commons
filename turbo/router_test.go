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
		URL:              testUrl,
	}}
	tests := []struct {
		name   string
		fields fields
		args args
		want   *Route
		want1  context.Context
	}{
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
				val: 1231241,
				r:   req,
			},
			want: 1231241,
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
				t.Errorf("GetIntPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_GetFloatPathParams(t *testing.T) {
	req := &http.Request{}
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		val float64
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
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
				val: 21.34,
				r: req,
			},
			want: 21.34,
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
				val: 12.3124124123123,
				r:   req,
			},
			want: 12.3124124123123,
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
			if got := router.GetFloatPathParams(tt.args.id, tt.args.r.WithContext(ctx)); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetFloatPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_GetBoolPathParams(t *testing.T) {
	req := &http.Request{}
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		val bool
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
				val: true,
				r: req,
			},
			want: true,
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
				val: false,
				r:   req,
			},
			want: false,
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
			if got := router.GetBoolPathParams(tt.args.id, tt.args.r.WithContext(ctx)); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetBoolPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

var strUrl, _ = url.Parse("https://foo.com?test1=value1&test2=value2&test3=")

func TestRouter_GetQueryParams(t *testing.T) {
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
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
				id: "test1",
				r: &http.Request{URL: strUrl},
			},
			want: "value1",
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test2",
				r: &http.Request{URL: strUrl},
			},
			want: "value2",
		},
		{
			name: "Test3",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test3",
				r: &http.Request{URL: strUrl},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			if got := router.GetQueryParams(tt.args.id, tt.args.r); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetQueryParams() Type Got = %v, want %v", got, tt.want)
			}
			if got := router.GetQueryParams(tt.args.id, tt.args.r); got != tt.want {
				t.Errorf("GetQueryParams() Value Got = %v, want %v", got, tt.want)
			}
		})
	}
}

var intUrl, _ = url.Parse("https://foo.com?test1=1&test2=2")
var intUrlFail, _ = url.Parse("https://foo.com?test1=foo")

func TestRouter_GetIntQueryParams(t *testing.T) {
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
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
				id: "test1",
				r: &http.Request{URL: intUrl},
			},
			want: 1,
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test2",
				r: &http.Request{URL: intUrl},
			},
			want: 2,
		},
		{ // Failure Test Case
			name: "Test3",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test1",
				r: &http.Request{URL: intUrlFail},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			if got := router.GetIntQueryParams(tt.args.id, tt.args.r); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetIntQueryParams() = %v, want %v", got, tt.want)
			}
			if got := router.GetIntQueryParams(tt.args.id, tt.args.r); got != tt.want {
				t.Errorf("GetIntQueryParams() Value Got = %v, want %v", got, tt.want)
			}
		})
	}
}

var floatUrl, _ = url.Parse("https://foo.com?test1=1.1&test2=2.2332323")
var floatUrlFail, _ = url.Parse("https://foo.com?test1=hello")

func TestRouter_GetFloatQueryParams(t *testing.T) {
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Test1",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test1",
				r: &http.Request{URL: floatUrl},
			},
			want: 1.1,
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test2",
				r: &http.Request{URL: floatUrl},
			},
			want: 2.2332323,
		},
		{ // Failure Test Case
			name: "Test3",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test1",
				r: &http.Request{URL: floatUrlFail},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			if got := router.GetFloatQueryParams(tt.args.id, tt.args.r); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetFloatQueryParams() = %v, want %v", got, tt.want)
			}
			if got := router.GetFloatQueryParams(tt.args.id, tt.args.r); got != tt.want {
				t.Errorf("GetFloatQueryParams() Value Got = %v, want %v", got, tt.want)
			}
		})
	}
}

var boolUrl, _ = url.Parse("https://foo.com?test1=true&test2=false")
var boolUrlFail, _ = url.Parse("https://foo.com?test1=fail")

func TestRouter_GetBoolQueryParams(t *testing.T) {
	type fields struct {
		unManagedRouteHandler    http.Handler
		unsupportedMethodHandler http.Handler
		topLevelRoutes           map[string]*Route
	}
	type args struct {
		id string
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test1",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test1",
				r: &http.Request{URL: boolUrl},
			},
			want: true,
		},
		{
			name: "Test2",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test2",
				r: &http.Request{URL: boolUrl},
			},
			want: false,
		},
		{ // Failure Test Case
			name: "Test3",
			fields: fields{
				unManagedRouteHandler:    nil,
				unsupportedMethodHandler: nil,
				topLevelRoutes:           nil,
			},
			args: args{
				id: "test1",
				r: &http.Request{URL: boolUrlFail},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &Router{
				unManagedRouteHandler:    tt.fields.unManagedRouteHandler,
				unsupportedMethodHandler: tt.fields.unsupportedMethodHandler,
				topLevelRoutes:           tt.fields.topLevelRoutes,
			}
			if got := router.GetBoolQueryParams(tt.args.id, tt.args.r); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetBoolQueryParams() = %v, want %v", got, tt.want)
			}
			if got := router.GetBoolQueryParams(tt.args.id, tt.args.r); got != tt.want {
				t.Errorf("GetBoolQueryParams() Value Got = %v, want %v", got, tt.want)
			}
		})
	}
}