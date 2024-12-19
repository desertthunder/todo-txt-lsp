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
}

func EncodeMessage(m interface{}) (string, error) {
	content, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

	logger.Debugf("message to client: %s", msg)

	return msg, nil
}

// DecodeMessage takes the input stream and decodes the message
// by cutting the Content-Length & carriage returns, then
// deserializing the json.
func DecodeMessage(r *bufio.Reader, m *RawMessage) (Message, error) {
	msg := Message{}
	payload := make([]byte, m.ContentLength)

	if _, err := io.ReadFull(r, payload); err != nil {
		return msg, fmt.Errorf("error reading payload: %v", err)
	}

	m.Payload = payload

	if err := json.Unmarshal(payload, &msg); err != nil {
		logger.Errorf("error unmarshalling payload: %s", err.Error())
		return msg, err
	}

	logger.Infof("Received message id: %d, %s", msg.ID, msg.Method)

	return msg, nil
}

type RawMessage struct {
	ContentLengthHeader []byte
	Payload             []byte
	ContentLength       int
}

func ReadMessage(r *bufio.Reader, d string) (*RawMessage, error) {
	var buf bytes.Buffer

	m := RawMessage{}
	delim := []byte(d)

	for {
		b, err := r.ReadByte()
		if err != nil {
			logger.Errorf("error reading buffer: %v", err.Error())
			return &RawMessage{}, err
		}

		buf.WriteByte(b)

		if buf.Len() >= len(delim) && bytes.HasSuffix(buf.Bytes(), delim) {
			data := buf.Bytes()[:buf.Len()-len(delim)]
			m.ContentLengthHeader = data

			if err := m.ParseContentLength(); err != nil {
				logger.Fatalf("invalid Content-Length: %v", err)
			}

			return &m, nil
		}
	}
}

func (m *RawMessage) ParseContentLength() error {
	var err error
	parts := strings.SplitN(string(m.ContentLengthHeader), ":", 2)

	if len(parts) != 2 {
		m.ContentLength = 0
		return fmt.Errorf("invalid header line: %s with parts: %d", m.ContentLengthHeader, len(parts))
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	if key != "Content-Length" {
		m.ContentLength = 0
		return errors.New("invalid key:value pair")
	}

	m.ContentLength, err = strconv.Atoi(value)
	return err
}

func WriteResponse(w io.Writer, resp string) error {
	logger.Debug("writing response")

	_, err := w.Write([]byte(resp))
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
