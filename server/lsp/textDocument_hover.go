package lsp

import "encoding/json"

type MarkupKind string

const (
	markdown  MarkupKind = "markdown"
	plaintext MarkupKind = "plaintext"
)

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type RangePos struct {
	Line      int
	Character int
}

type Range struct {
	Start RangePos `json:"start"`
	End   RangePos `json:"end"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

func (m *MarkupContent) FromFile() {
	m.Kind = markdown
	m.Value = "File Contents+"
}

type HoverResult struct {
	Contents MarkupContent `json:"contents"`
}

type HoverResultWithRange struct {
	HoverResult
	Range Range `json:"range"`
}

type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

func HandleHoverMessage(data []byte) (*HoverParams, error) {
	params := HoverParams{}
	if err := json.Unmarshal(data, &params); err != nil {
		return nil, err
	}

	return &params, nil
}

func CreateHoverResult(params HoverParams) HoverResult {
	return HoverResult{
		MarkupContent{markdown,
			`# Hello World

			This is likely where I'll import documentation from markdown files.
			`,
		}}
}

func CreateHoverResponse(id *int, result HoverResult) HoverResponse {
	return HoverResponse{BaseResponse(id), result}
}
