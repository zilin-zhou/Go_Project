package logs

import "io"

var writerAdapter = make(map[string]InitLogWriteFunc, 0)

type InitLogWriteFunc func() logWriter

type logWriter interface {
	Flush()
	io.Writer
}

func RegisterInitWriterFunc(adapter string, writerFunc InitLogWriteFunc) {
	writerAdapter[adapter] = writerFunc
}
