package config

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestNewProperties(t *testing.T) {
	tests := []struct {
		name string
		want *Properties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProperties(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_Get(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		d string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			if got := p.Get(tt.args.k, tt.args.d); got != tt.want {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_GetAsBool(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k          string
		defaultVal bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.GetAsBool(tt.args.k, tt.args.defaultVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_GetAsDecimal(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k          string
		defaultVal float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.GetAsDecimal(tt.args.k, tt.args.defaultVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsDecimal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_GetAsInt(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k          string
		defaultVal int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.GetAsInt(tt.args.k, tt.args.defaultVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_GetAsInt64(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k          string
		defaultVal int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.GetAsInt64(tt.args.k, tt.args.defaultVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAsInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_Put(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			if got := p.Put(tt.args.k, tt.args.v); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_PutBool(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		v bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.PutBool(tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_PutDecimal(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		v float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.PutDecimal(tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutDecimal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_PutInt(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.PutInt(tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_PutInt64(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		k string
		v int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			got, err := p.PutInt64(tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties_ReadFrom(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			if err := p.ReadFrom(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProperties_WriteTo(t *testing.T) {
	type fields struct {
		props         map[string]*value
		resolvedProps map[string]string
		RWMutex       sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Properties{
				props:         tt.fields.props,
				resolvedProps: tt.fields.resolvedProps,
				RWMutex:       tt.fields.RWMutex,
			}
			w := &bytes.Buffer{}
			err := p.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteTo() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
