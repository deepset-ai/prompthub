package index

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type PromptIndex map[string]*Prompt

var prompts PromptIndex
var cards map[string]string

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

func GetCard(name string) (string, error) {
	val, ok := cards[name]
	if ok {
		return val, nil
	}
	return "", errors.New("not found")
}

func Init(path string) error {

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	prompts = PromptIndex{}
	cards = map[string]string{}

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

		// If the prompt has a related markdown with more detailed explanation
		// add it to a separate index. Note we use the prompt file name 
		// as the name of the card, not the actual name of the prompt.
		cardName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		cardFile := filepath.Join(path, fmt.Sprintf("%s.md", cardName))
		if _, err := os.Stat(cardFile); errors.Is(err, os.ErrNotExist) {
			log.Printf("Card not found for prompt %s", p.Name)
			continue
		}
		cardData, err := os.ReadFile(cardFile)
		if err != nil {
			log.Printf("Readfile error:  %v", err)
		}
		cards[p.Name] = string(cardData)
	}

	return nil
}
