package cli

import (
	openai "github.com/sashabaranov/go-openai"
)

type oaiClients struct {
	openAIClient openai.Client
}
