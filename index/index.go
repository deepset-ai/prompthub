package index

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type PromptIndex map[string]*Prompt

var prompts PromptIndex

func GetPrompt(name string) (*Prompt, error) {
	val, ok := prompts[name]
	if ok {
		return val, nil
	}
	return nil, errors.New("not found")
}

func GetPrompts() []*Prompt {
	v := make([]*Prompt, 0, len(prompts))

	for _, value := range prompts {
		v = append(v, value)
	}

	return v
}

func Init(path string) error {

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	prompts = PromptIndex{}

	for _, file := range files {
		// yaml files only
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		var p Prompt
		yamlFile, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			log.Printf("Readfile error:  %v", err)
		}
		err = yaml.Unmarshal(yamlFile, &p)
		if err != nil {
			log.Fatalf("Unmarshal %s: %v", file.Name(), err)
		}

		prompts[p.Name] = &p
	}

	return nil
}
