# Todo.txt Language Server & Grammar

## Packages

1. `lsp` - a golang implementation of the language server protocol for todo.txt
files.
2. `syntax-highlight` - a vscode extension that provides syntax highlighting for
todo.txt files.

```plaintext
.
|
└── syntax-highlight
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
