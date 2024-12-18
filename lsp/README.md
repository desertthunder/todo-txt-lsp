# LSP

The lsp package for todo.txt language support.

## Notes

### Lifecycle

The current protocol specification defines that the lifecycle of a server is
managed by the client (e.g. an editor). It is up to the client to decide when to
start (process-wise) and when to shutdown a server.

### LSP Docs

See [LSP](https://microsoft.github.io/language-server-protocol/specification).

```typescript
/**
 * Completion options.
 */
export interface CompletionOptions extends WorkDoneProgressOptions {
	/**
	 * The additional characters, beyond the defaults provided by the client (typically
	 * [a-zA-Z]), that should automatically trigger a completion request. For example
	 * `.` in JavaScript represents the beginning of an object property or method and is
	 * thus a good candidate for triggering a completion request.
	 *
	 * Most tools trigger a completion request automatically without explicitly
	 * requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	 * do so when the user starts to type an identifier. For example if the user
	 * types `c` in a JavaScript file code complete will automatically pop up
	 * present `console` besides others as a completion item. Characters that
	 * make up identifiers don't need to be listed here.
	 */
	triggerCharacters?: string[];

	/**
	 * The list of all possible characters that commit a completion. This field
	 * can be used if clients don't support individual commit characters per
	 * completion item. See client capability
	 * `completion.completionItem.commitCharactersSupport`.
	 *
	 * If a server provides both `allCommitCharacters` and commit characters on
	 * an individual completion item the ones on the completion item win.
	 *
	 * @since 3.2.0
	 */
	allCommitCharacters?: string[];

	/**
	 * The server provides support to resolve additional
	 * information for a completion item.
	 */
	resolveProvider?: boolean;

	/**
	 * The server supports the following `CompletionItem` specific
	 * capabilities.
	 *
	 * @since 3.17.0
	 */
	completionItem?: {
		/**
		 * The server has support for completion item label
		 * details (see also `CompletionItemLabelDetails`) when receiving
		 * a completion item in a resolve call.
		 *
		 * @since 3.17.0
		 */
		labelDetailsSupport?: boolean;
	}
}
```

```typescript
export interface SignatureHelpOptions extends WorkDoneProgressOptions {
	/**
	 * The characters that trigger signature help
	 * automatically.
	 */
	triggerCharacters?: string[];

	/**
	 * List of characters that re-trigger signature help.
	 *
	 * These trigger characters are only active when signature help is already
	 * showing. All trigger characters are also counted as re-trigger
	 * characters.
	 *
	 * @since 3.15.0
	 */
	retriggerCharacters?: string[];
}
```
