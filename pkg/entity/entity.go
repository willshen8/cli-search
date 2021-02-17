package entity

// entity defines the relationship of one table with other tables
type entity struct {
	Table       string
	ForeignKeys []string
}

var OrganisationEnity = entity{Table: "organisation", ForeignKeys: []string{"organization_id"}}
var UserEntity = entity{Table: "user", ForeignKeys: []string{"assignee_id", "submitter_id"}}
