package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
)

// DataRecord represent a single data record read from the JSON file
type DataRecord map[string]interface{}

// getFileSuffix returns the suffix of a file name without extension, and is used for creating DB name
func GetFileSuffix(file string) string {
	var extension = filepath.Ext(file)
	return file[0 : len(file)-len(extension)]
}

// GetFileNamesInDir returns a list of all file in a directory specified
func GetFileNamesInDir(dir string) ([]string, error) {
	var fileNames []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fileNames, err
	}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames, nil
}

// ReadJsonFile reads contents from a reader and returns the data in bytes[]
func ReadJsonFile(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalData reads the data and convert the JSON format data into a slice of dataRecords
func UnmarshalData(data []byte) ([]DataRecord, error) {
	var marshalledData []DataRecord
	if err := json.Unmarshal(data, &marshalledData); err != nil {
		return marshalledData, err
	}
	return marshalledData, nil
}
