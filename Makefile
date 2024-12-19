.PHONY: build help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "	build -	build lsp binary and compile vsc client"
	@echo "	help  -	display this help message"

build:
	@cd server; \
		echo "building lsp binary"; \
		go build; \
		echo "adding todo_txt_lsp to GOPATH"; \
		go install; \
		echo "adding todo_txt_lsp to PATH. please wait"; \
		asdf reshim; \
		echo "";
	@cd client; \
		echo "compiling vscode client"; \
		pnpm compile;


