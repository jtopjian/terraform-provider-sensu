package sensu

import (
	"fmt"

	"github.com/sensu/sensu-go/cli/client"
	"github.com/sensu/sensu-go/cli/client/config"
	"github.com/sensu/sensu-go/types"
)

// Config represents a configuration struct for sensu.
type Config struct {
	// apiUrl is the URL to the Sensu API service.
	apiUrl string

	// username is the username.
	username string

	// password is the password.
	password string

	// edition is the sensu edition.
	edition string

	// environment is the sensu environment.
	environment string

	// organization is the sensu organization.
	organization string

	// Tokens
	tokens *types.Tokens

	// client is the sensu client
	client *client.RestClient
}

// LoadAndValidate is a method used to initiate a client.
func (c *Config) LoadAndValidate() error {
	c.client = client.New(c)
	tokens, err := c.client.CreateAccessToken(c.apiUrl, c.username, c.password)
	if err != nil {
		return err
	}

	if tokens == nil {
		return fmt.Errorf("bad username or password")
	}

	err = c.SaveTokens(tokens)
	if err != nil {
		return err
	}

	return nil
}

// The following methods are to implement the Sensu Config interface.

// APIUrl implements APIUrl method for the config.Config interface.
func (c *Config) APIUrl() string {
	return c.apiUrl
}

// SaveAPIUrl implements the SaveAPIUrl method for the config.Config interface.
func (c *Config) SaveAPIUrl(url string) error {
	c.apiUrl = url
	return nil
}

// Edition implements the Edition method for the config.Config interface.
func (c *Config) Edition() string {
	if c.edition == "" {
		return config.DefaultEdition
	}

	return c.edition
}

// SaveEdition implements the SaveEdition method for the config.Config interface.
func (c *Config) SaveEdition(edition string) error {
	c.edition = edition
	return nil
}

// Environment implements the Environment method for the config.Config interface.
func (c *Config) Environment() string {
	if c.environment == "" {
		return config.DefaultEnvironment
	}

	return c.environment
}

// SaveEnvironment implements the SaveEnvironment method for the config.Config interface.
func (c *Config) SaveEnvironment(environment string) error {
	c.environment = environment
	return nil
}

// Format implements the Format method of the config.Config interface.
func (c *Config) Format() string {
	return config.DefaultFormat
}

// SaveFormat implements the SaveFormat method of the config.Config interface.
func (c *Config) SaveFormat(format string) error {
	return nil
}

// Organization implements the Organization method for the config.Config interface.
func (c *Config) Organization() string {
	if c.organization == "" {
		return config.DefaultOrganization
	}

	return c.organization
}

// SaveOrganization implements the SaveOrganization method for the config.Config interface.
func (c *Config) SaveOrganization(organization string) error {
	c.organization = organization
	return nil
}

// Tokens implements the Tokens method for the config.Config interface.
func (c *Config) Tokens() *types.Tokens {
	return c.tokens
}

// SaveTokens implements the SaveTokens method for the config.Config interface.
func (c *Config) SaveTokens(tokens *types.Tokens) error {
	c.tokens = tokens
	return nil
}
