package server

type ClientInfo struct {
	Name string `json:"name"`
	// Client version
	// Optional
	Version string `json:"version"`
}

type ClientCapabilities struct {
	// Optional
	Workspace Workspace `json:"workspace"`
	// Optional
	TextDocument TextDocument `json:"textDocument"`
	// Optional
	NotebookDocument NotebookDocument `json:"notebookDocument"`
	// Window specific capabilities
	// Optional
	Window Window `json:"window"`
	// Optional
	General General `json:"general"`
	// Optional
	Experimental interface{} `json:"experimental"`
}

type WorkspaceFolder struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

type TraceValue string

const (
	Off      TraceValue = "off"
	Verbose  TraceValue = "verbose"
	Messages TraceValue = "messages"
)

type InitializeParams struct {
	// Parent process ID (nullable)
	ProcessID int `json:"processId"`
	// Client metadata
	// Optional
	ClientInfo ClientInfo `json:"clientInfo"`
	// IETF Language Tag
	// Optional (default is en-us)
	Locale  string `json:"locale"`
	RootURI string `json:"rootUri"`
	// Initialization options from user
	// Optional
	InitializationOptions interface{} `json:"initializationOptions"`
	// The capabilities provided by the client (editor or tool)
	// See ClientCapabilities in spec
	Capabilities ClientCapabilities `json:"capabilities"`
	// If omitted, this value should be set to 'off'.
	// Optional
	Trace TraceValue `json:"trace"`
	// Can be null if the client does not support workspace folders
	// or none are configured.
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders"`
}

// TODO: See WorkspaceEditClientCapabilities in spec
type WorkspaceEdit struct{}

// TODO: See DidChangeConfigurationClientCapabilities in spec
type DidChangeConfiguration struct{}

// TODO: See DidChangeWatchedFilesClientCapabilities in spec
type DidChangeWatchedFiles struct{}

// TODO: See SymbolClientCapabilities in spec
type Symbol struct{}

// TODO: See ExecuteCommandClientCapabilities in spec
type ExecuteCommand struct{}

// TODO: See SemanticTokensClientCapabilities in spec
type SemanticTokens struct{}

// TODO: See CodeLensClientCapabilities in spec
type CodeLens struct{}

type FileOperations struct {
	DynamicRegistration bool `json:"dynamicRegistration"`
	DidCreate           bool `json:"didCreate"`
	WillCreate          bool `json:"willCreate"`
	DidRename           bool `json:"didRename"`
	WillRename          bool `json:"willRename"`
	DidDelete           bool `json:"didDelete"`
	WillDelete          bool `json:"willDelete"`
}

// TODO: See InlineValueWorkspaceClientCapabilities in spec
type InlineValue struct{}

// TODO: See InlayHintWorkspaceClientCapabilities in spec
type InlayHint struct{}

// TODO: See DiagnosticWorkspaceClientCapabilities in spec
type Diagnostics struct{}

type Workspace struct {
	ApplyEdit      bool                  `json:"applyEdit"`
	WorkspaceEdit  WorkspaceEdit         `json:"workspaceEdit"`
	DidChangeWatch DidChangeWatchedFiles `json:"didChangeWatchedFiles"`
	Symbol         Symbol                `json:"symbol"`
	ExecuteCommand ExecuteCommand        `json:"executeCommand"`
	SemanticTokens SemanticTokens        `json:"semanticTokens"`
	CodeLens       CodeLens              `json:"codeLens"`
	InlineValue    InlineValue           `json:"inlineValue"`
	InlayHint      InlayHint             `json:"inlayHint"`
	Diagnostics    Diagnostics           `json:"diagnostics"`
	FileOperations FileOperations        `json:"fileOperations"`
}

// TODO: See NotebookDocumentClientCapabilities in spec
type NotebookDocument struct{}

// TODO: See TextDocumentClientCapabilities in spec
type TextDocument struct{}

// TODO: See ShowMessageRequestClientCapabilities in spec
type ShowMessage struct{}

// TODO: See ShowDocumentClientCapabilities in spec
type ShowDocument struct{}

type Window struct {
	WorkDoneProgress bool         `json:"workDoneProgress"`
	ShowMessage      ShowMessage  `json:"showMessage"`
	ShowDocument     ShowDocument `json:"showDocument"`
}

type StaleRequestSupport struct {
	Cancel bool     `json:"cancel"`
	Retry  []string `json:"retryOnContentModified"`
}

// TODO: See RegularExpressionsClientCapabilities in spec
type RegularExpressions struct{}

// TODO: See MarkdownClientCapabilities in spec
type Markdown struct{}

// TODO: See PositionEncodingKind in spec
type PositionEncoding struct{}

type General struct {
	// Optional
	StaleRequestSupport StaleRequestSupport `json:"staleRequestSupport"`
	// Optional
	RegularExpressions RegularExpressions `json:"regularExpressions"`
	// Optional
	Markdown Markdown `json:"markdown"`
	// Optional
	PositionEncodings []PositionEncoding `json:"positionEncodings"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

const InitializeMethod Method = "initialize"
