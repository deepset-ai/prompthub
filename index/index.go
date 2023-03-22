package index

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type PromptIndex map[string]*Prompt

var data map[string]*Prompt

func init() {
	data, _ = readFiles("./prompts") // FIXME: get path from config
}

func GetPrompt(name string) (*Prompt, error) {
	val, ok := data[name]
	if ok {
		return val, nil
	}
	return nil, errors.New("not found")
}

func readFiles(path string) (PromptIndex, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	retval := PromptIndex{}

	for _, file := range files {
		var p Prompt
		yamlFile, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			log.Printf("Readfile error:  %v", err)
		}
		err = yaml.Unmarshal(yamlFile, &p)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		retval[p.Name] = &p
	}

	return retval, nil
}
