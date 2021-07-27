package turbo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
)

// BenchmarkFindRouteStatic: Static Path Test
func BenchmarkFindRouteStatic(b *testing.B) {
	router.Get("/api/v1/health", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]byte("hello from turbo"))
	})
	testUrl, _ := url.Parse("/api/v1/health")
	req:= &http.Request{
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
	}
	for i := 0; i < b.N; i++ {
		router.findRoute(req)
	}
}

// BenchmarkFindRoutePathParam: Static Path Test
func BenchmarkFindRoutePathParam(b *testing.B) {
	router.Get("/api/v1/health/:id", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]byte("hello from turbo"))
	})
	testUrl, _ := url.Parse("/api/v1/health/123")
	req:= &http.Request{
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
	}
	for i := 0; i < b.N; i++ {
		router.findRoute(req)
	}
}