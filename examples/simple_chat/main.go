package main

import (
	"context"
	"log"
	"os"

	"github.com/xpzouying/moonshot"
)

func main() {
	apiKey := os.Getenv("MOONSHOT_API_KEY")

	m := moonshot.New(apiKey)

	req := &moonshot.RequestCompletionChat{
		Model: moonshot.ModelV18K,
		Messages: []moonshot.Message{
			{
				Role:    moonshot.RoleSystem,
				Content: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一些涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。",
			},
			{
				Role:    moonshot.RoleUser,
				Content: "你好，我叫李雷，1+1等于多少？",
			},
		},
		Temperature: 0.3,
	}

	result, err := m.CreateChatCompletions(context.Background(), req)
	if err != nil {
		log.Fatalf("CreateChatCompletions: error: %v", err)
	}

	log.Printf("CreateChatCompletions: result: %v", result.Choices[0].Message.Content)
}
