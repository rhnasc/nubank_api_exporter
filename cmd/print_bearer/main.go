package main

import (
	"fmt"
	"os"

	"github.com/nubank_api_exporter/nubank"
)

func main() {
	login := os.Getenv("NUBANK_LOGIN")
	password := os.Getenv("NUBANK_PASSWORD")

	client := nubank.NewNubankHttpClient(login, password)

	err := client.Discover()
	must(err)

	err = client.Authenticate()
	must(err)

	fmt.Println(client.AuthToken)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
