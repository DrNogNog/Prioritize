package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Message string `json:"message"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	var req ChatRequest
	json.NewDecoder(r.Body).Decode(&req)

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: "user",
					Content: req.Message,
				},
			},
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reply := resp.Choices[0].Message.Content

	json.NewEncoder(w).Encode(ChatResponse{
		Message: reply,
	})
}