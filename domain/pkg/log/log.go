package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	mw := io.MultiWriter(os.Stdout)
	Logger.SetOutput(mw)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&MyFormatter{})
}

type MyFormatter struct {
}

func CorlorHandler(msg string) string {
	msg = strings.ToUpper(msg)
	switch msg {
	case "DEBUG":
		return "\033[1;36m" + msg + "\033[0m"
	case "INFO":
		return "\033[1;32m" + msg + "\033[0m"
	case "WARN":
		return "\033[1;33m" + msg + "\033[0m"
	case "ERROR":
		return "\033[1;31m" + msg + "\033[0m"
	case "FATAL":
		return "\033[1;35m" + msg + "\033[0m"
	default:
		return msg
	}
}

// Format implement the Formatter interface
func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	var newLog string
	// entry.Message 就是需要打印的日志
	if entry.HasCaller() {
		level := CorlorHandler(entry.Level.String())
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] - [%s] - %s- [%v:%v %s]",
			entry.Time.Format("2006-01-02 15:04:05"),
			level,
			entry.Message, fName,
			entry.Caller.Line, entry.Caller.Function)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
	}
	b.WriteString(newLog + "\n")
	return b.Bytes(), nil
}
