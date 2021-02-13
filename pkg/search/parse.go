package search

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// ParseJsonToMapOfMap reads the data from a file and convert the result into
// a map of maps for efficient search
func ParseJsonToMapOfMap(r io.Reader) (map[string]map[string]interface{}, error) {
	var mapSlice []map[string]interface{}
	result := make(map[string]map[string]interface{}, len(mapSlice))
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
