package logging

import (
	"bufio"
	"io"
	"os"
)

// ConsoleWriter struct
type ConsoleWriter struct {
	errorWriter, warnWriter, infoWriter, debugWriter, traceWriter io.Writer
}

// InitConfig ConsoleWriter
func (cw *ConsoleWriter) InitConfig(w *WriterConfig) {
	if w.Console.WriteErrToStdOut {
		cw.errorWriter = bufio.NewWriter(os.Stdout)
	} else {
		cw.errorWriter = bufio.NewWriter(os.Stderr)
	}
	if w.Console.WriteWarnToStdOut {
		cw.warnWriter = bufio.NewWriter(os.Stdout)
	} else {
		cw.warnWriter = bufio.NewWriter(os.Stderr)
	}

	cw.infoWriter = bufio.NewWriter(os.Stdout)
	cw.debugWriter = bufio.NewWriter(os.Stdout)
	cw.traceWriter = bufio.NewWriter(os.Stdout)

}

// DoLog consoleWriter
func (cw *ConsoleWriter) DoLog(logMsg *LogMessage) {
	var writer io.Writer

	switch logMsg.Sev {
	case Off:
		break
	case ErrLvl:
		writer = cw.errorWriter
	case WarnLvl:
		writer = cw.warnWriter
	case InfoLvl:
		writer = cw.infoWriter
	case DebugLvl:
		writer = cw.debugWriter
	case TraceLvl:
		writer = cw.traceWriter
	}

	if writer != nil {

		writeLogMsg(writer, logMsg)
	}
}

// Close close ConsoleWriter
func (cw *ConsoleWriter) Close() error {
	return nil
}
