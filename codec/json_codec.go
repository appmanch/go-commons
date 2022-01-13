package codec

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type JsonCodec struct {
	//Json JsonCodec specific code will come here.

}

func NewJsonCodec(v interface{}) Codec {
	return BaseCodec{readerWriter: JsonRW(v)}
}

func JsonRW(v interface{}) *JsonCodec {
	//Case the defn here and return the codec
	// base codec's reader writer
	return &JsonCodec{}
}

func (c *JsonCodec) Write(v interface{}, w io.Writer) error {
	// marshal wrapper
	// if the validation is successful then use the core json-codec marshal to generate the json-codec from the struct and write it back to the buffer
	output, err := json.Marshal(v)
	if err != nil {
		// in case of error during marshaling
		return errors.New(fmt.Sprintf("marshal error: %d", err))
	}
	_, errW := w.Write(output)
	if errW != nil {
		return errW
	}
	return nil
}

func (c *JsonCodec) Read(r io.Reader, v interface{}) error {
	// unmarshal wrapper
	// read the data from reader and map it to the interface
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New(fmt.Sprintf("input error: %d", err))
	}
	if errU := json.Unmarshal(b, v); err != nil {
		return errors.New(fmt.Sprintf("unmarshal error: %d", errU))
	}
	return nil
}

// Commenting for now, to be used later for the info during caching, if required!!
/*func (c *JsonCodec) Load(v interface{}) error {
	// check if the input interface constraints are present in the cache or not
	// caching the struct constraints
	// map all the constraints
	// parse the constraints to save to the cache while the struct comes in
	// make sure the map is synchronised
	return errors.New("register is not implemented in base codec")
}*/
