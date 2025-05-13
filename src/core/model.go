package core

type Website struct {
	URL   string
	Title string
	Text  string
	Links []string
}

type LinkSuggestion struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type LinkResponse struct {
	Links []LinkSuggestion `json:"links"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model          string        `json:"model"`
	Messages       []ChatMessage `json:"messages"`
	ResponseFormat string        `json:"response_format,omitempty"` // Optional for Ollama
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
