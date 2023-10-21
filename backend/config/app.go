package config

import (
	"io/ioutil"
	"log"

	"github.com/abdulsalam/pixel8labs/internal/entity"
	"gopkg.in/yaml.v2"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func LoadOAuthConfig(config entity.Config) oauth2.Config {
	var oauthConfig = oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user", "username", "repo", "gist"},
	}

	return oauthConfig
}

func LoadAppConfig(configFile string) (entity.Config, error) {
	// Create an instance of the struct.
	var (
		config entity.Config
		err    error
	)

	// Read the YAML file into a byte slice.
	data, err := ioutil.ReadFile(configFile + "/app.yaml")
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
		return config, err
	}

	// Unmarshal the YAML data into the struct.
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
		return config, err
	}

	return config, nil
}
