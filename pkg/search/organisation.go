package search

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Organisation struct {
	ID             uint32   `json:"_id"`
	Url            string   `json:"url"`
	External_Id    string   `json:"external_id"`
	Domain_names   []string `json:"domain_names"`
	Created_at     string   `json:"created_at"`
	Details        string   `json:"details"`
	Shared_tickets bool     `json:"shared_tickets"`
	Tags           []string `json:"tags"`
}

// OrgMap is used to efficiently check whether a particular field is part of the table
var OrgMap = map[string]bool{
	"_id":            true,
	"url":            true,
	"external_id":    true,
	"domain_names":   true,
	"created_at":     true,
	"details":        true,
	"shared_tickets": true,
	"tags":           true,
}

// ParseJsonOrgs will read user data from input and unmarshall them into Organisation struct
func ParseJsonOrgs(r io.Reader) ([]Organisation, error) {
	var organisations []Organisation
	organization_data, _ := ioutil.ReadAll(r) // swallow the error because it is already handled by kingpin library
	if err := json.Unmarshal(organization_data, &organisations); err != nil {
		return nil, err
	}
	return organisations, nil
}
