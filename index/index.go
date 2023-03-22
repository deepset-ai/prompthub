package index

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type PromptIndex map[string]*Prompt

var data PromptIndex

func GetPrompt(name string) (*Prompt, error) {
	val, ok := data[name]
	if ok {
		return val, nil
	}
	return nil, errors.New("not found")
}

func Init(path string) error {

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	data = PromptIndex{}

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

		data[p.Name] = &p
	}

	return nil
}
