package codec

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"go.appmanch.org/commons/codec/json"
	"go.appmanch.org/commons/textutils"
)

// StructMeta Struct
type StructMeta struct {
	Fields map[string]string
}

// FieldMeta struct captures the basic information for a field
type FieldMeta struct {
	// Name of the field
	Name string
	// Dimension holds the field dimension
	Dimension int
	// Required flag indicating if the field is a required field.
	Required bool
	// TargetNames stores map for known format types. This allows
	TargetNames map[string]string
	// TargetConfig stores configuration that is required by the target format . for Eg. Attribute config for XML etc.
	TargetConfig map[string]string
	// Sequence specifies the order  of the fields in the source/target format
	Sequence int
}

// StringFieldMeta Struct
type StringFieldMeta struct {
	FieldMeta
	DefaultVal string
	Pattern    string
	Format     string
	OmitEmpty  bool
	Length     int
}

// Int8FieldMeta Struct
type Int8FieldMeta struct {
	FieldMeta
	DefaultVal int8
	Min        int8
	Max        int8
}

// Int16FieldMeta Struct
type Int16FieldMeta struct {
	FieldMeta
	DefaultVal int16
	Min        int16
	Max        int16
}

// Int32FieldMeta Struct
type Int32FieldMeta struct {
	FieldMeta
	DefaultVal int16
	Min        int32
	Max        int32
}

// IntFieldMeta Struct
type IntFieldMeta struct {
	FieldMeta
	DefaultVal int
	Min        int
	Max        int
}

// UInt8FieldMeta Struct
type UInt8FieldMeta struct {
	FieldMeta
	DefaultVal uint8
	Min        uint8
	Max        uint8
}

// UInt16FieldMeta Struct
type UInt16FieldMeta struct {
	FieldMeta
	DefaultVal uint16
	Min        uint16
	Max        uint16
}

// UInt32FieldMeta Struct
type UInt32FieldMeta struct {
	FieldMeta
	DefaultVal uint32
	Min        uint32
	Max        uint32
}

// UIntFieldMeta Struct
type UIntFieldMeta struct {
	FieldMeta
	DefaultVal uint
	Min        uint
	Max        uint
}

// UInt64FieldMeta Struct
type UInt64FieldMeta struct {
	FieldMeta
	DefaultVal uint64
	Min        uint64
	Max        uint64
}

// Float32FieldMeta Struct
type Float32FieldMeta struct {
	FieldMeta
	DefaultVal float32
	Min        float32
	Max        float32
}

// Float64FieldMeta Struct
type Float64FieldMeta struct {
	FieldMeta
	DefaultVal float64
	Min        float64
	Max        float64
}

// BooleanFieldMeta Struct
type BooleanFieldMeta struct {
	FieldMeta
	DefaultVal bool
}

// StringEncoder Interface
type StringEncoder interface {
	//EncodeToString will encode a type to string
	EncodeToString(v interface{}) (string, error)
}

// BytesEncoder Interface
type BytesEncoder interface {
	// EncodeToBytes will encode the provided type to []byte
	EncodeToBytes(v interface{}) ([]byte, error)
}

// StringDecoder Interface
type StringDecoder interface {
	//DecodeString will decode  a type from string
	DecodeString(s string, v interface{}) error
}

// BytesDecoder Interface
type BytesDecoder interface {
	//DecodeBytes will decode a type from an array of bytes
	DecodeBytes(b []byte, v interface{}) error
}

type JSONEncoder interface {
	JSONParser(v interface{}) ([]byte, error)
}

type JSONDecoder interface {
	JSONMapper(data []byte, v interface{}) error
}

// Encoder Interface
type Encoder interface {
	StringEncoder
	JSONEncoder
}

// Decoder Interface
type Decoder interface {
	StringDecoder
	JSONDecoder
}

type ReaderWriter interface {
	//Write a type to writer
	Write(v interface{}, w io.Writer) error
	//Read a type from a reader
	Read(r io.Reader, v interface{}) error
}

// Codec Interface
type Codec interface {
	Decoder
	Encoder
	ReaderWriter
	Register(v interface{}) error
}

type defaultCodec struct {
}

func Get() Codec {
	return defaultCodec{}
}

func (d defaultCodec) DecodeString(s string, v interface{}) error {
	r := strings.NewReader(s)
	return d.Read(r, v)
}

// JSONMapper : maps the incoming JSON Bytes to the provided Struct
// Decoding of the JSON Bytes to the Struct
func (d defaultCodec) JSONMapper(data []byte, v interface{}) error {
	// placeholder logic
	// logic WIP
	r := bytes.NewReader(data)
	return d.Read(r, v)
}

// DecodeBytes : might not be needed as we would be moving the JSON operations to separate file
// ----->>> JSONMapper
/*func (d defaultCodec) DecodeBytes(b []byte, v interface{}) error {
	r := bytes.NewReader(b)
	return d.Read(r, v)
}*/

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
	fmt.Println(reflect.ValueOf(v).Type())
	j, err := json.Serialize(v)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// EncodeToBytes : might not be needed as we would be moving the JSON operations to separate file
// ----->>> JSONParser
/*func (d defaultCodec) EncodeToBytes(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	e := d.Write(v, buf)
	if e == nil {
		return buf.Bytes(), e
	} else {
		return nil, e
	}
}*/

func (d defaultCodec) EncodeToString(v interface{}) (string, error) {
	buf := &bytes.Buffer{}
	e := d.Write(v, buf)
	if e == nil {
		return buf.String(), e
	} else {
		return textutils.EmptyStr, e
	}
}

func (d defaultCodec) Read(r io.Reader, v interface{}) error {
	return errors.New("reader is not implemented in base codec")
}

func (d defaultCodec) Write(v interface{}, w io.Writer) error {

	return errors.New("writer is not implemented in base codec")
}

func (d defaultCodec) Register(v interface{}) error {
	return errors.New("register is not implemented in base codec")
}
