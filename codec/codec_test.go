package codec

import (
	"bytes"
	"testing"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

type XMLMessage struct {
	Name string `xml:"name"`
	Body string `xml:"body"`
	Time int64  `xml:"time"`
}

func TestNewJsonCodec(t *testing.T) {
	m := Message{"Test", "Hello", 123124124}
	c := Get("application/json")
	buf := new(bytes.Buffer)
	if err := c.Write(m, buf); err != nil {
		t.Errorf("error in write: %d", err)
	}

	const want = `{"name":"Test","body":"Hello","time":123124124}`
	if got := buf; got.String() != want {
		t.Errorf("got %q, want %q", got.String(), want)
	}
}

func TestNewXmlCodec(t *testing.T) {
	m := XMLMessage{"Test", "Hello", 123124124}
	c := Get("text/xml")
	buf := new(bytes.Buffer)
	if err := c.Write(m, buf); err != nil {
		t.Errorf("error in write: %d", err)
	}
	// fmt.Println(buf)
	const want = `<XMLMessage><name>Test</name><body>Hello</body><time>123124124</time></XMLMessage>`
	if got := buf; got.String() != want {
		t.Errorf("got %q, want %q", got.String(), want)
	}
}
