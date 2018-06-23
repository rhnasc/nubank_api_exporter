package nubank

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	EntryPoint = "https://prod-s0-webapp-proxy.nubank.com.br/api/discovery"
)

type DiscoveryRoutes struct {
	Login string `json:"login"`
}

func (c *NubankHttpClient) Discover() error {
	req, err := http.NewRequest(
		http.MethodGet,
		EntryPoint,
		nil,
	)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	discoveryRoutes := &DiscoveryRoutes{}
	err = json.Unmarshal(body, &discoveryRoutes)
	if err != nil {
		return err
	}

	c.authHref = discoveryRoutes.Login

	return nil
}
