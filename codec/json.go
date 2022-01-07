package codec

import "bytes"

// JSON Module of the codec used to perform the JSON manipulation with the help of codec
// Encoding and Decoding of the JSON Data performed

// JSONParser : parses the incoming struct and converts it to the JSON Bytes
// Encoding of Struct to the JSON Bytes
func (d defaultCodec) JSONParser(v interface{}) ([]byte, error) {
	// placeholder logic
	// logic WIP
	buf := &bytes.Buffer{}
	e := d.Write(v, buf)
	if e == nil {
		return buf.Bytes(), e
	} else {
		return nil, e
	}
}

// JSONMapper : maps the incoming JSON Bytes to the provided Struct
// Decoding of the JSON Bytes to the Struct
func (d defaultCodec) JSONMapper(data []byte, v interface{}) error {
	// placeholder logic
	// logic WIP
	r := bytes.NewReader(data)
	return d.Read(r, v)
}

