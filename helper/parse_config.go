package helper

import (
	"log"
	"os"

	"github.com/gopher5889/interview/config"
	"github.com/spf13/viper"
)

func New(path string) *config.Config {
	if _, err := os.Stat(path); err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	config := &config.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
