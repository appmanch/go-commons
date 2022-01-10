package codec

import (
	"errors"
	"io"

	"go.appmanch.org/commons/codec/json"
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

	if err := json.Validate(v); err != nil {
		// if the input struct is not validated against the `constraints` then we fail the encoding part here
		return nil
	}
	// if the validation is successful then use the core json marshal to generate the json from the struct and write it back to the buffer

	return nil //TODO Implement
}

func (c *JsonCodec) Read(r io.Reader, v interface{}) error {
	return nil //TODO Implement

}

func (c *JsonCodec) Load(v interface{}) error {
	return errors.New("register is not implemented in base codec")
}
