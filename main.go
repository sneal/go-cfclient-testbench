package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfclient/v3/client"
)

func main() {
	fmt.Println("Starting test...")
	err := execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func execute() error {
	c, err := client.NewConfigFromCFHome()
	if err != nil {
		return err
	}
	c.SkipSSLValidation(true)
	cf, err := client.New(c)
	if err != nil {
		return err
	}

	_, err = cf.Applications.Start("6a62ec20-738e-498b-bcde-89bba7ea8c96")
	return err
}
