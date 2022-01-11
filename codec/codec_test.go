package codec

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Message struct {
	Name string `json-codec:"name" contraints:"min-length:3"`
	Body string `json-codec:"body" contraints:"min-length:10"`
	Time int64  `json-codec:"time" contraints:"min-length:0"`
}

func TestEncoder(t *testing.T) {
	m := Message{"Test", "Hello", 123124124}
	output, _ := json.Marshal(m)
	fmt.Println("output\n", string(output))
	//c := Get()
	//b, _ := c.JSONParser(m)
	//fmt.Println(b)
}
