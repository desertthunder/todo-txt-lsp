import * as path from "path";
import { workspace, ExtensionContext } from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";
import { tmpdir } from "os";

const log = "lsp.log";
let client: LanguageClient;

export function activate(context: ExtensionContext) {
  // Golang LSP - Built to /tmp/bin
  const serverModule = context.asAbsolutePath(
    path.join("tmp", "bin", "todo_txt_lsp")
  );

  const serverOptions = {
    run: { command: "todo_txt_lsp", transport: TransportKind.stdio },
    debug: {
      command: "todo_txt_lsp",
      transport: TransportKind.stdio,
      args: ["--file", `${tmpdir}/${log}`],
    },
  } as ServerOptions;

  const clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: "file", language: "plaintext" }],
  };

  client = new LanguageClient(
    "todolsp",
    "Todo.txt LSP",
    serverOptions,
    clientOptions
  );

  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  return client ? client.stop() : undefined;
}
