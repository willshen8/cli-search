package search

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
