package turbo

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestFunction(t *testing.T) {
	router := RegisterTurboEngine()
	router.Get("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from turbo"))
	})
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
				fmt.Printf("got = %#v, want = %#v\n", got, tt.want)
				t.Errorf("RegisterTurboEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
