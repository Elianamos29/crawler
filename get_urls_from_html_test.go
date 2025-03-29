package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct{
		name      		string
		inputURL  		string
		inputBody 		string
		expected  		[]string
		errorContains 	string
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
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
				<span>Boot.dev</span>
		</a>
	</body>
</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
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
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
	`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
	`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
	`,
			expected: nil,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, err := url.Parse(tt.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tt.name, err)
				return
			}

			urls, err := getURLsFromHTML(tt.inputBody, baseURL)
			if err != nil {
				t.Errorf("an error occurred: %s", err)
			}

			if err != nil && !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tt.name, err)
				return
			} else if err != nil && tt.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tt.name, err)
				return
			} else if err == nil && tt.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tt.name, tt.errorContains)
				return
			}

			if !reflect.DeepEqual(urls, tt.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tt.name, tt.expected, urls)
				return
			}
		})
	}
}