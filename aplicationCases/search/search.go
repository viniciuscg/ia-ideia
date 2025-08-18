package search

import (
	"context"
	"encoding/json"
	"log"

	globalconfig "github.com/viniciuscg/ia-ideia.git/aplicationCases/globalConfig"
	"github.com/viniciuscg/ia-ideia.git/aplicationCases/prompt"
	"github.com/viniciuscg/ia-ideia.git/aplicationCases/text"
	"google.golang.org/genai"
)

type LastfmQuery struct {
	Type    string `json:"type"`
	Action  string `json:"action"`
	Time    string `json:"time"`
	User    string `json:"user"`
	Consult string `json:"consult"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Search struct {
	userPrompt string
}

func NewSearch(userPrompt string) *Search {
	return &Search{
		userPrompt: userPrompt,
	}
}

func (search Search) SearchPrompt() LastfmQuery {
	apiKey := globalconfig.LoadEnv("IAKEY")

	p := prompt.ConcatPrompts(search.userPrompt)
	parts := []*genai.Part{
		{Text: p},
	}

	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]*genai.Content{{Parts: parts}},
		nil,
	)

	response := text.GetText(result)

	var query LastfmQuery
	err := json.Unmarshal([]byte(response), &query)
	if err != nil {
		log.Println("Invalid JSON from model:", err)
		return LastfmQuery{}
	}

	return query
}
