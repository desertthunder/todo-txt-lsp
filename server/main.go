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
	"github.com/desertthunder/todo_txt_lsp/lsp"
)

var logger = libs.GetLogger()

type State struct {
	RecInitialize  bool
	RecInitialized bool
}

func (s State) Empty() bool {
	return !s.RecInitialize && !s.RecInitialized
}

func NewStateMachine() State {
	return State{RecInitialized: false}
}

func main() {
	state := NewStateMachine()
	writer := os.Stdout
	reader := bufio.NewReader(os.Stdin)
	logger.Infof("Starting Server (%s)...", time.Now().Format(time.Kitchen))

	for {
		m, err := jrpc.ReadMessage(reader, "\r\n\r\n")
		if err != nil {
			if err == io.EOF {
				break
			}

			logger.Fatalf("error reading message: %v", err)
		}

		logger.Infof("received payload of len %d", m.ContentLength)

		msg, err := jrpc.DecodeMessage(reader, m)
		if err != nil {
			logger.Errorf("error decoding message: %v", err)
			continue
		}

		if msg.Method == string(lsp.InitializeMethod) {
			state.RecInitialize = true
		}

		if state.RecInitialize && msg.Method == string(lsp.Initialized) {
			logger.Infof("transport opened 🎉")
			state.RecInitialized = true
			continue
		}

		if state.Empty() {
			logger.Error("initialized notification not received from client.")
			os.Exit(1)
		}

		r, err := lsp.HandleMessage(msg, m.Payload)
		if err != nil {
			logger.Errorf("error handling message (%s): %v", string(m.Payload), err.Error())
			continue
		}

		resp, err := jrpc.EncodeMessage(r)
		logger.Debugf("encoded msg: %s", resp)
		if err != nil {
			logger.Errorf("error encoding message: %v", err.Error())
			continue
		}

		if err = jrpc.WriteResponse(writer, resp); err != nil {
			logger.Errorf("unable to write response %v", err.Error())
		}
	}

	logger.Info("reached end of file")
	os.Exit(0)
}
