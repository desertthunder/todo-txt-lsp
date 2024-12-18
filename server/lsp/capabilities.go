package lsp

type PositionEncodingKind string
type TextDocumentSyncKind int

const (
	UTF8  PositionEncodingKind = "utf-8"
	UTF16 PositionEncodingKind = "utf-16"
	UTF32 PositionEncodingKind = "utf-32"
)

const (
	NoSync          TextDocumentSyncKind = 0
	FullSync        TextDocumentSyncKind = 1
	IncrementalSync TextDocumentSyncKind = 2
)

type TextDocumentSyncOptions struct {
	// Whether or not open/close notifications should be sent.
	OpenClose bool                 `json:"openClose"`
	Change    TextDocumentSyncKind `json:"change"`
}

type CompletionItem struct {
	LabelDetailsSupport bool `json:"labelDetailsSupport"`
}

type CompletionOptions struct {
	TriggerCharacters   []string       `json:"triggerCharacters"`
	AllCommitCharacters []string       `json:"allCommitCharacters"`
	ResolveProvider     bool           `json:"resolveProvider"`
	CompletionItem      CompletionItem `json:"completionItem"`
}

type SignatureHelpOptions struct {
	TriggerCharacters   []string `json:"triggerCharacters"`
	RetriggerCharacters []string `json:"retriggerCharacters"`
}

type ServerCapabilities struct {
	PositionEncoding                 PositionEncodingKind    `json:"positionEncoding"`
	TextDocumentSync                 TextDocumentSyncOptions `json:"textDocumentSync"`
	CompletionProvider               CompletionOptions       `json:"completionProvider"`
	SignatureHelpProvider            SignatureHelpOptions    `json:"signatureHelpProvider"`
	DocumentFormattingProvider       bool                    `json:"documentFormattingProvider"`
	DocumentRangeFormattingProvider  bool                    `json:"documentRangeFormattingProvider"`
	DocumentOnTypeFormattingProvider bool                    `json:"documentOnTypeFormattingProvider"`
	DeclarationProvider              bool                    `json:"declarationProvider"`
	RenameProvider                   bool                    `json:"renameProvider"`
	FoldingRangeProvider             bool                    `json:"foldingRangeProvider"`
	InlineValueProvider              bool                    `json:"inlineValueProvider"`
	HoverProvider                    bool                    `json:"hoverProvider"`
}
