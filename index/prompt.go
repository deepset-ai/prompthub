package index

// Prompt data model.
type Prompt struct {
	Name        string                 `json:"name"`
	Tags        []string               `json:"tags"`
	Meta        map[string]interface{} `json:"meta"`
	Version     string                 `json:"version"`
	Text        string                 `json:"prompt_text"`
	Description string                 `json:"description"`
}
