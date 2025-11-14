package utils

import (
	"os"
	"encoding/json"
)

type Config struct {
	Crypto string `json:"crypto_type"`
	Address string `json:"address"`
}

var Cfg Config

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if CheckError(err) { return err }
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Cfg)
	if CheckError(err) { return err }

	return nil
}