package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Structure Declarations
type Result struct {
	Path     string `json:"path"`
	Language string `json:"language"`
}

type FindLanguageResult struct {
	Summary map[string]float32 `json:"summary"`
	Results []Result           `json:"results"`
}

type Rule struct {
	Language  string `json:"language"`
	Stratergy string `json:"stratergy"`
	Value     string `json:"value"`
}

type Config struct {
	FindLanguageRules []Rule `json:"rules"`
}

const (
	CONFIG_FILE_PATH = "/go/src/github.com/CodeQualityProject/Task2/config.json"
	ROOT_FOLDER      = "/app"
	OUTPUT_FILE_PATH = "/app/findlanguageoutput"
)

func main() {

	var (
		files []string
		err   error
	)

	files, err = fetchFiles(ROOT_FOLDER)
	if err != nil {
		panic(err)
	}

	config, err := fetchConfig()
	if err != nil {
		log.Fatal("error fetching config")
		panic(err)
	}

	findLanguageResult := findLanguages(files, config)

	err = writeOutput(findLanguageResult)
	if err != nil {
		log.Fatal("Error writing result")
		panic(err)
	}
}

func fetchConfigFile() (*os.File, error) {
	jsonFile, err := os.Open(CONFIG_FILE_PATH)
	if err != nil {
		fmt.Printf("failed to open json file: config.json, error: %v", err)
		return nil, err
	}
	return jsonFile, nil
}

func fetchConfig() (Config, error) {
	data := Config{}
	configJsonFile, err := fetchConfigFile()
	if err != nil {
		log.Fatal("Error reading config")
		panic(err)
	}
	jsonData, err := ioutil.ReadAll(configJsonFile)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return data, err
	}
	return data, nil
}

func fetchFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func findLanguages(files []string, config Config) FindLanguageResult {
	var findLanguageResult FindLanguageResult
	countBuckets := make(map[string]int)
	findLanguageResult.Results = make([]Result, 0)
	for _, file := range files {
		var result Result
		fileExtension := filepath.Ext(file)
		filename := filepath.Base(file)
		for _, rule := range config.FindLanguageRules {
			if rule.Stratergy == "file_name" && filename == rule.Value {
				result.Path = file
				result.Language = rule.Language
				findLanguageResult.Results = append(findLanguageResult.Results, result)
				countBuckets[result.Language]++

			} else if rule.Stratergy == "extension" && fileExtension == rule.Value {
				result.Path = file
				result.Language = rule.Language
				findLanguageResult.Results = append(findLanguageResult.Results, result)
				countBuckets[result.Language]++
			} else {
				continue
			}
		}
		findLanguageResult.Summary = map[string]float32{}
		for k, v := range countBuckets {
			findLanguageResult.Summary[k] = (float32(v) / float32(len(files)))

		}
	}
	return findLanguageResult
}

func writeOutput(findLanguageResult FindLanguageResult) error {
	resJSON, err := json.MarshalIndent(findLanguageResult, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(OUTPUT_FILE_PATH, resJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
