package main

import (
	"fmt"
	"github.com/frtatmaca/go-speak/pkg"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("ASSEMBLY_AI_API_KEY")

	client := pkg.NewClient(apiKey)
	uploadUrl, err := client.UploadFile("test2.wav")
	if err != nil {
		log.Fatalf("Error While Uplaod File. %v", err)
	}

	id, err := client.StartTranscription(uploadUrl)
	if err != nil {
		log.Fatalf("Error While Start Transcription. %v", err)
	}

	text, err := client.GetTranscribedText(id)
	if err != nil {
		log.Fatalf("Error While Get Transcribed Text. %v", err)
	}

	fmt.Println(text)
}
