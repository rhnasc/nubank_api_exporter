package nubank

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AccountWrapper struct {
	Account *Account `json:"account"`
}

type Account struct {
	Links map[string]Link `json:"_links"`

	PaymentMethod *PaymentMethod `json:"payment_method"`

	DueDay               int     `json:"due_day"`
	ProductID            string  `json:"product_id"`
	AvailableBalance     int     `json:"available_balance"`
	TemporaryLimitAmount int     `json:"temporary_limit_amount"`
	LimitRangeMin        int     `json:"limit_range_min"`
	FutureBalance        int     `json:"future_balance"`
	CurrentInterestRate  string  `json:"current_interest_rate"`
	PreciseCreditLimit   string  `json:"precise_credit_limit"`
	NextDueDate          string  `json:"next_due_date"`
	InterestRate         float64 `json:"interest_rate"`
	Status               string  `json:"status"`
	ID                   string  `json:"id"`

	Cards interface{}

	NetAvailable int `json:"net_available"`

	LimitRangeMax  int    `json"limit_range_max:"`
	NextCloseDate  string `json:"next_close_date"`
	CurrentBalance int    `json:"current_balance"`
	CreatedAt      string `json:"created_at"`
	RequestID      string `json:"request_id"`
	CustomerID     string `json:"customer_id"`
	CreditLimit    int    `json:"credit_limit"`
}

type PaymentMethod struct {
	AccountID string `json:"account_id"`
	Kind      string `json:"kind"`
}

type Balances struct {
	Future    int `json:"future"`
	Open      int `json:"open"`
	Due       int `json:"due"`
	Prepaid   int `json:"prepaid"`
	Available int `json:"available"`
}

func (c *NubankHttpClient) Account() (*Account, error) {
	creds := NewWebAppCredentials(c.login, c.password)

	marshaledCreds, err := json.Marshal(creds)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		c.accountHref,
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

	accountWrapper := &AccountWrapper{}
	err = json.Unmarshal(body, &accountWrapper)
	if err != nil {
		return nil, err
	}

	return accountWrapper.Account, err
}
