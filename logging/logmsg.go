package logging

import (
	"bytes"
	"fmt"
	"go.appmanch.org/commons/textutils"
	"sync"
	"time"
)

var logMsgPool = &sync.Pool{
	New: func() interface{} {
		lm := &LogMessage{
			Content: &bytes.Buffer{},
			Buf:     &bytes.Buffer{},
		}
		lm.Content.Grow(1024)
		lm.Buf.Grow(1280)
		return lm
	},
}

// LogMessage struct.
type LogMessage struct {
	Time    time.Time     `json-codec:"timestamp"`
	FnName  string        `json-codec:"function,omitempty"`
	Line    int           `json-codec:"line,omitempty"`
	Content *bytes.Buffer `json-codec:"msg"`
	Sev     Severity      `json-codec:"sev"`
	Buf     *bytes.Buffer
	//SevBytes []byte
}

func getLogMessageF(sev Severity, f string, v ...interface{}) *LogMessage {
	msg := logMsgPool.Get().(*LogMessage)
	msg.Sev = sev
	msg.Time = time.Now()
	msg.FnName = textutils.EmptyStr
	msg.Line = 0
	_, _ = fmt.Fprintf(msg.Content, f, v...)
	return msg
}

func getLogMessage(sev Severity, v ...interface{}) *LogMessage {
	msg := logMsgPool.Get().(*LogMessage)
	msg.Sev = sev
	msg.Time = time.Now()
	msg.FnName = textutils.EmptyStr
	msg.Line = 0
	_, _ = fmt.Fprint(msg.Content, v...)
	return msg
}

func putLogMessage(logMsg *LogMessage) {
	logMsg.Content.Reset()
	logMsgPool.Put(logMsg)
}
