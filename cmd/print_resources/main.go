package main

import (
	"fmt"
	"os"

	"github.com/rhnasc/nubank_api_exporter/nubank"
)

func main() {
	login := os.Getenv("NUBANK_LOGIN")
	password := os.Getenv("NUBANK_PASSWORD")

	client := nubank.NewNubankHttpClient(login, password)

	err := client.Discover()
	must(err)

	err = client.Authenticate()
	must(err)

	account, err := client.Account()
	must(err)

	events, err := client.Events()
	must(err)

	fmt.Println(account)
	fmt.Println(events[0])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
