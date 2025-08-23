package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		hasError  bool
	}{
		{
			name:     "simple HTML body",
			inputURL: "https://blog.boot.dev",
			hasError: false,
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
			name:     "complete HTML",
			inputURL: "https://one.com",
			hasError: false,
			inputBody: `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Link Parser Test</title>
</head>
<body>
    <h1>Link Parser Test Page</h1>
    <p>Here are some links for testing:</p>

    <ul>
        <li><a href="https://example.com">Example Domain</a></li>
        <li><a href="https://openai.com/">OpenAI</a></li>
        <li><a href="/local/path/page.html">Local Relative Path</a></li>
        <li><a href="https://example.com?query=parser&test=1">Link with Query Parameters</a></li>
        <li><a href="https://example.org/download.zip" download>Download File</a></li>
    </ul>

		<h1>Example header</h1>
    <p>Example text.</p>
</body>
</html>
`,
			expected: []string{"https://example.com", "https://openai.com/", "https://one.com/local/path/page.html", "https://example.com?query=parser&test=1", "https://example.org/download.zip"},
		},
		{
			name:     "bad HTML",
			inputURL: "https://app.com",
			hasError: false,
			inputBody: `
<body>
	<a href="https://interpreterbook.com/"></a>
	<a href="https://compilerbook.com/"></a>
</body>
			`,
			expected: []string{"https://interpreterbook.com/", "https://compilerbook.com/"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualURLs, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !tc.hasError {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if err == nil && tc.hasError {
				t.Errorf("Test %v - '%s' FAIL: program succeded unexpectedly: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actualURLs, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actualURLs)
			}
		})
	}
}
