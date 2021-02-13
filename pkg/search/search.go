package search

import (
	"sort"
)

// search takes any table, a field name and a value and returns a slice of IDs of the records found
func Search(table map[string]map[string]interface{}, field string, value string) ([]string, error) {
	if _, found := OrgMap[field]; !found {
		return nil, ErrInvalidSearchField
	}
	var result []string
	if value == "" {
		for k := range table {
			result = append(result, string(k))
		}
		sort.Strings(result)
		return result, nil
	}

	for k, v := range table {
		if v[field] == value {
			result = append(result, string(k))
		}
	}
	sort.Strings(result)
	return result, nil
}
