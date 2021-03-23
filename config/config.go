package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func readYamlConfig() {
	// Config
	viper.SetConfigName("default") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Set default value
	viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

	// Declare var
	env := viper.GetString("app.env")
	producerbroker := viper.GetString("app.producerbroker")
	consumerbroker := viper.GetString("app.consumerbroker")
	linetoken := viper.GetString("app.linetoken")

	// Print
	fmt.Println("---------- Example ----------")
	fmt.Println("app.env :", env)
	fmt.Println("app.producerbroker :", producerbroker)
	fmt.Println("app.consumerbroker :", consumerbroker)
	fmt.Println("app.linetoken :", linetoken)

}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
