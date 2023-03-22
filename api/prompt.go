package api

import (
	"net/http"

	"github.com/deepset-ai/prompthub/index"
)

// PromptResponse is the response payload for the Prompt data model.
type PromptResponse struct {
	*index.Prompt
}

func (rd *PromptResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewPromptResponse(prompt *index.Prompt) *PromptResponse {
	resp := &PromptResponse{Prompt: prompt}

	return resp
}
