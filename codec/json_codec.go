package codec

import (
	"errors"
	"io"
)

type JsonCodec struct {
	//Json JsonCodec specific code will come here.

}

func NewJsonCodec(v interface{}) Codec {
	return BaseCodec{readerWriter: JsonRW(v)}
}

func JsonRW(v interface{}) *JsonCodec {
	//Case the defn here and return the codec

	return &JsonCodec{}
}

func (c *JsonCodec) Write(v interface{}, w io.Writer) error {
	return nil //TODO Implement
}

func (c *JsonCodec) Read(r io.Reader, v interface{}) error {
	return nil //TODO Implement

}

func (c *JsonCodec) Load(v interface{}) error {
	return errors.New("register is not implemented in base codec")
}
