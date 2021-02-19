package copy

import (
	"fmt"
	"io/ioutil"

	"github.com/willshen8/cli-search/internal/errors"
)

// CopyFile copy the srcFile to the destFile
func CopyFile(srcFile string, destFile string) error {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return errors.NewError(err, fmt.Sprintf("Error reading file %s", srcFile))
	}
	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		return err
	}
	return nil
}
