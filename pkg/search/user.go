package search

// UserMap is used to efficiently check whether a particular field is part of the table
var UserMap = map[string]bool{
	"_id":             true,
	"url":             true,
	"external_id":     true,
	"name":            true,
	"alias":           true,
	"created_at":      true,
	"active":          true,
	"verified":        true,
	"shared":          true,
	"locale":          true,
	"timezone":        true,
	"last_login_at":   true,
	"email":           true,
	"phone":           true,
	"signature":       true,
	"organization_id": true,
	"tags":            true,
	"suspended":       true,
	"role":            true,
}

var UserFields = []string{"_id", "url", "external_id", "name", "alias", "created_at", "active",
	"verified", "shared", "locale", "timezone", "last_login_at", "email", "phone", "signature", "organization_id",
	"tags", "suspended", "role"}
