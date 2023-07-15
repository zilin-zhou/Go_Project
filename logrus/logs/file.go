package logs

import "os"

const LOGPATH = "runtime/logs/logrus.log"

type fileWriter struct {
	*os.File
}

func (f *fileWriter) Flush() {
	f.Sync()
}
func newFileWriter() logWriter {
	//打开文件
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		file = os.Stderr
	}
	return &fileWriter{
		file,
	}
}

// 初始化函数
func init() {
	RegisterInitWriterFunc("file", newFileWriter)
}
