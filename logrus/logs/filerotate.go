package logs

//日志文件分割
import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

//实现接口

type fileRotateWriter struct {
	*rotatelogs.RotateLogs
}

func (frw *fileRotateWriter) Flush() {
	frw.Close()
}

func newFileRotateWriter() logWriter {
	write, err := getRotate()
	if err != nil {
		fmt.Println(err)
		return newStdWriter()
	}
	return &fileRotateWriter{
		write,
	}
}
func getRotate() (*rotatelogs.RotateLogs, error) {
	path := LOGPATH
	logf, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName("/path/to/access_log"),
		rotatelogs.WithMaxAge(time.Second*180),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	return logf, err
}

// 初始化
func init() {
	RegisterInitWriterFunc("fileRotate", newFileRotateWriter)
}
