package cmd

// Config holds configuration settings.
type Config struct {
	Http      bool // Whether to use HTTP
	WebSocket bool // Whether to use WebSockets
	Port      int  // Port number
}

// NewDefaultConfig creates a new Config instance with default values.
func NewDefaultConfig() *Config {
	return &Config{
		Http:      true,
		WebSocket: false,
		Port:      8080,
	}
}
