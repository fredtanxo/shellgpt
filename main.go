package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strings"
)

const (
	promptDefault = "\u001b[32m> \u001b[0m"
	promptMulLine = "\u001b[32m. \u001b[0m"
	promptExit    = "\nBye"
)

// The OpenAI API client params
var (
	apiUrl string
	apiKey string
)

// Recent AI response content
var output string

func main() {
	flag.StringVar(&apiUrl, "api-url", "https://api.openai.com/v1", "OpenAI API base url")
	flag.StringVar(&apiKey, "api-key", "", "OpenAI API key")
	flag.Parse()

	envKey := os.Getenv("OPENAI_API_KEY")
	if len(apiKey) == 0 {
		apiKey = envKey
	}

	if apiKey == "" {
		log.Fatalln("No API key provided")
	}

	ShowBanner()

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = apiUrl
	client := openai.NewClientWithConfig(config)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          promptDefault,
		InterruptPrompt: promptExit,
		EOFPrompt:       promptExit,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = rl.Close() }()

	var inputs []string
	for {
		// Read user input
		input, err := rl.Readline()
		if err != nil {
			break
		}

		input = strings.TrimSpace(input)
		if len(input) == 0 && len(inputs) == 0 {
			continue
		}

		inputs = append(inputs, input)
		if strings.HasSuffix(input, "\\") {
			rl.SetPrompt(promptMulLine)
			continue
		}

		// Generate request data
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: strings.Join(inputs, "\n"),
			},
		}

		if len(output) > 0 {
			c := openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: output,
			}
			messages = append(messages, c)
		}

		// Send request
		stream, err := client.CreateChatCompletionStream(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)
		if err != nil {
			log.Fatalln("error creating stream request", err)
		}

		HandleResponse(stream, func(s string) {
			fmt.Print(s)
		})

		// Cleanup resources
		stream.Close()
		fmt.Println()
		inputs = inputs[:0]
		rl.SetPrompt(promptDefault)
	}
}
