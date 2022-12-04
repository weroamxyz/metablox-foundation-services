package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var logWriter io.Writer

type Config struct {
	Path        string `json:"path"`
	File        string `json:"file"`
	Level       string `json:"level"`
	Stdout      bool   `json:"stdout"`
	RotateTime  int    `json:"rotateTime"`
	RotateLimit uint   `json:"rotateLimit"`
}

func (c Config) ToLevel() logrus.Level {
	switch c.Level {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	default:
		return logrus.InfoLevel
	}
}

func Init(c *Config) error {

	logrus.SetReportCaller(true)
	logrus.SetLevel(c.ToLevel())
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           time.RFC3339Nano,
	})

	if c.Stdout {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(new(LogNullWriter))
	}

	writer, err := rotatelogs.New(
		c.Path+"/"+c.File+".%Y%m%d",
		rotatelogs.WithLinkName(c.Path+"/"+c.File),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(c.RotateTime)),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(c.RotateLimit),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	logrus.AddHook(lfsHook)

	if c.Stdout {
		logWriter = io.MultiWriter(os.Stdout, writer)
	} else {
		logWriter = writer
	}

	return nil
}

type LogNullWriter struct {
}

func (*LogNullWriter) Write(b []byte) (n int, err error) {
	return 0, nil
}

func GetLogWriter() io.Writer {
	return logWriter
}
