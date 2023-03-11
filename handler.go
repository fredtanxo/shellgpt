package main

import (
	"errors"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
)

// HandleResponse handles the stream response and then apply the handler to the partial content
func HandleResponse(stream *openai.ChatCompletionStream, handler func(string)) {
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
