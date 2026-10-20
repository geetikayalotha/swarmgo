# llm

The provider adapters. Each file implements the same client interface for one backend
(OpenAI, Anthropic, Gemini, Ollama); `llm.go` defines the request, response and tool types the
 swarm runtime passes around.
