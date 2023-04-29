package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
)

var once sync.Once
var instance *logrus.Logger
func GetLogger() *logrus.Logger{
	once.Do(func(){
		instance = NewLogger()
	})
	
	return instance
}

type writeHook struct{
	Writers []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error{
	line, err := entry.String()
	if err != nil{
		return err
	}
	
	for _, w := range hook.Writers{
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	
	return err
}

func (hook *writeHook) Levels() []logrus.Level{
	return hook.LogLevels
}

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	})
	
	err := os.MkdirAll("logs", 0740)
	if err != nil{
		panic(err)
	}
	
	file, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil{
		panic(err)
	}
	
	l.SetOutput(io.Discard)
	
	l.AddHook(&writeHook{
		Writers: []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	
	l.SetLevel(logrus.TraceLevel)
		return l
}