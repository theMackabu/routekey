package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	_ "embed"
	"os"
	"errors"
)

//go:embed config.example
var data []byte

type Config struct {
	Production bool   `json:"production"`
	Port       string `json:"port"`
	Words      []string `json:"words"`
}

func ReadConfig() Config {
	var cfg Config
	
	if _, err := os.Stat("./config.json"); errors.Is(err, os.ErrNotExist) {
	  err := ioutil.WriteFile("./config.json", data, 0644)
	  
		if err != nil {
			 log.Fatal(err)
		}
	}
	
	fileData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Not able to read file")
	}
	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}
