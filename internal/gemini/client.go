package gemini

import (
	"chatbot/config"
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetChatResponse(prompt string) (string, error) {
	apiKey := config.LoadAPIKey()

	// Create a new Gemini AI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %v", err)
	}
	defer client.Close()

	// Get the model
	model := client.GenerativeModel("gemini-pro")

	// Generate response
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("Gemini API error: %v", err)
	}

	// Ensure response has candidates
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "No response from AI.", nil
	}

	var responseText string
	for _, part := range resp.Candidates[0].Content.Parts {
		switch p := part.(type) {
		case genai.Text:
			responseText += string(p)
		default:
			return "Unexpected part type in response.", nil
		}
	}

	if responseText == "" {
		return "AI did not return any text.", nil
	}

	return responseText, nil
}
