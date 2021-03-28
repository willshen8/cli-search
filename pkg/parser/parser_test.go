package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileSuffix(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected string
	}{
		{
			testName: "Valid file name",
			input:    "hello.json",
			expected: "hello",
		},
		{
			testName: "File without an extension",
			input:    "file",
			expected: "file",
		},
		{
			testName: "File has no extension name",
			input:    "file.",
			expected: "file",
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual := GetFileSuffix(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetFileNamesInDir(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected []string
	}{
		{
			testName: "Current directory",
			input:    ".",
			expected: []string{"parser.go", "parser_test.go"},
		},
		{
			testName: "One directory up",
			input:    "../print",
			expected: []string{"print.go", "print_test.go"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual, err := GetFileNamesInDir(tc.input)
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestReadJsonFile(t *testing.T) {
	var data = strings.NewReader(`helloWorld!`)
	expectedData := "helloWorld!"
	actual, err := ReadJsonFile(data)
	assert.Equal(t, expectedData, string(actual))
	assert.Equal(t, nil, err)
}
