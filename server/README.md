# Server

The server package for todo.txt language support.

## LSP Package

Files prefixed with method contain types.

Files prefixed with sync or feature contain request/response handlers.

## Debug

```bash
go build && go install && asdf reshim && pnpm compile

# After starting the debugger
tail -f $(echo "console.log(require('os').tmpdir())" | node)/lsp.log
```

## Notes

### Lifecycle

The current protocol specification defines that the lifecycle of a server is
managed by the client (e.g. an editor). It is up to the client to decide when to
start (process-wise) and when to shutdown a server.
