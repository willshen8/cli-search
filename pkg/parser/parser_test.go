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

func TestGetFileNamesInDirSuccess(t *testing.T) {
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

func TestGetFileNamesInDirFail(t *testing.T) {
	input := "../dummyFolder"
	expected := []string(nil)
	actual, err := GetFileNamesInDir(input)
	assert.NotNil(t, err)
	assert.Equal(t, expected, actual)
}

func TestReadJsonFile(t *testing.T) {
	var data = strings.NewReader(`helloWorld!`)
	expectedData := "helloWorld!"
	actual, err := ReadJsonFile(data)
	assert.Equal(t, expectedData, string(actual))
	assert.Equal(t, nil, err)
}

func TestUnmarshalData(t *testing.T) {
	data := []byte(`[{"_id": "101"}]`)
	actual, err := UnmarshalData(data)
	expected := []DataRecord{{"_id": "101"}}
	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func TestUnmarshalDataWithFailure(t *testing.T) {
	data := []byte(`[{"_id"}]`)
	_, err := UnmarshalData(data)
	assert.NotNil(t, err)
}
