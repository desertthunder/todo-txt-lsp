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
