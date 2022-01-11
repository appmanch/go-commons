package codec

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Message struct {
	Name string `json-codec:"name" constraints:"min-length:3"`
	Body string `json-codec:"body" constraints:"min-length:10"`
	Time int64  `json-codec:"time" constraints:"min-length:0"`
}

func TestInBuiltJson(t *testing.T) {
	m := Message{"Test", "Hello", 123124124}
	output, _ := json.Marshal(m)
	fmt.Println("output\n", string(output))
	//c := Get()
	//b, _ := c.JSONParser(m)
	//fmt.Println(b)
}

func TestGet(t *testing.T) {
	m := Message{"Test", "Hello", 123124124}
	codec := Get("application/json", m)
	codec.Write(m, nil)
}
