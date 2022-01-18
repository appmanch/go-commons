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

type Validator interface {
	Validate() (bool, []error)
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

func Get(contentType string) Codec {
	var readerWriter ReaderWriter
	switch contentType {
	case JSON:
		{
			readerWriter = JsonRW()
		}
	case XML:
		{
			readerWriter = XmlRW()
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
