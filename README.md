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
└── server
|   └── lsp
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

1. Inlay hints/Intellisense
   Hover
2. Project awareness
3. Context awareness
4. Go To Definition (first instance of a project, or key)
5. Folding by date
6. Rename symbols (projects, contexts, keys)
7. Warnings for out of order (date)
8. Filtered views

## Server

The server package for todo.txt language support.

### LSP Package

Files prefixed with method contain types.

Files prefixed with sync or feature contain request/response handlers.

### Debugging

```bash
go build && go install && asdf reshim && pnpm compile

# After starting the debugger
tail -f $(echo "console.log(require('os').tmpdir())" | node)/lsp.log
```
