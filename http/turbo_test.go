package http

import (
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
		want *TurboRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterTurbo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoutesHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction(t *testing.T) {
	router := RegisterTurbo()
	router.
}