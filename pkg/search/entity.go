package search

// entity defines the relationship of one table with other tables
type entity struct {
	table       string
	foreignKeys []string
}

var organisationEnity = entity{table: "organisation", foreignKeys: []string{"organization_id"}}
var userEntity = entity{table: "user", foreignKeys: []string{"assignee_id", "submitter_id"}}
