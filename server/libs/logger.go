// package libs contains helpers and utilities
package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var logger *log.Logger

func CreateLogDir() error {
	return os.Mkdir("logs", 0755)
}

// CreateLogFile creates a file with the format:
// Name Format: {context}_{year}-{month}-{day}-{hour}-{minute}-{second}_lsp.log
func CreateLogFile(context string) *os.File {
	var f *os.File
	var err error

	if len(os.Args) < 3 {
		d := time.Now()
		fdate := d.Format(time.DateOnly)
		ftime := fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
		fname := fmt.Sprintf("%s_%s_%s.log", context, fdate, ftime)
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

// CreateLogger creates a file logger
func CreateLogger(ctx string) *log.Logger {
	f := CreateLogFile(ctx)

	if f == nil {
		return log.Default()
	}

	return log.NewWithOptions(f, log.Options{
		Level:           log.DebugLevel,
		Prefix:          "[lsp 🚀]",
		TimeFormat:      time.RFC3339,
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}

func GetLogger() *log.Logger {
	return CreateLogger("")
}
