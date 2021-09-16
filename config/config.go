package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	EmailTemplatePath    string `mapstructure:"EMAIL_TEMPLATE_PATH"`
	CustomerDataFilePath string `mapstructure:"CUSTOMER_DATA_FILE_PATH"`
	OutputEmailPath      string `mapstructure:"OUTPUT_EMAIL_PATH"`
	ErrorPath            string `mapstructure:"ERROR_PATH"`
}

func New(path string) *Config {
	if _, err := os.Stat(path); err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	config := &Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
