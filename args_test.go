package swarmgo

import "testing"

func TestParseToolArguments(t *testing.T) {
	cases := []struct {
		name    string
		raw     string
		wantRaw bool
	}{
		{"valid", `{"city":"Paris"}`, false},
		{"empty", ``, false},
		{"truncated", `{"city":`, true},
	}
	for _, tc := range cases {
		got := parseToolArguments("demo", tc.raw)
		m, ok := got.(map[string]interface{})
		if !ok {
			t.Fatalf("%s: expected an object, got %T", tc.name, got)
		}
		if _, has := m["raw"]; has != tc.wantRaw {
			t.Fatalf("%s: raw presence = %v, want %v", tc.name, has, tc.wantRaw)
		}
	}
}
