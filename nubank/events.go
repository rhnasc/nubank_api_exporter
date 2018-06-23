package nubank

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type EventsWrapper struct {
	AsOf       time.Time `json:"as_of"`
	CustomerID string    `json:"customer_id"`

	Events []*Event `json:"events"`
}

type Event struct {
	Links map[string]Link `json:"_links"`

	Amount      int                    `json:"amount"`
	Category    string                 `json:"category"`
	Details     map[string]interface{} `json:"details"`
	Href        string                 `json:"href"`
	ID          string                 `json:"id"`
	Time        time.Time              `json:"time"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
}

func (c *NubankHttpClient) Events() ([]*Event, error) {
	creds := NewWebAppCredentials(c.login, c.password)

	marshaledCreds, err := json.Marshal(creds)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		c.eventsHref,
		bytes.NewReader(marshaledCreds),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add(
		"Authorization",
		"Bearer "+c.AuthToken,
	)
	applyWebAppHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	eventsWrapper := &EventsWrapper{}
	err = json.Unmarshal(body, &eventsWrapper)
	if err != nil {
		return nil, err
	}

	return eventsWrapper.Events, err
}
