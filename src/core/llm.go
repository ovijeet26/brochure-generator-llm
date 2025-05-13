package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRelevantLinks(website *Website) (*LinkResponse, error) {

	endpoint, model, apiKey, err := getLLMConfig()
	if err != nil {
		return nil, err
	}

	requestBody := ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "system", Content: GetLinkSystemPrompt()},
			{Role: "user", Content: GetLinksUserPrompt(website)},
		},
		//ResponseFormat: "json", // Only supported by GPT-4
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("LLM error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned from model")
	}

	response := result.Choices[0].Message.Content
	var linkResp LinkResponse
	if err := json.Unmarshal([]byte(response), &linkResp); err != nil {
		return nil, err
	}

	return &linkResp, nil
}

func getLLMConfig() (string, string, string, error) {
	endpoint := os.Getenv("OPENAI_API_URL")
	model := os.Getenv("OPENAI_MODEL")
	apiKey := os.Getenv("OPENAI_API_KEY")

	if endpoint == "" || model == "" || apiKey == "" {
		return "", "", "", fmt.Errorf("missing one or more required env vars")
	}
	return endpoint, model, apiKey, nil
}
