package codec

import (
	"bytes"
	"fmt"
	"reflect"
)

// JSON Module of the codec used to perform the JSON manipulation with the help of codec
// Encoding and Decoding of the JSON Data performed

// JSONParser : parses the incoming struct and converts it to the JSON Bytes
// Encoding of Struct to the JSON Bytes
/**
{
	"name": "test"
}
*/
func (d defaultCodec) JSONParser(v interface{}) ([]byte, error) {
	// placeholder logic
	// logic WIP
	buf := &bytes.Buffer{}
	fmt.Println(reflect.ValueOf(v).Type())
	reflectValue(reflect.ValueOf(v))
	e := d.Write(v, buf)
	if e == nil {
		return buf.Bytes(), e
	} else {
		return nil, e
	}
}

func reflectValue(v reflect.Value) {
	valueEncoder(v)
}

type serializedState struct {
	bytes.Buffer
	scratch  [64]byte
	ptrLevel uint
	ptrSeen  map[interface{}]struct{}
}

type serializerFunc func(e *serializedState, v reflect.Value)

func valueEncoder(v reflect.Value) serializerFunc {

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

}

func stringSerializer(e *serializedState, v reflect.Value) {}

func unsupportedSerializer(e *serializedState, v reflect.Value) {}

// fetches the fields from the input struct and returns the fields mapped to the structFields
func fetchFields(t reflect.Type) structFields {
	return structFields{}
}

// JSONMapper : maps the incoming JSON Bytes to the provided Struct
// Decoding of the JSON Bytes to the Struct
func (d defaultCodec) JSONMapper(data []byte, v interface{}) error {
	// placeholder logic
	// logic WIP
	r := bytes.NewReader(data)
	return d.Read(r, v)
}
