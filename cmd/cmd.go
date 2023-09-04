package cmd

import (
	"flag"
)

// ParseConfigFromArgs parses command line arguments and returns a Config instance.
func ParseConfigFromArgs() *Config {
	config := NewDefaultConfig()

	// Define command line flags
	httpFlag := flag.Bool("http", config.Http, "Use HTTP")
	webSocketFlag := flag.Bool("ws", config.WebSocket, "Use WebSockets")
	portFlag := flag.Int("port", config.Port, "Port number")

	flag.Parse()

	config.Http = *httpFlag
	config.WebSocket = *webSocketFlag
	config.Port = *portFlag

	return config
}

func ParseRPCConfigFromArgs() *Config {
	config := RPCDefaultConfig()

	// Define command line flags
	httpFlag := flag.Bool("http", config.Http, "Use HTTP")
	webSocketFlag := flag.Bool("ws", config.WebSocket, "Use WebSockets")
	portFlag := flag.Int("port", config.Port, "Port number")

	flag.Parse()

	config.Http = *httpFlag
	config.WebSocket = *webSocketFlag
	config.Port = *portFlag

	return config
}
