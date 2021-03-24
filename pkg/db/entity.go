package db

// entity defines the relationship of one table with other tables
type entity struct {
	Table       string
	ForeignKeys []string
}

var OrganizationsEntity = entity{Table: "organizations", ForeignKeys: []string{"organization_id"}}
var UsersEntity = entity{Table: "users", ForeignKeys: []string{"assignee_id", "submitter_id"}}
