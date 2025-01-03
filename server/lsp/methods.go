package lsp

type Method string

const (
	NotImplementedMethod Method = ""
	InitializeMethod     Method = "initialize"
	Initialized          Method = "initialized"
	HoverMethod          Method = "textDocument/hover"
)

func GetMethod(m string) Method {
	methods := []Method{
		InitializeMethod,
		HoverMethod,
	}

	for _, method := range methods {
		if m == string(method) {
			return method
		}
	}

	return NotImplementedMethod
}

// Parking Lot: These methods need to be implemented and require type
// definitions. See the document synchronization section of the spec
// for more information.
const (
	// Sync Methods
	DidChangeTextDocument         Method = "textDocument/didChange"
	WillSaveTextDocument          Method = "textDocument/willSave"
	WillSaveWaitUntilTextDocument Method = "textDocument/willSaveWaitUntil"
	DidSaveTextDocument           Method = "textDocument/didSave"
	DidCloseTextDocument          Method = "textDocument/didClose"
	RenameTextDocument            Method = "textDocument/rename"

	// Language Features
	CompletionFeature Method = "textDocument/completion"
)
