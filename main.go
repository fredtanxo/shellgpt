package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
)

// The OpenAI API Key
var apiKey string

// Recent AI response content
var output string

func main() {
	flag.StringVar(&apiKey, "api-key", "", "OpenAI API key")
	flag.Parse()

	if apiKey == "" {
		log.Fatalln("No API key provided")
	}

	client := openai.NewClient(apiKey)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[0;32m> \033[0m")

		// Read user input
		input, err := getUserInput(scanner)
		if err != nil {
			continue
		}

		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue
		}

		// Generate request data
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: input,
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

		handleResponse(stream, func(s string) {
			fmt.Print(s)
		})

		// Release resources
		fmt.Println()
		stream.Close()
	}
}

// handleResponse handles the stream response
// and then apply the handler to the partial content
func handleResponse(stream *openai.ChatCompletionStream, handler func(string)) {
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalln("error receiving stream response", err)
		}

		c := response.Choices[0].Delta.Content

		output += c

		if err != nil {
			log.Fatal("error rendering response into markdown format", err)
		}

		handler(c)
	}
}

// getUserInput read the string from the scanner.
// Multiple lines are separated by a backslash(\) at the end of the line.
func getUserInput(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		t := scanner.Text()
		s := strings.TrimSuffix(t, "\\")
		if strings.HasSuffix(t, "\\") {
			n, err := getUserInput(scanner)
			if err != nil {
				return "", err
			}
			return strings.Join([]string{s, n}, "\n"), nil
		}
		return s, nil
	}
	return "", scanner.Err()
}
