package config

type TransmissionServer struct {
	Host     string
	Port     uint16
	Protocol string
	Username string
	Password string
}

type RestServer struct {
	Port        int
	ApiBasePath string
	MainEntity  string
}

type Log struct {
	Level string
}

type Config struct {
	TransmissionServer TransmissionServer
	RestServer         RestServer
	Log                Log
}

// TODO: unmarshal config from JSON
func Load() *Config {
	config := &Config{
		TransmissionServer: TransmissionServer{
			Host:     "127.0.0.1",
			Port:     9091,
			Protocol: "https",
			Username: "John",
			Password: "Smith",
		},
		RestServer: RestServer{
			Port:        5000,
			ApiBasePath: "api",
			MainEntity:  "torrents",
		},
		Log: Log{
			Level: "INFO",
		},
	}
	return config
}
