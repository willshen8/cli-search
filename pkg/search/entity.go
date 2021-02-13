package search

// entity defines the relationship of one table with other tables
type entity struct {
	foreignKeys []string
}

var organisationEnity = entity{foreignKeys: []string{"organization_id"}}
var userEntity = entity{foreignKeys: []string{"assignee_id", "submitter_id"}}
var ticketEntity = entity{foreignKeys: []string{}}
