package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/xpzouying/moonshot"
)

func main() {
	apiKey := os.Getenv("MOONSHOT_API_KEY")

	m := moonshot.New(apiKey)

	result, err := m.ListModels(context.Background())
	if err != nil {
		log.Fatalf("ListModels: error: %v", err)
	}

	models, _ := json.Marshal(result.ModelDetails)

	log.Printf("ListModels: %s", models)
}
