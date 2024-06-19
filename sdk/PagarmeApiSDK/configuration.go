package pagarmeapisdk

import (
	"os"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
	serviceRefererName string
	environment        Environment
	httpConfiguration  HttpConfiguration
	basicAuthUserName  string
	basicAuthPassword  string
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
	config := Configuration{}

	for _, option := range options {
		option(&config)
	}
	return config
}

// WithServiceRefererName is an option that sets the ServiceRefererName in the Configuration.
func WithServiceRefererName(serviceRefererName string) ConfigurationOptions {
	return func(c *Configuration) {
		c.serviceRefererName = serviceRefererName
	}
}

// WithEnvironment is an option that sets the Environment in the Configuration.
func WithEnvironment(environment Environment) ConfigurationOptions {
	return func(c *Configuration) {
		c.environment = environment
	}
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
	}
}

// WithBasicAuthUserName is an option that sets the BasicAuthUserName in the Configuration.
func WithBasicAuthUserName(basicAuthUserName string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthUserName = basicAuthUserName
	}
}

// WithBasicAuthPassword is an option that sets the BasicAuthPassword in the Configuration.
func WithBasicAuthPassword(basicAuthPassword string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthPassword = basicAuthPassword
	}
}

// ServiceRefererName returns the serviceRefererName from the Configuration.
func (c *Configuration) ServiceRefererName() string {
	return c.serviceRefererName
}

// Environment returns the Environment from the Configuration.
func (c *Configuration) Environment() Environment {
	return c.environment
}

// HttpConfiguration returns the HttpConfiguration from the Configuration.
func (c *Configuration) HttpConfiguration() HttpConfiguration {
	return c.httpConfiguration
}

// BasicAuthUserName returns the BasicAuthUserName from the Configuration.
func (c *Configuration) BasicAuthUserName() string {
	return c.basicAuthUserName
}

// BasicAuthPassword returns the BasicAuthPassword from the Configuration.
func (c *Configuration) BasicAuthPassword() string {
	return c.basicAuthPassword
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	serviceRefererName := os.Getenv("PAGARMEAPISDK_SERVICE_REFERER_NAME")
	if serviceRefererName != "" {
		config.serviceRefererName = serviceRefererName
	}
	environment := os.Getenv("PAGARMEAPISDK_ENVIRONMENT")
	if environment != "" {
		config.environment = Environment(environment)
	}
	basicAuthUserName := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_USER_NAME")
	if basicAuthUserName != "" {
		config.basicAuthUserName = basicAuthUserName
	}
	basicAuthPassword := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_PASSWORD")
	if basicAuthPassword != "" {
		config.basicAuthPassword = basicAuthPassword
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

// Server represents available servers.
type Server string

const (
	ENUMDEFAULT Server = "default"
)

// Environment represents available environments.
type Environment string

const (
	PRODUCTION Environment = "production"
)

// CreateRetryConfiguration creates a new RetryConfiguration with the provided options.
func CreateRetryConfiguration(options ...RetryConfigurationOptions) RetryConfiguration {
	config := DefaultRetryConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateHttpConfiguration creates a new HttpConfiguration with the provided options.
func CreateHttpConfiguration(options ...HttpConfigurationOptions) HttpConfiguration {
	config := DefaultHttpConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateConfiguration creates a new Configuration with the provided options.
func CreateConfiguration(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}
