package main

import (
	"reflect"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "no scheme",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "empty path",
			inputURL: "https://example.com",
			expected: "example.com/",
		},
		{
			name: "trailing slash",
			inputURL: "https://example.com/",
			expected: "example.com/",
		},
		{
			name: "with query",
			inputURL: "https://example.com/path?query=1",
			expected: "example.com/path",
		},
		{
			name: "with fragment",
			inputURL: "https://example.com/path#section",
			expected: "example.com/path",
		},
		{
			name: "empty string",
			inputURL: "",
			expected: "",
		},
		{
			name: "malformed URL",
			inputURL: "https://%%%",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if tc.expected != "" && err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if tc.expected == "" && err == nil {
				t.Errorf("Expected error for %q, got none", tc.inputURL)
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct{
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urls, err := getURLsFromHTML(tt.inputBody, tt.inputURL)
			if err != nil {
				t.Errorf("an error occurred: %s", err)
			}

			if !reflect.DeepEqual(tt.expected, urls) {
				t.Errorf("expected urls: %+v, got %+v", tt.expected, urls)
			}
		})
	}
}