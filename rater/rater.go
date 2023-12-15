package rater

import (
	"encoding/json"
	"os"
)

const fileName = "./rater/rater.json"

func writeMapToFile(data map[string]int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func readMapFromFile() (map[string]int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var result map[string]int
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func IncrementWordScore(word string) error {
	data, err := readMapFromFile()
	if err != nil {
		return err
	}

	data[word]++

	err = writeMapToFile(data)
	if err != nil {
		return err
	}

	return nil
}

func GetWordScore(word string) (int, error) {
	data, err := readMapFromFile()
	if err != nil {
		return 0, err
	}

	return data[word], nil
}

func DecrementWordScore(word string) error {
	data, err := readMapFromFile()
	if err != nil {
		return err
	}

	data[word]--

	err = writeMapToFile(data)
	if err != nil {
		return err
	}

	return nil
}
