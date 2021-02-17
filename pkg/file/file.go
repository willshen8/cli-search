package file

import (
	"io/ioutil"
)

// CopyFile copy the srcFile to the destFile
func CopyFile(srcFile string, destFile string) error {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		return err
	}
	return nil
}
