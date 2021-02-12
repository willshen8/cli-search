package search

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Ticket struct {
	ID              string   `json:"_id"`
	Url             string   `json:"url"`
	External_Id     string   `json:"external_id"`
	Created_at      string   `json:"created_at"`
	Type            string   `json:"type"`
	Subject         string   `json:"subject"`
	Description     string   `json:"description"`
	Priority        string   `json:"priority"`
	Status          string   `json:"status"`
	Submitter_id    uint32   `json:"submitter_id"`
	Assignee_id     uint32   `json:"assignee_id"`
	Organization_id uint32   `json:"organization_id"`
	Tags            []string `json:"tags"`
	Has_incident    bool     `json:"has_incident"`
	Due_at          string   `json:"due_at"`
	Via             string   `json:"via"`
}

// TicketMap is used to efficiently check whether a particular field is part of the table
var TicketMap = map[string]bool{
	"_id":             true,
	"url":             true,
	"external_id":     true,
	"created_at":      true,
	"type":            true,
	"active":          true,
	"verified":        true,
	"shared":          true,
	"locale":          true,
	"timezone":        true,
	"last_login_at":   true,
	"email":           true,
	"phone":           true,
	"signature":       true,
	"organisation_id": true,
	"tags":            true,
	"suspended":       true,
	"role":            true,
}

// ParseJsonOrgs will read user data from input and unmarshall them into Organisation struct
func ParseJsonTickets(r io.Reader) ([]Ticket, error) {
	var tickets []Ticket
	tickets_data, _ := ioutil.ReadAll(r) // swallow the error because it is already handled by kingpin library
	if err := json.Unmarshal(tickets_data, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}
