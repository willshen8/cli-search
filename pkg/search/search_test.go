package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSuccess(t *testing.T) {
	organisationMap, err := ParseJsonToMapOfMap(parseTestData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
}

func TestSearchByID(t *testing.T) {
	organisationMap, err := ParseJsonToMapOfMap(parseTestData)
	fmt.Println("organisationMap", organisationMap)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "_id", "101")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
}

func TestSearchWithoutSpecifiedValue(t *testing.T) {
	organisationMap, err := ParseJsonToMapOfMap(parseTestData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "external_id", "")
	var expectedResultIDs = []string{"101", "102"}
	assert.Equal(t, expectedResultIDs, actual)
	assert.Equal(t, nil, err)
}
