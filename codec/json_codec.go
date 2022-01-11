package codec

import (
	"encoding/json"
	"errors"
	"io"

	"go.appmanch.org/commons/codec/json-codec"
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

	if err := json_codec.Validate(v); err != nil {
		// if the input struct is not validated against the `constraints` then we fail the encoding part here
		return err
	}
	// if the validation is successful then use the core json-codec marshal to generate the json-codec from the struct and write it back to the buffer
	_, err := json.Marshal(v)
	if err != nil {
		// in case of error during marshaling
		return nil
	}
	return nil //TODO Implement
}

func (c *JsonCodec) Read(r io.Reader, v interface{}) error {
	return nil //TODO Implement

}

func (c *JsonCodec) Load(v interface{}) error {
	return errors.New("register is not implemented in base codec")
}
