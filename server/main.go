// package main is the application entry point
//
// Runs the server. This uses stdin & stdout as the transport
// mechanism for communication between client and server. i.e.
// each entity receives through stdin and sends through stdout,
// thus logs need to be stored in files or sent to stderr.
//
// Messages are decoded by handlers in `jrpc` package and then
// processed by an appropriate handler in the `lsp` package.
// This handler returns a Result that is encoded by jrpc and sent
// to the client.
package main

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/desertthunder/todo_txt_lsp/jrpc"
	"github.com/desertthunder/todo_txt_lsp/libs"
)

var logger = libs.GetLogger()

func main() {
	// writer := os.Stdout
	reader := bufio.NewReader(os.Stdin)
	logger.Infof("Starting Server (%s)...", time.Now().Format(time.Kitchen))

	for {
		rawMsg, err := jrpc.ReadMessage(reader, "\r\n\r\n")
		if err != nil {
			if err == io.EOF {
				break
			}

			logger.Fatalf("error reading message: %v", err)
		}

		contentLength, err := rawMsg.ParseContentLength()
		if err != nil {
			logger.Fatalf("Invalid Content-Length: %v", err)
		}

		logger.Infof("received payload of len %d", contentLength)

		// TODO: use the message struct
		_, err = jrpc.DecodeMessage(reader, contentLength)

		if err != nil {
			logger.Printf("error handling message: %v", err)
		}

	}

	os.Exit(0)
}
