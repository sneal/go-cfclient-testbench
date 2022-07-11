package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient"
	"net/url"
	"os"
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
	config, err := cfclient.NewConfigFromCF()
	if err != nil {
		return err
	}
	client, err := cfclient.NewClient(config)
	if err != nil {
		return err
	}

	vars, err := client.ListV3SecGroupsByQuery(url.Values{})
	bytes, err := json.Marshal(vars)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
