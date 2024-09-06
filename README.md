# About
A simple program to perform sentiment analysis. The main purpose is to experiment with OpenAI in golang.

# Usage
You simply need to provide the input texts and the target labels you want to classify to.

## As module
```
sentimentAnalyzer = sentiment_analysis.New(llms.OpenAI, cfg.LLMS.OpenAIConfig.ApiKey)

sentiments, _ := sentimentAnalyzer.GetSentiments(models.Request{
		Texts:               <Some texts u wanna classify>,
		Model:               cfg.LLMS.OpenAIConfig.Model,
		Temperature:         0,
		SentimentCategories: <The unique labels u wanna classify to>,
	})
```

## As a cli tool
To use this as a cli tool, you need to define your job as a json file. Your job file should be in the following format:
```
{
    "texts": [
        <your texts>
    ],
    "sentiment_labels": [
        <your labels>
    ]
}
```
Refer to `req.example.json` for reference.

Use the following command to run:
```
OPENAI_API_KEY="Bring your own key" go run cmd/sentiment_analysis/cli_tool/cli_tool.go -config=configs/sentiment_analysis.yml
```