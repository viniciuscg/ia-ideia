package text

import (
	"strings"

	"google.golang.org/genai"
)

func GetText(result *genai.GenerateContentResponse) string {
	for _, cand := range result.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if part.Text != "" {
					return sanitizeJSON(part.Text)
				}
			}
		}
	}
	return ""
}

func sanitizeJSON(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}
