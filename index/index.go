package index

func GetPrompt(name string) (*Prompt, error) {
	return &Prompt{Name: name}, nil
}
