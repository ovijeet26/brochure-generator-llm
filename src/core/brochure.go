package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateBrochure(companyName, rootUrl string, links []LinkSuggestion) (string, error) {

	endpoint, model, apiKey, err := getLLMConfig()
	if err != nil {
		return "", err
	}

	userPrompt, err := GetBrochureUserPrompt(companyName, rootUrl, links)
	if err != nil {
		return "", err
	}
	requestBody := ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "system", Content: GetBrochureHumorousSystemPrompt()}, // use GetBrochureSystemPrompt() for a more formal brochure.
			{Role: "user", Content: userPrompt},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("LLM error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from model")
	}

	return result.Choices[0].Message.Content, nil
}
