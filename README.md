# Todo.txt Language Server & Grammar

TextMate grammar for todo.txt files for semantic token highlighting and server
for completions and hover information. The server is based on LSP v3.17.0. The
purpose of this set of extensions is to provide a more seamless experience for
using todo.txt as an in-workspace/in-repo solution for managing tasks for
projects.

## Packages

1. `lsp` - a golang implementation of the language server protocol for todo.txt
files.
2. `syntax` - a vscode extension that provides syntax highlighting for
todo.txt files.

```plaintext
.
|
└── syntax
|   └── examples
|   └── syntaxes
|       └── todo-txt.tmLanguage.json
└── lsp
|   └── server
|   └── jrpc
└── (treesitter)
└── (zed-lang)
└── (todo.nvim)
```

*Note* that the packages in parens are planned.

## LSP

The `lsp` package is a golang implementation of the language server protocol for
todo.txt files. It is reliant on only a single external dependency, `charmbracelet/log`,
but otherwise builds parsing and request-response handling from the standard library.

## Proposed Features
