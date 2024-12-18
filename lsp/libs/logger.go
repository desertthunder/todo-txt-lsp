// package libs contains helpers and utilities
package libs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

func CreateLogDir() error {
	return os.Mkdir("logs", 0755)
}

// CreateLogFile creates a file with the format:
// Name Format: {context}_{year}-{month}-{day}-{hour}-{minute}-{second}_lsp.log
func CreateLogFile(context string) *os.File {
	err := CreateLogDir()

	logger := log.Default()

	if err != nil && !strings.Contains(err.Error(), "file exists") {
		logger.Errorf("error creating logs directory: %v", err)
		return nil
	}

	d := time.Now()
	fdate := d.Format(time.DateOnly)
	ftime := fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
	fname := fmt.Sprintf("%s_%s_%s.log", context, fdate, ftime)

	f, err := os.Create(fmt.Sprintf("logs/%s", fname))

	if err != nil {
		logger.Errorf("error creating log file: %v", err)

		return nil
	}

	return f
}

// CreateLogger creates a file logger
func CreateLogger(ctx string) *log.Logger {
	f := CreateLogFile(ctx)
	return log.NewWithOptions(f, log.Options{
		Level:      log.InfoLevel,
		Prefix:     fmt.Sprintf("[lsp|%s 🚀]", ctx),
		TimeFormat: time.RFC3339,
	})
}
