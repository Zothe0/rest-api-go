package store

// Config ...
type Config struct {
	DatabaseURL string `json:"databaseURL"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
