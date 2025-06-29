package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	fmt.Println("GOOGLE_API_KEY= =", os.Getenv("GOOGLE_API_KEY="))
	llm, err := googleai.New(context.Background(), googleai.WithDefaultModel("gemini-1.5-pro-latest"))
	if err != nil {
		log.Fatal("Error creating OpenAI LLM:", err)
	}

	chatmemory := memory.NewConversationBuffer()
	if err != nil {
		log.Fatal("Error creating chat memory:", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	for {
		log.Print("Enter a message (or type 'exit' to quit): ")
		if !scanner.Scan() {
			log.Println("Error reading input:", scanner.Err())
			continue
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			log.Println("Exiting...")
			break
		}

		response, err := llm.GenerateContent(ctx, []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, input),
		})
		if err != nil {
			log.Println("Error generating response:", err)
			continue
		}
		aiResponse := response.Choices[0].Content
		fmt.Printf("AI: %s\n\n", aiResponse)

		// Add the user input to the chat memory
		err = chatmemory.ChatHistory.AddUserMessage(ctx, input)
		if err != nil {
			log.Println("Error adding user message to chat memory:", err)
			continue
		}
		chatmemory.ChatHistory.AddAIMessage(ctx, aiResponse)
	}
	// Example usage of the LLM and chat memory
}
