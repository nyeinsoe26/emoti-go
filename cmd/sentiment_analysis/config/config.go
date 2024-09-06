package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type OpenAIConfig struct {
	Model  string `yaml:"model"`
	ApiKey string `yaml:"api_key"`
}

type LLMS struct {
	OpenAIConfig OpenAIConfig `yaml:"openai"`
}

type Config struct {
	LLMS   LLMS   `yaml:"llms"`
	UseLLM string `yaml:"use_llm"`
}

// LoadConfig loads the configuration from a YAML file.
func LoadConfig(filePath string) (*Config, error) {
	// Read the YAML file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the YAML into the Config struct
	var cfg Config
	err = yaml.Unmarshal(fileContent, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.UseLLM == "" {
		cfg.UseLLM = "openai"
	}

	if cfg.UseLLM == "openai" {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatalln("OPENAI_API_KEY must be provided")
		} else {
			cfg.LLMS.OpenAIConfig.ApiKey = apiKey
		}

	}
	return &cfg, nil
}
