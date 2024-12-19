package lsp

import (
	"fmt"

	"github.com/desertthunder/todo_txt_lsp/jrpc"
	"github.com/desertthunder/todo_txt_lsp/libs"
)

var logger = libs.GetLogger()

// HandleMessage handles a JSON-RPC message
//
//  1. Unmarshals the JSON message and deserializes the message content (params)
//
//  2. Based on the method in the message, a handler is dispatched.
//     This Handler creates a response (a Result object/struct)
//
//  3. This method serializes the response to JSON and calls writes to stdout
func HandleMessage(msg jrpc.Message, content []byte) (interface{}, error) {
	var err error

	switch GetMethod(msg.Method) {
	case InitializeMethod:
		_, err = HandleInitializeMessage(content)
		if err != nil {
			logger.Errorf("error handling initialize request message: %v", err.Error())
			break
		}

		r := CreateInitializeResult()
		res := CreateInitializeResponse(&msg.ID, r)
		return res, nil
	case HoverMethod:
		p, err := HandleHoverMessage(content)
		if err != nil {
			logger.Errorf("error handling hover request message: %v", err.Error())
			break
		}

		r := CreateHoverResult(*p)
		res := CreateHoverResponse(&msg.ID, r)
		return res, nil
	default:
		err = fmt.Errorf("method %s not implemented", msg.Method)
	}

	return nil, err
}
