# Examples

Each directory is a runnable program against the public API:

- `basic/` - one agent, one completion.
- `agenthandoff/` - two agents handing a task over.
- `collaborative/` - a shared conversation with a tool call.

Run one with `go run ./examples/<name>`; they read the provider keys from the environment.
