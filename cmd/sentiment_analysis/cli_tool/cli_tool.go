package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nyeinsoe26/emoti-go/cmd/sentiment_analysis/config"
	"github.com/nyeinsoe26/emoti-go/pkg/models"
	"github.com/nyeinsoe26/emoti-go/pkg/sentiment_analysis"
	"github.com/nyeinsoe26/emoti-go/pkg/sentiment_analysis/llms"
)

type ReqInput struct {
	Texts           []string `json:"texts"`
	SentimentLabels []string `json:"sentiment_labels"`
}

func main() {
	// Command-line flag for YAML file input
	configFilepath := flag.String(
		"config",
		"configs/sentiment_analysis.yml",
		"Path to a YAML file with texts and optional sentiment categories.",
	)
	reqInputFilepath := flag.String(
		"input",
		"req.example.json",
		"Path to request file",
	)
	outputFilepath := flag.String(
		"output",
		"",
		"Path to the output JSON file. If not specified, prints to stdout.",
	)
	flag.Parse()
	if *configFilepath == "" {
		log.Fatalln("Invalid config filepath")
		flag.Usage()
		return
	}

	// Read the YAML file
	cfg, err := config.LoadConfig(*configFilepath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Read input
	inputContent, err := os.ReadFile(*reqInputFilepath)
	if err != nil {
		log.Fatalf("Error reading input file: %v \n", err)
	}
	var req ReqInput
	if err := json.Unmarshal(inputContent, &req); err != nil {
		log.Fatalf("Error parsing input JSON: %v \n", err)
	}

	var sentimentAnalyzer *sentiment_analysis.SentimentAnalyzer
	if cfg.UseLLM == llms.OpenAI {
		sentimentAnalyzer = sentiment_analysis.New(llms.OpenAI, cfg.LLMS.OpenAIConfig.ApiKey)
	}

	sentiments, err := sentimentAnalyzer.GetSentiments(models.Request{
		Texts:               req.Texts,
		Model:               cfg.LLMS.OpenAIConfig.Model,
		Temperature:         0,
		SentimentCategories: req.SentimentLabels,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	outputJSON, err := json.MarshalIndent(sentiments, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling output JSON: %v \n", err)
	}

	// Handle output
	if *outputFilepath != "" {
		// Write the output to a file
		err := os.WriteFile(*outputFilepath, outputJSON, 0644)
		if err != nil {
			log.Fatalf("Error writing to output file: %v \n", err)
		}
		log.Printf("Results written to %s\n", *outputFilepath)
	} else {
		// Print to stdout
		fmt.Println(string(outputJSON))
	}
}
