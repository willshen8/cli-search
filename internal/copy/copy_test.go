package copy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFileFail(t *testing.T) {
	var dummySrcFile = "/hello.txt"
	var dummyDestFile = "/world.txt"
	err := CopyFile(dummySrcFile, dummyDestFile)
	assert.NotNil(t, err)
}

func TestCopyFileWithWriteFailure(t *testing.T) {
	var dummySrcFile = "../../config/organizations.json"
	var dummyDestFile = "/world.txt"
	err := CopyFile(dummySrcFile, dummyDestFile)
	assert.NotNil(t, err)
}
