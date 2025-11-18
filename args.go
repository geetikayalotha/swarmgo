package swarmgo

import (
	"encoding/json"
	"log"
)

// parseToolArguments decodes the model's tool arguments without losing a malformed payload.
// An unparsable string is recorded as {"raw": ...} so the call stays visible in the result.
func parseToolArguments(toolName, raw string) interface{} {
	if raw == "" {
		return map[string]interface{}{}
	}
	var args interface{}
	if err := json.Unmarshal([]byte(raw), &args); err != nil {
		log.Printf("tool %s returned invalid arguments: %v", toolName, err)
		return map[string]interface{}{"raw": raw}
	}
	return args
}
