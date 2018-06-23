package nubank

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Credentials struct {
	GrantType    string `json:"grant_type"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

const (
	WebAppGrantType    = "password"
	WebAppClientID     = "other.conta"
	WebAppClientSecret = "yQPeLzoHuJzlMMSAjC-LgNUJdUecx8XO"
)

func NewWebAppCredentials(login, password string) *Credentials {
	return &Credentials{
		GrantType:    WebAppGrantType,
		Login:        login,
		Password:     password,
		ClientID:     WebAppClientID,
		ClientSecret: WebAppClientSecret,
	}
}

func (c *NubankHttpClient) Authenticate() error {
	creds := NewWebAppCredentials(c.login, c.password)

	marshaledCreds, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.authHref,
		bytes.NewReader(marshaledCreds),
	)
	if err != nil {
		return err
	}

	applyWebAppHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	grant := &Grant{}
	err = json.Unmarshal(body, &grant)
	if err != nil {
		return err
	}

	c.AuthToken = grant.AccessToken

	c.accountHref = grant.Links["account"].Href
	c.eventsHref = grant.Links["events"].Href

	return err
}
