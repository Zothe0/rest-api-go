package main

import (
	"flag"
	"log"

	"github.com/Zothe0/rest-api-go/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.json", "path to config file")
}

func main() {
	flag.Parse()

	config, err := apiserver.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
