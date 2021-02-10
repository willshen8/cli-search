package search

import "time"

type Ticket struct {
	ID              uint32    `json:"ticket_id"`
	Url             string    `json:"url"`
	External_Id     string    `json:"external_id"`
	Created_at      time.Time `json:"created_at"`
	Type            string    `json:"type"`
	Subject         string    `json:"subject"`
	Description     string    `json:"description"`
	Priority        string    `json:"priority"`
	Status          string    `json:"status"`
	Submitter_id    uint32    `json:"submitter_id"`
	Assignee_id     uint32    `json:"assignee_id"`
	Organization_id uint32    `json:"organization_id"`
	Tags            []string  `json:"tags"`
	Has_incident    bool      `json:"has_incident"`
	Due_at          string    `json:"due_at"`
	Via             string    `json:"via"`
}
