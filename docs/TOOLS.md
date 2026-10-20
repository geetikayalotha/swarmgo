# Tool results

A tool call produces a `ToolResult` with the tool name, the parsed arguments and the handler result.

- Valid JSON arguments are decoded into an object.
- An empty argument string yields `{}` (a tool that takes no arguments is normal).
- Unparsable arguments are recorded as `{"raw": "<the original string>"}` and logged with the tool
  name, so a truncated model output is visible instead of looking like a call with no arguments.

The same helper is used by the streaming and non-streaming paths, so both report identically.
