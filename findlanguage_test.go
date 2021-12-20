package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var filepathExt = filepath.Ext

func TestFindLanguages(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	files := []string{exPath + "/findlanguage.go", exPath + "/run.sh", exPath + "/Dockerfile", exPath + "/findlanguage_test.go"}
	config := Config{}
	filepathExt = func(path string) string {
		return ".go0000"
	}
	config.FindLanguageRules = []Rule{
		{
			Language:  "go",
			Stratergy: "extension",
			Value:     ".go",
		},
		{
			Language:  "py",
			Stratergy: "extension",
			Value:     ".py",
		},
	}
	got := findLanguages(files, config)
	want := FindLanguageResult{
		Summary: map[string]float32{
			"go": 0.5,
		},
		Results: []Result{
			{
				Path:     "findlanguage.go",
				Language: "go",
			},
			{
				Path:     "findlanguage_test.go",
				Language: "go",
			},
		},
	}
	if len(got.Results) != len(want.Results) {
		t.Error("got wrong number of results for files wanted")
	}
	if got.Summary["go"] != want.Summary["go"] {
		t.Error("got wrong results for summary")
	}
}

func Test_UnknownExtensions_FindLanguages(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	files := []string{exPath + "/findlanguage.go1", exPath + "/run.sh1", exPath + "/Dockerfile1", exPath + "/findlanguage_test.go1"}
	config := Config{}
	filepathExt = func(path string) string {
		return ".go0000"
	}
	config.FindLanguageRules = []Rule{
		{
			Language:  "go",
			Stratergy: "extension",
			Value:     ".go",
		},
		{
			Language:  "py",
			Stratergy: "extension",
			Value:     ".py",
		},
	}
	got := findLanguages(files, config)
	want := FindLanguageResult{
		Summary: map[string]float32{},
		Results: []Result{},
	}
	if len(got.Results) != len(want.Results) {
		t.Error("got wrong number of results for files wanted")
	}
	if got.Summary["go"] != want.Summary["go"] {
		t.Error("got wrong results for summary")
	}
}
