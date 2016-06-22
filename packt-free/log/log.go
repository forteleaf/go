package log

import (
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

const (
	// LogFile is path to log file
	LogFile = "log/packtFree.log"
)

// Log is custom logrus logger
var Log = logrus.New()

func init() {
	f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	os.MkdirAll(filepath.Dir(LogFile), os.ModePerm)
	if err != nil {
		f, err = os.Create(LogFile)
	}
	Log.Out = f
}
