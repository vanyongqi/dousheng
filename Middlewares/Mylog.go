package Middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type MyHook struct{}

// 设置一个field
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["api"] = "yongqi"
	return nil
}

// 哪些等级的日志才会生效 过滤等级
func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	/*
		var levelColor int
		switch entry.Level {
		case logrus.DebugLevel, logrus.TraceLevel:
			levelColor = gray
		case logrus.WarnLevel:
			levelColor = yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = red
		default:
			levelColor = blue
		}
	*/

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%5s] ▶▶ %5s ▍ %5s %5s ▍ %5s\n", timestamp, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%5s] ▶▶ %5s ▍ %5s ▍\n", timestamp, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.SetFormatter(&LogFormatter{})
		// 日志的打开格式是追加，所以不能用os.Create
		f, _ := os.OpenFile("./logRecord.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		logrus.SetReportCaller(true) // 显示文件和代码行数
		logrus.SetOutput(io.MultiWriter(f, os.Stdout))
		//logrus.Info("test Info")
		//logrus.Warn("test Warning")
		////logrus.AddHook(&MyHook{})

		//log := logrus.New()
		//log.SetFormatter(&LogFormatter{})
		//log.SetReportCaller(true)
		//log.SetOutput(os.Stdout)
		//log.Info("Info")
		//log.Debug("Debug")
		//log.Warning("warning")
		//log.Error("error")

	}
}
