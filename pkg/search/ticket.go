package search

// TicketMap is used to efficiently check whether a particular field is part of the table
var TicketMap = map[string]bool{
	"_id":             true,
	"url":             true,
	"external_id":     true,
	"created_at":      true,
	"type":            true,
	"subject":         true,
	"description":     true,
	"priority":        true,
	"status":          true,
	"submitter_id":    true,
	"assignee_id":     true,
	"organization_id": true,
	"tags":            true,
	"has_incidents":   true,
	"due_at":          true,
	"via":             true,
}
