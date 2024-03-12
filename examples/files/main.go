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

	ctx := context.Background()

	{
		fileInfo, err := m.UploadFile(ctx, "./hello.txt")
		if err != nil {
			log.Fatalf("UploadFile: error: %v", err)
		}

		logJSONStr(fileInfo)
	}

	{

		fileInfos, err := m.ListFiles(ctx)
		if err != nil {
			log.Fatalf("ListModels: error: %v", err)
		}

		logJSONStr(fileInfos)
	}
}

func logJSONStr(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	log.Printf("%s", data)
}
