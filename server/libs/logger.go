// package libs contains helpers and utilities
package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var logger = CreateLogger()

func CreateLogDir() error {
	return os.Mkdir("logs", 0755)
}

// CreateLogFile creates a file with the format:
// Name Format: lsp_{year}-{month}-{day}-{hour}-{minute}-{second}.log
// if args are not provided
func CreateLogFile() *os.File {
	var f *os.File
	var err error

	if len(os.Args) < 3 {
		d := time.Now()
		fdate := d.Format(time.DateOnly)
		ftime := fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
		fname := fmt.Sprintf("lsp_%s_%s.log", fdate, ftime)
		f, err = os.CreateTemp("", fname)
	} else {
		fname := os.Args[2]
		f, err = os.Create(fname)
	}

	if err != nil {
		log.Default().Errorf("unable to create temp file %s", err.Error())
		return nil
	}

	return f
}

func CreateLogger() *log.Logger {
	f := CreateLogFile()

	if f == nil {
		return log.Default()
	}

	return log.NewWithOptions(f, log.Options{
		Level:           log.DebugLevel,
		TimeFormat:      time.DateTime,
		Prefix:          "[lsp 🚀]",
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}

func GetLogger() *log.Logger {
	return logger
}
