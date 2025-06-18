package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel("gemma3:1b"))
	if err != nil {
		log.Fatal(err)
	}

	chatMemory := memory.NewConversationBuffer()
	fmt.Println("Chat Application Started! Type 'quit' to exit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" {
			break
		}
		response, err := llm.GenerateContent(ctx, []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, input),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.Choices[0].Content)
		chatMemory.ChatHistory.AddUserMessage(ctx, input)
		chatMemory.ChatHistory.AddAIMessage(ctx, response.Choices[0].Content)
	}

}
