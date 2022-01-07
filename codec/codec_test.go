package codec

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func TestEncoder(t *testing.T) {
	m := Message{"Test", "Hello", 123124124}
	b, _ := json.Marshal(m)
	fmt.Println(b)
}