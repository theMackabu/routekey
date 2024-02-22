package config

import (
	_ "embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

//go:embed config.example
var data []byte

type Config struct {
	Production bool     `json:"production"`
	Port       string   `json:"port"`
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

// -------- don't really feel like dealing with theMackabu's code rn - blobbybilb, made functions to keep config.example and config.json synced, and changes change both --------

func ReadExampleConfig() Config {
	var cfg Config

	data, err := ioutil.ReadFile("./config/config.example")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return cfg
}

func SaveConfig(cfg Config) {
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("./config/config.example", b, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func CopyExampleConfigToConfig() {
	data, err := ioutil.ReadFile("./config/config.example")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("config.json", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func AddWord(word string) {
	cfg := ReadExampleConfig()

	cfg.Words = append(cfg.Words, word)

	SaveConfig(cfg)

	CopyExampleConfigToConfig()
}

func RemoveWord(word string) {
	cfg := ReadExampleConfig()

	for i, w := range cfg.Words {
		if w == word {
			cfg.Words = append(cfg.Words[:i], cfg.Words[i+1:]...)
			break
		}
	}

	SaveConfig(cfg)

	CopyExampleConfigToConfig()
}

func ReadWords() []string {
	return ReadExampleConfig().Words
}
