package lsp

import "encoding/json"

func HandleInitialize(data []byte) (*InitializeParams, error) {
	params := InitializeParams{}

	if err := json.Unmarshal(data, &params); err != nil {
		return nil, err
	}

	logger.Debug("params:", params)

	return &params, nil
}

func CreateInitializeResult(p InitializeParams) InitializeResult {
	return InitializeResult{
		Capabilities: ServerCapabilities{
			RenameProvider:                   true,
			HoverProvider:                    true,
			FoldingRangeProvider:             true,
			DocumentFormattingProvider:       true,
			DocumentRangeFormattingProvider:  true,
			DocumentOnTypeFormattingProvider: true,
			CompletionProvider: CompletionOptions{
				ResolveProvider: true,
				TriggerCharacters: []string{
					"(", "[", "-",
				},
				AllCommitCharacters: []string{
					")", "]", "_", "x",
				},
				CompletionItem: CompletionItem{
					LabelDetailsSupport: true,
				},
			},
			PositionEncoding: UTF8,
			TextDocumentSync: TextDocumentSyncOptions{
				OpenClose: true,
				Change:    IncrementalSync,
			},
		},
		ServerInfo: ServerInfo{
			Name:    "todo.txt language server",
			Version: "0.1.0",
		},
	}
}
