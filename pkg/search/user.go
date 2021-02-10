package search

import "time"

type User struct {
	ID              uint32    `json:"user_id"`
	Url             string    `json:"url"`
	External_Id     string    `json:"external_id"`
	Name            string    `json:"name"`
	Alias           string    `json:"alias"`
	Created_at      time.Time `json:"created_at"`
	Active          bool      `json:"active"`
	Verified        bool      `json:"verified"`
	Shared          bool      `json:"shared"`
	Locale          string    `json:"locale"`
	Timezone        string    `json:"timezone"`
	Last_login_at   time.Time `json:"last_login_at"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Signature       string    `json:"signature"`
	Organisation_id uint32    `json:"organization_id"`
	Tags            []string  `json:"tags"`
	Suspended       bool      `json:"suspended"`
	Role            string    `json:"role"`
}
