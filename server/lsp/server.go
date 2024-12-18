// package lsp
package lsp

import (
	"fmt"
	"io"
	"os"

	"github.com/desertthunder/todo_txt_lsp/jrpc"
	"github.com/desertthunder/todo_txt_lsp/libs"
)

var logger = libs.CreateLogger("server")

func WriteResponse(w io.Writer, resp string) error {
	w.Write([]byte(resp))
	return nil
}

// HandleMessage handles a JSON-RPC message
//
// 1: Unmarshals the JSON message and deserializes the message content
// 2. Handler creates a response
// 3. This method marshals the response and calls WriteResponse
func HandleMessage(m string, c []byte) error {
	w := os.Stdout
	switch GetMethod(m) {
	case InitializeMethod:
		p, err := HandleInitialize(c)
		if err != nil {
			return err
		}

		r := CreateInitializeResult(*p)

		resp, err := jrpc.EncodeMessage(r)
		if err != nil {
			return err
		}

		return WriteResponse(w, resp)
	case HoverFeature:
		logger.Info("Hover feature not implemented")
		return fmt.Errorf("method %s not implemented", m)
	default:
		return fmt.Errorf("method %s not implemented", m)
	}
}
