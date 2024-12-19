package lsp

import (
	"fmt"
	"io"

	"github.com/desertthunder/todo_txt_lsp/jrpc"
	"github.com/desertthunder/todo_txt_lsp/libs"
)

var logger = libs.GetLogger()

func WriteResponse(w io.Writer, resp string) error {
	_, err := w.Write([]byte(resp))

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// HandleMessage handles a JSON-RPC message
//
//  1. Unmarshals the JSON message and deserializes the message content (params)
//
//  2. Based on the method in the message, a handler is dispatched.
//     This Handler creates a response (a Result object/struct)
//
//  3. This method serializes the response to JSON and calls writes to stdout
func HandleMessage(msg jrpc.Message, content []byte, w io.Writer) error {
	var resp string
	var err error

	switch GetMethod(msg.Method) {
	case InitializeMethod:
		logger.Info("handling initialize message")

		_, err = HandleInitializeMessage(content)

		if err != nil {
			logger.Error(err.Error())
			return err
		}

		r := CreateInitializeResult(&msg.ID)
		resp, err = jrpc.EncodeMessage(r)
		break
	case HoverMethod:
		var p *HoverParams
		var r *HoverResult
		p, err = HandleHoverMessage(content)
		if err != nil {
			return err
		}

		r, err = CreateHoverResult(*p)
		resp, err = jrpc.EncodeMessage(r)
		break
	case DidSaveTextDocument:
		err = fmt.Errorf("did save text doc")
	default:
		err = fmt.Errorf("method %s not implemented", msg.Method)
	}

	if err != nil {
		return err
	}

	return WriteResponse(w, resp)
}
