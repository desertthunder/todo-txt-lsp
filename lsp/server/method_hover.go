package server

const HoverFeature Method = "textDocument/hover"

type HoverParams struct{}

type RangePos struct {
	Line      int
	Character int
}

type Range struct {
	Start RangePos `json:"start"`
	End   RangePos `json:"end"`
}

type HoverResult struct {
	Contents string `json:"contents"`
	Range    Range  `json:"range"`
}
