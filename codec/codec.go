package codec

import (
	"bytes"
	"io"
	"strings"

	"go.appmanch.org/commons/textutils"
)

const (
	JSON = "application/json"
	XML  = "text/xml"
	YAML = "text/x-yaml"
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
	//SkipField indicates that if the value of the field is absent/nil then skip the field while writing to data
	//This is similar to omitempty
	SkipField bool
}

// StringFieldMeta Struct
type StringFieldMeta struct {
	FieldMeta
	DefaultVal *string
	Pattern    *string
	Format     *string
	MinLength  *int
	MaxLength  *int
}

// Int8FieldMeta Struct
type Int8FieldMeta struct {
	FieldMeta
	DefaultVal *int8
	Min        *int8
	Max        *int8
}

// Int16FieldMeta Struct
type Int16FieldMeta struct {
	FieldMeta
	DefaultVal *int16
	Min        *int16
	Max        *int16
}

// Int32FieldMeta Struct
type Int32FieldMeta struct {
	FieldMeta
	DefaultVal *int16
	Min        *int32
	Max        *int32
}

// IntFieldMeta Struct
type IntFieldMeta struct {
	FieldMeta
	DefaultVal *int
	Min        *int
	Max        *int
}

// UInt8FieldMeta Struct
type UInt8FieldMeta struct {
	FieldMeta
	DefaultVal *uint8
	Min        *uint8
	Max        *uint8
}

// UInt16FieldMeta Struct
type UInt16FieldMeta struct {
	FieldMeta
	DefaultVal *uint16
	Min        *uint16
	Max        *uint16
}

// UInt32FieldMeta Struct
type UInt32FieldMeta struct {
	FieldMeta
	DefaultVal *uint32
	Min        *uint32
	Max        *uint32
}

// UIntFieldMeta Struct
type UIntFieldMeta struct {
	FieldMeta
	DefaultVal *uint
	Min        *uint
	Max        *uint
}

// UInt64FieldMeta Struct
type UInt64FieldMeta struct {
	FieldMeta
	DefaultVal *uint64
	Min        *uint64
	Max        *uint64
}

// Float32FieldMeta Struct
type Float32FieldMeta struct {
	FieldMeta
	DefaultVal *float32
	Min        *float32
	Max        *float32
}

// Float64FieldMeta Struct
type Float64FieldMeta struct {
	FieldMeta
	DefaultVal *float64
	Min        *float64
	Max        *float64
}

// BooleanFieldMeta Struct
type BooleanFieldMeta struct {
	FieldMeta
	DefaultVal *bool
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

// Encoder Interface
type Encoder interface {
	StringEncoder
	BytesEncoder
}

// Decoder Interface
type Decoder interface {
	StringDecoder
	BytesDecoder
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
}

type BaseCodec struct {
	readerWriter ReaderWriter
}

func Get(contentType string, v interface{}) Codec {
	var readerWriter ReaderWriter
	switch contentType {
	case JSON:
		{
			readerWriter = JsonRW(v)
		}
	}

	return BaseCodec{
		readerWriter: readerWriter,
	}

}

func (bc BaseCodec) DecodeString(s string, v interface{}) error {
	r := strings.NewReader(s)
	return bc.Read(r, v)
}

func (bc BaseCodec) DecodeBytes(b []byte, v interface{}) error {
	r := bytes.NewReader(b)
	return bc.Read(r, v)
}

// EncodeToBytes :
func (bc BaseCodec) EncodeToBytes(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	e := bc.Write(v, buf)
	if e == nil {
		return buf.Bytes(), e
	} else {
		return nil, e
	}
}

func (bc BaseCodec) EncodeToString(v interface{}) (string, error) {
	buf := &bytes.Buffer{}
	e := bc.Write(v, buf)
	if e == nil {
		return buf.String(), e
	} else {
		return textutils.EmptyStr, e
	}
}

func (bc BaseCodec) Read(r io.Reader, v interface{}) error {
	return bc.readerWriter.Read(r, v)
}

func (bc BaseCodec) Write(v interface{}, w io.Writer) error {

	return bc.readerWriter.Write(v, w)
}
