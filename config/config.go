package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	ListenAddr    string   `json:"listen_address"`
	BannedDomains []string `json:"banned_domains"`
	UpstreamAddr  string   `json:"upstreamAddr"`
}

func LoadConfig(configpath string) (*Config, error) {

	file, err := os.Open(configpath)

	if err != nil {
		fmt.Printf("Error in opening condfig file %s\n", err)
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("Error in reading condfig file %s\n", err)
		return nil, err
	}

	var conf *Config

	err = json.Unmarshal(bytes, &conf)

	if err != nil {
		fmt.Printf("Error unmarshalling json: %v", err)
		return nil, err
	}

	return conf, nil

}
