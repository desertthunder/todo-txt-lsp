// package libs contains helpers and utilities
package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// CreateLogFile creates a file with the format:
// Name Format: {year}-{month}-{day}-{hour}-{minute}-{second}_lsp.log
func CreateLogFile() *os.File {
	d := time.Now()
	fdate := d.Format(time.DateOnly)
	ftime := fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
	fname := fmt.Sprintf("%s_%s.log", fdate, ftime)

	f, err := os.Create(fname)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

// CreateLogger creates a file logger
func CreateLogger() *log.Logger {
	f := CreateLogFile()

	return log.NewWithOptions(f, log.Options{
		Level:      log.InfoLevel,
		Prefix:     "[lsp 🚀]",
		TimeFormat: time.RFC3339,
	})
}
