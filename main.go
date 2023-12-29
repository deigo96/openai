package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  apiKey := os.Getenv("API_KEY")

	fmt.Print("Please enter something: ")

	scanner.Scan()
	userInput := scanner.Text()

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content:userInput,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
