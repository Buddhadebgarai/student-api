package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env:"HTTP_ADDR" env-default:":8080"`
}

type Config struct {
	HTTPServer  HTTPServer `yaml:"http_server"`
	Env         string     `yaml:"env" env:"ENV" env-default:"prod"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-default:"./data"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" { // check if CONFIG_PATH is set, if not, try to get it from flages from command line
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}
	var cfg Config

	// Read config from file
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}
	return &cfg
}
