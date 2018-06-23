package nubank

type NubankHttpClient struct {
	login    string
	password string

	AuthToken string

	authHref string

	eventsHref  string
	accountHref string
}

func NewNubankHttpClient(login, password string) *NubankHttpClient {
	return &NubankHttpClient{
		login:    login,
		password: password,
	}
}
