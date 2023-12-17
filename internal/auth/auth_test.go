package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input    http.Header
		expected string
	}{
		"header present":        {input: http.Header{"Authorization": []string{"ApiKey 1235"}}, expected: "12345"},
		"header missing":        {input: http.Header{"Foo": []string{"123"}}, expected: ""},
		"malformed auth header": {input: http.Header{"Authorization": []string{"ApiKey12345"}}, expected: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, _ := GetAPIKey(tc.input)
			if actual != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, actual)
			}
		})
	}
}
