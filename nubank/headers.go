package nubank

import (
	"net/http"
)

func applyWebAppHeaders(req *http.Request) {
	header := req.Header

	header.Add("X-Correlation-Id", "WEB-APP.WnhxM")
	header.Add("Origin", "https://app.nubank.com.br")
	header.Add("Accept-Language", "en-US,en;q=0.9,pt;q=0.8,es;q=0.7")
	header.Add("Content-Type", "application/json;charset=UTF-8")
	header.Add("Accept", "application/json, text/plain, */*")
	header.Add("User-Agent", "https://github.com/rhnasc/nubank_api_exporter")
}
