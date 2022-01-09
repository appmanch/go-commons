package json

import (
	"bytes"
	"reflect"
)

func Serialize(v interface{}) ([]byte, error) {
	s := &serializedState{}
	err := s.serialize(v)
	if err != nil {
		return nil, err
	}
	s.reflectValue(reflect.ValueOf(v))
	buf := append([]byte(nil), s.Bytes()...)
	return buf, nil
}

func (s *serializedState) serialize(v interface{}) (err error) {
	s.reflectValue(reflect.ValueOf(v))
	return nil
}

func (s *serializedState) reflectValue(v reflect.Value) {
	valueSerializer(v)
}

type serializedState struct {
	bytes.Buffer
	scratch  [64]byte
	ptrLevel uint
	ptrSeen  map[interface{}]struct{}
}

type serializerFunc func(e *serializedState, v reflect.Value)

func valueSerializer(v reflect.Value) serializerFunc {

	switch v.Type().Kind() {
	case reflect.Bool:
		return boolSerializer
	case reflect.Int:
		return intSerializer
	case reflect.Struct:
		return structSerializer(v.Type())
	case reflect.String:
		return stringSerializer
	default:
		return unsupportedSerializer
	}
}

func boolSerializer(e *serializedState, v reflect.Value) {}

func intSerializer(e *serializedState, v reflect.Value) {}

type field struct {
	name       string
	typ        reflect.Type
	serializer serializerFunc
}

type structSerializerStruct struct {
	fields structFields
}

type structFields struct {
	list []field
	idx  map[string]int
}

func structSerializer(t reflect.Type) serializerFunc {
	ss := structSerializerStruct{fields: fetchFields(t)}
	return ss.serialize
}

func (ss structSerializerStruct) serialize(e *serializedState, v reflect.Value) {
	// TODO
}

func stringSerializer(e *serializedState, v reflect.Value) {}

func unsupportedSerializer(e *serializedState, v reflect.Value) {}

// fetches the fields from the input struct and returns the fields mapped to the structFields
func fetchFields(t reflect.Type) structFields {

	// TODO

	return structFields{}
}
