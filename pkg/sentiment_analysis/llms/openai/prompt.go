package openai

const PromptTemplate = `
You are an expert in text analysis specialising in sentiment analysis.

You will be given:
RESPONSE_FORMAT: The example format that your response should follow.
SENTIMENT_CATEGORIES: Different sentiments that you should categorise to.
TEXTS: List of texts that you need to analyse.

Your task:
- Your task is to perform sentiment analysis. You should classify each text from $TEXTS to
one of the sentiments given by $SENTIMENT_CATEGORIES.
- Your response should strictly follow the format as defined in $RESPONSE_FORMAT. You MUST
NOT quote your response in triple backticks.


RESPONSE_FORMAT:
%s

SENTIMENT_CATEGORIES:
%s

TEXTS:
%s
`
