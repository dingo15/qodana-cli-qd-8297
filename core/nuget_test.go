package core

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTempFileWithContent(content string) string {
	tempFile, _ := os.CreateTemp("", "test")
	_, _ = tempFile.WriteString(content)
	err := tempFile.Close()
	if err != nil {
		return ""
	}
	return tempFile.Name()
}

func TestCheckForPrivateFeed(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "FileWithPrivateFeed",
			filename: createTempFileWithContent(`<add key= 'xxx' value = 'http://'>`),
			expected: true,
		},
		{
			name:     "FileWithPrivateFeed2",
			filename: createTempFileWithContent(`<add key= 'xxx' value = 'https://'>`),
			expected: true,
		},
		{
			name:     "FileWithoutPrivateFeed",
			filename: createTempFileWithContent(`<add key= 'xxx' value='yyy'>`),
			expected: false,
		},
		{
			name:     "EmptyFile",
			filename: createTempFileWithContent(""),
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := checkForPrivateFeed(test.filename)
			if result != test.expected {
				t.Errorf("got/want mismatch, got %v, want %v", result, test.expected)
			}
			_ = os.Remove(test.filename)
		})
	}
}

func TestPrepareNugetConfig(t *testing.T) {
	_ = os.Setenv(qodanaNugetName, "qdn")
	_ = os.Setenv(qodanaNugetUrl, "test_url")
	_ = os.Setenv(qodanaNugetUser, "test_user")
	_ = os.Setenv(qodanaNugetPassword, "test_password")

	// create temp dir
	tmpDir, _ := os.MkdirTemp("", "test")
	defer func(tmpDir string) {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Fatal(err)
		}
	}(tmpDir)

	prepareNugetConfig(tmpDir)

	expected := `<?xml version="1.0" encoding="utf-8"?>
<configuration>
  <packageSources>
    <clear />
    <add key="nuget.org" value="https://api.nuget.org/v3/index.json" />
    <add key="qdn" value="test_url" />
  </packageSources>
  <packageSourceCredentials>
    <qdn>
      <add key="Username" value="test_user" />
      <add key="ClearTextPassword" value="test_password" />
    </qdn>
  </packageSourceCredentials>
</configuration>`

	file, err := os.Open(filepath.Join(tmpDir, ".nuget", "NuGet", "NuGet.Config"))
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	text = strings.TrimSuffix(text, "\n")
	if text != expected {
		t.Fatalf("got:\n%s\n\nwant:\n%s", text, expected)
	}
}
