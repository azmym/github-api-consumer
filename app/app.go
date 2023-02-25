package app

import (
	"fmt"
	"github-api-consumer/config"
	"github-api-consumer/httpclient"
	"github.com/kelseyhightower/envconfig"
	"net/http"
)

func loadConfig() error {
	err := envconfig.Process("", config.EnvVariable)
	if err != nil {
		return err
	}
	return nil
}

func StartUp() {
	//load the configuration
	err := loadConfig()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	//
	client := &httpclient.GithubClient{
		Client: http.DefaultClient,
	}
	repos, err := client.GetRepoLanguages("azmym", "Build-RESTful-Spring")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%+v", repos)
}
