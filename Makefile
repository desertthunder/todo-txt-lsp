.PHONY: build help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "	build -	build lsp binary"
	@echo "	help  -	display this help message"

build:
	@cd lsp; \
		mkdir -p ./tmp/bin; \
		echo "building & compiling language server..."
		go build -o ./tmp/bin ./...; \
		echo "built lsp to ./tmp/bin"; \
		echo "binary path: ./tmp/bin/$$(ls ./tmp/bin/)"


