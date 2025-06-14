package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetChatGPTResponse(ctx context.Context, messages []map[string]string) (string, error) {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"
	body := map[string]interface{}{
		"model":       "gpt-4o-mini",
		"messages":    messages,
		"temperature": 0.7,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", fmt.Errorf("OpenAI API error: %s", string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", err
	}

	choicesRaw, ok := result["choices"].([]interface{})
	if !ok || len(choicesRaw) == 0 {
		return "", errors.New("no choices returned by OpenAI")
	}

	choice, ok := choicesRaw[0].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid choice format from OpenAI")
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid message format from OpenAI")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", errors.New("missing content in OpenAI response")
	}

	return content, nil
}
