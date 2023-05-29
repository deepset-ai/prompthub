package index

// Prompt data model.
type Prompt struct {
	Name        string                 `json:"name"`
	Tags        []string               `json:"tags,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Version     string                 `json:"version"`
	Text        string                 `json:"text"`
	Description string                 `json:"description"`
}
