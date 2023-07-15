package logs

import "github.com/sirupsen/logrus"

type Log struct {
	*logrus.Entry
	logWriter
}

func (l *Log) Flush() {
	l.logWriter.Flush()
}

type LogConf struct {
	Level       logrus.Level
	AdapterName string
}

// 初始化
func InitLog(conf LogConf) *Log {
	adapterName := "std"
	if conf.AdapterName != "" {
		adapterName = conf.AdapterName
	}
	write, ok := writerAdapter[adapterName]
	if !ok {
		adapterName = "std"
		write, _ = writerAdapter[adapterName]
	}

	log := &Log{
		logrus.NewEntry(logrus.New()),
		write(),
	}

	log.Logger.SetOutput(log.logWriter)
	if conf.Level != 0 {
		log.Logger.SetLevel(conf.Level)
	}
	//设置字段格式为json
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	log.Logger.SetReportCaller(true)
	return log
}
