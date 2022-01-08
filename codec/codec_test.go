package codec

import (
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
	c := Get()
	b, _ := c.JSONParser(m)
	fmt.Println(b)
}