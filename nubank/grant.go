package nubank

type Grant struct {
	Links map[string]Link `json:"_links"`

	AccessToken string `json:"access_token"`
	TokenType   string `json:"bearer"`

	RefreshToken  string `json:"refresh_token"`
	RefreshBefore string `json:"refresh_before"`
}
