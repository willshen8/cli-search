package search

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type User struct {
	ID              uint32   `json:"_id"`
	Url             string   `json:"url"`
	External_Id     string   `json:"external_id"`
	Name            string   `json:"name"`
	Alias           string   `json:"alias"`
	Created_at      string   `json:"created_at"`
	Active          bool     `json:"active"`
	Verified        bool     `json:"verified"`
	Shared          bool     `json:"shared"`
	Locale          string   `json:"locale"`
	Timezone        string   `json:"timezone"`
	Last_login_at   string   `json:"last_login_at"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Signature       string   `json:"signature"`
	Organisation_id uint32   `json:"organization_id"`
	Tags            []string `json:"tags"`
	Suspended       bool     `json:"suspended"`
	Role            string   `json:"role"`
}

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
	"organisation_id": true,
	"tags":            true,
	"suspended":       true,
	"role":            true,
}

// ParseJsonUsers will read user data from input and unmarshall them into User struct
func ParseJsonUsers(r io.Reader) ([]User, error) {
	var users []User
	user_data, _ := ioutil.ReadAll(r) // swallow the error because it is already handled by kingpin library
	if err := json.Unmarshal(user_data, &users); err != nil {
		return nil, err
	}
	return users, nil
}
