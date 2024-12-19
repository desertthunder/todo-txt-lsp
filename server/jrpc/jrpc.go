// package jrpc defines methods for managing the transport layer
// between the language server and client (editor)
package jrpc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/desertthunder/todo_txt_lsp/libs"
)

type Params struct{}

var logger = libs.GetLogger()

/*
Message represents a JSON-RPC message as defined in the LSP spec:
https://microsoft.github.io/language-server-protocol/specification

Example:

	{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "textDocument/completion",
		"params": { ... }
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
func DecodeMessage(r *bufio.Reader, c int) (Message, error) {
	m := Message{}
	payload := make([]byte, c)

	if _, err := io.ReadFull(r, payload); err != nil {
		return m, fmt.Errorf("error reading payload: %v", err)
	}

	logger.Debugf("full message: %s", string(payload))

	if err := json.Unmarshal(payload, &m); err != nil {
		return m, err
	}

	logger.Infof("Received message %d: %s", m.ID, m.Method)

	return m, nil
}

type RawMessage struct {
	ContentLength []byte
	Payload       []byte
}

func ReadMessage(r *bufio.Reader, d string) (*RawMessage, error) {
	var buf bytes.Buffer

	m := RawMessage{}
	delim := []byte(d)

	for {
		b, err := r.ReadByte()

		if err != nil {
			return &RawMessage{}, err
		}

		buf.WriteByte(b)

		if buf.Len() >= len(delim) && bytes.HasSuffix(buf.Bytes(), delim) {
			data := buf.Bytes()[:buf.Len()-len(delim)]
			remaining := buf.Bytes()[buf.Len()-len(delim):]

			m.ContentLength = data
			m.Payload = remaining

			return &m, nil
		}
	}
}

func (m RawMessage) ParseContentLength() (int, error) {
	parts := strings.SplitN(string(m.ContentLength), ":", 2)

	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid header line: %s", m.ContentLength)
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	if key != "Content-Length" {
		return 0, errors.New("invalid key:value pair")
	}

	return strconv.Atoi(value)
}
