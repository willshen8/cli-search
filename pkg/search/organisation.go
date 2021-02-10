package search

type Organisation struct {
	ID             uint32   `json:"organisation_id"`
	Url            string   `json:"url"`
	External_Id    string   `json:"external_id"`
	Domain_names   []string `json:"domain_names"`
	Created_at     string   `json:"created_at"`
	Details        string   `json:"details"`
	Shared_tickets bool     `json:"shared_tickets"`
	Tags           []string `json:"tags"`
}
