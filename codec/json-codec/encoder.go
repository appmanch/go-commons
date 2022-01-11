package json_codec

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

func Validate(v interface{}) error {

	value := reflect.ValueOf(v)
	typ := value.Type()

	var fields []field

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		tag := f.Tag.Get("constraints")
		constraints := parseTag(tag)
		fieldValue := value.Field(i)
		fmt.Println("fieldName", f.Name)
		fmt.Println("fieldValue", value.Field(i))
		fmt.Println("fieldType", f.Type.Kind())
		fmt.Println("tagNames", constraints)

		executeValidators(fieldValue, f.Type, constraints)

		/*if err := executeValidators(fieldValue, f.Type, constraints); err != nil {
			return err
		}*/
	}

	nameIdx := make(map[string]int, len(fields))
	for i, field := range fields {
		nameIdx[field.name] = i
	}

	return nil
}

func parseTag(tag string) map[string]string {
	m := make(map[string]string)
	split := strings.Split(tag, ",")
	for _, str := range split {
		constraintName := strings.Split(str, ":")[0]
		constraintValue := strings.Split(str, ":")[1]
		m[constraintName] = constraintValue
	}
	return m
}

type validatorFunc func(value reflect.Value, constraint map[string]string) error

func executeValidators(value reflect.Value, typ reflect.Type, constraint map[string]string) validatorFunc {
	switch typ.Kind() {
	case reflect.Bool:
		return boolValidator
	case reflect.String:
		return stringValidator(value, constraint)
	default:
		return invalidValidator
	}
}

func stringValidator(value reflect.Value, constraint map[string]string) error {

	// constraints to be predefined and mapped to a particular validation
	fmt.Println("executing validator", value)
	for i, val := range constraint {
		fmt.Println(i, val)
	}
	return nil
}

func boolValidator(value reflect.Value, constraint map[string]string) error {
	return nil
}

func invalidValidator(value reflect.Value, constraint map[string]string) error {
	return nil
}

// TODO
// To be implemented
var serializerCache sync.Map

func (s *serializedState) reflectValue(v reflect.Value) {
	checkSerializer(v)(s, v)
}

type serializedState struct {
	bytes.Buffer
	scratch  [64]byte
	ptrLevel uint
	ptrSeen  map[interface{}]struct{}
}

type serializerFunc func(e *serializedState, v reflect.Value)

func checkSerializer(v reflect.Value) serializerFunc {
	if !v.IsValid() {
		//return invalidEncoding
	}
	if fi, ok := serializerCache.Load(v.Type()); ok {
		return fi.(serializerFunc)
	}
	f := valueSerializer(v.Type())
	serializerCache.Store(v.Type(), f)
	return f
}

func valueSerializer(t reflect.Type) serializerFunc {
	switch t.Kind() {
	case reflect.Bool:
		return boolSerializer
	case reflect.Int:
		return intSerializer
	case reflect.Struct:
		return structSerializer(t)
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
	ss := structSerializerStruct{fields: loadFields(t)}
	return ss.serialize
}

func (ss structSerializerStruct) serialize(e *serializedState, v reflect.Value) {
	// TODO
	//next := byte('{')
	//loopFields:
	for i := range ss.fields.list {
		field := &ss.fields.list[i]
		fmt.Println("field", field.name)
		fieldValue := v.Field(i)
		fmt.Println("value", fieldValue)
	}

}

func stringSerializer(e *serializedState, v reflect.Value) {}

func unsupportedSerializer(e *serializedState, v reflect.Value) {}

// fetches the fields from the input struct and returns the fields mapped to the structFields
func loadFields(t reflect.Type) structFields {

	// TODO
	// loop through the fields

	var fields []field

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		field := field{
			name: f.Name,
			typ:  f.Type,
		}
		fields = append(fields, field)
	}

	nameIdx := make(map[string]int, len(fields))
	for i, field := range fields {
		nameIdx[field.name] = i
	}

	return structFields{fields, nameIdx}
}
