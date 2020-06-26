package sensu

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
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

	// namespace is the sensu namespace.
	namespace string

	// Tokens
	tokens *types.Tokens

	// insecureSkipTLSVerify will allow self-signed certificates.
	insecureSkipTLSVerify bool

	// trustedCAFile is the path to a trusted CA.
	trustedCAFile string

	// client is the sensu client
	client *client.RestClient
}

// LoadAndValidate is a method used to initiate a client.
func (c *Config) LoadAndValidate() error {
	c.client = client.New(c)

	// Configure the client for SSL.
	tlsConfig := tls.Config{}

	if c.TrustedCAFile() != "" {
		caCertPool, err := corev2.LoadCACerts(c.TrustedCAFile())
		if err != nil {
			log.Printf("[DEBUG] Problem loading CA Certs: %s", err)
		}
		tlsConfig.RootCAs = caCertPool
	}

	tlsConfig.InsecureSkipVerify = c.InsecureSkipTLSVerify()

	tlsConfig.BuildNameToCertificate()
	tlsConfig.CipherSuites = corev2.DefaultCipherSuites

	c.client.SetTLSClientConfig(&tlsConfig)

	// Create an access token.
	tokens, err := c.client.CreateAccessToken(c.apiUrl, c.username, c.password)
	if err != nil {
		return err
	}

	if tokens == nil {
		return fmt.Errorf("bad username or password")
	}

	// Save the access token in the internal Sensu client.
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

// Namespace implements the Namespace method for the config.Config interface.
func (c *Config) Namespace() string {
	if c.namespace == "" {
		return config.DefaultNamespace
	}

	return c.namespace
}

// SaveNamespace implements the SaveNamespace method for the config.Config interface.
func (c *Config) SaveNamespace(namespace string) error {
	c.namespace = namespace
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

// Tokens implements the Tokens method for the config.Config interface.
func (c *Config) Tokens() *types.Tokens {
	return c.tokens
}

// SaveTokens implements the SaveTokens method for the config.Config interface.
func (c *Config) SaveTokens(tokens *types.Tokens) error {
	c.tokens = tokens
	return nil
}

// InsecureSkipTLSVerify implements the InsecureSkipTLSVerify method for the
// config.Config interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) InsecureSkipTLSVerify() bool {
	return c.insecureSkipTLSVerify
}

// SaveInsecureSkipTLSVerify implements the SaveInsecureSkipTLSVerify method
// for the config.Config interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) SaveInsecureSkipTLSVerify(v bool) error {
	c.insecureSkipTLSVerify = v
	return nil
}

// TrustedCAFile implements the TrustedCAFile method for the config.Config
// interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) TrustedCAFile() string {
	return c.trustedCAFile
}

// SaveTrustedCAFile implements the SaveTrustedCAFile method for the
// config.Config interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) SaveTrustedCAFile(v string) error {
	c.trustedCAFile = v
	return nil
}

// Timeout implements the Timeout method for the config.Config
// interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) Timeout() time.Duration {
	return 0
}

// SaveTimeout implements the SaveTimeout method for the
// config.Config interface.
//
// Note: this is not actually implemented in this provider. This is only
// used to satisfy the Sensu client interface.
func (c *Config) SaveTimeout(val time.Duration) error {
	return nil
}

// determineNamespace will figure out the right namespace setting to use.
func (c *Config) determineNamespace(d *schema.ResourceData) string {
	if v, ok := d.Get("namespace").(string); ok && v != "" {
		return v
	}

	return c.Namespace()
}
