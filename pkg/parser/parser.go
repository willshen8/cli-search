package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

// ParseJsonToMapOfMap reads the data from a file and convert the result into
// a map of maps for efficient search
func ParseJsonToMapOfMap(r io.Reader) (map[string]map[string]interface{}, error) {
	var mapSlice []DataRecord
	result := make(map[string]map[string]interface{})
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &mapSlice); err != nil {
		return result, err
	}

	for _, v := range mapSlice {
		innerMap := make(map[string]interface{})
		for key, val := range v {
			innerMap[string(key)] = val
		}
		result[string(fmt.Sprintf("%v", v["_id"]))] = innerMap
	}
	return result, nil
}

// ParseFileAndStoreInDb opens a file and store the information into the database
func ParseFileAndStoreInDb(table string, file string, dB map[string]map[string]map[string]interface{}) (map[string]map[string]map[string]interface{}, error) {
	openedFile, err := os.Open(file)
	if err != nil {
		return dB, err
	}
	createdMap, err := ParseJsonToMapOfMap(openedFile)
	if err != nil {
		return dB, err
	}
	dB[table] = createdMap
	return dB, err
}
