package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
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

	chain := chains.NewConversation(llm, chatMemory)
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
		response, err := chains.Run(ctx, chain, input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("AI: %s\n\n", response)
	}

}
