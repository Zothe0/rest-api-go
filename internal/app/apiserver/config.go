package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Zothe0/rest-api-go/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr string `json:"bindAddr"`
	LogLevel string `json:"logLevel"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig(path string) (*Config, error) {
	filename, err := os.Open("configs/apiserver.json")
	if err != nil {
		return nil, err
	}
	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		return nil, err
	}

	parseData := &Config{
		// Default config, rewrites by configs/apiserver.json
		BindAddr: ":3000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}

	err = json.Unmarshal(data, parseData)
	if err != nil {
		return nil, err
	}
	return parseData, nil
}
