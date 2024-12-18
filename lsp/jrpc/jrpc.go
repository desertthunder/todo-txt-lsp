// package jrpc
package jrpc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/log"
)

type Params struct{}

var Separator = []byte{'\r', '\n', '\r', '\n'}
var SepNotFound = errors.New("carriage returns not found in message")

/*
Message represents a JSON-RPC message as defined in the LSP spec:
https://microsoft.github.io/language-server-protocol/specification

Example:

	{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "textDocument/completion",
		"params": {
			...
		}
	}
*/
type Message struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

func EncodeMessage(m interface{}) (string, error) {
	content, err := json.Marshal(m)

	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

	return msg, nil
}

// DecodeMessage takes the input stream and decodes the message
// by cutting the Content-Length & carriage returns, then
// deserializing the json.
func DecodeMessage(b []byte, logger *log.Logger) (Message, error) {
	m := Message{}
	headers, contents, found := bytes.Cut(b, Separator)

	if !found {
		return m, SepNotFound
	}

	len_bytes := headers[len("Content-Length: "):]
	c, err := strconv.Atoi(string(len_bytes))

	if err != nil {
		return m, err
	}

	if len(contents) < c {
		err = errors.New("message too short")
		return m, err
	}

	t := len(headers) + 4 + c // Note: the +4 is for the \r\n\r\n

	logger.Debug("Message length:", t)

	err = json.Unmarshal(contents[:c], &m)

	if err != nil {
		return m, err
	}

	logger.Info("Received message:", m.ID, m.Method, m.Params)

	return m, nil
}

/*
Split is a bufio.SplitFunc that tokenizes JSON messages from the input stream.
We take the content length and two carriage returns and return the remaining
bytes.

Example: Content-Length: ...\r\n
\r\n

	{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "textDocument/completion",
		"params": {
			...
		}
	}

```
*/
var Split bufio.SplitFunc = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	headers, contents, found := bytes.Cut(data, Separator)

	if !found {
		return 0, nil, SepNotFound
	}

	content, err := strconv.Atoi(string(headers[len("Content-Length: "):]))

	if err != nil {
		return 0, nil, err
	}

	if len(contents) < content {
		return 0, nil, errors.New("message should be larger in size than the provided content length")
	}

	total := len(headers) + 4 + content // Note: the +4 is for the \r\n\r\n

	return total, contents[:total], nil
}
