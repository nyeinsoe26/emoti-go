package openai

const ResponseFormat string = `
{
	"sentiments": [
		<sentiment of text_1>,
		...
		<sentiment of text_N>
	]
}
`

var DefaultSentiments = []string{
	"Extremely Positive",
	"Positive",
	"Mixed",
	"Negative",
	"Extremely Negative",
}
