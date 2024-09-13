package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

var (
	booksFile = "books.json"
	mu        sync.Mutex
)

func readJSONFile(filename string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty data
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func writeJSONFile(filename string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}
