package lsp

import "encoding/json"

const docSyncDidOpen Method = "textDocument/didOpen"

type DidOpenParams struct{}

func HandleDocDidOpen(data []byte) (*DidOpenParams, error) {
	params := DidOpenParams{}
	if err := json.Unmarshal(data, &params); err != nil {
		return nil, err
	}

	return &params, nil
}
