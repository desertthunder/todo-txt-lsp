// package main is the application entry point
//
// Runs the server. This uses stdin & stdout as the transport
// mechanism for communication between client and server.
package main

import (
	"bufio"
	"os"

	"github.com/charmbracelet/log"

	"github.com/desertthunder/todo_txt_lsp/jrpc"
	"github.com/desertthunder/todo_txt_lsp/libs"
)

var logger *log.Logger

func init() {
	logger = libs.CreateLogger("scanner")
}

func main() {
	logger.Info("Starting server...")

	s := bufio.NewScanner(os.Stdin)
	s.Split(jrpc.Split)

	for s.Scan() {
		b := s.Bytes()
		logger.Info("Received message:", string(b))
	}
}
