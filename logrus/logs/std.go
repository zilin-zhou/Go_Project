package logs

import "os"

// 实现接口
type stdWriter struct {
	*os.File
}

func (s *stdWriter) Flush() {
	s.Sync()
}
func newStdWriter() logWriter {
	return &stdWriter{
		os.Stderr,
	}
}

// 初始化方法
func init() {
	RegisterInitWriterFunc("std", newStdWriter)
}
