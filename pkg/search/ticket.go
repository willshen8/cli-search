package search

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
