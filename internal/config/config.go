package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct{
	Addr string `yaml:"address"`
} 

//env-default:"production"
// type Config struct{
// 	Env 		string `yaml:"env" env:"ENV" env-required:"true"`
// 	StoragePath string `yaml:"storage_path" env-required:"true"`
// 	HTTPServer	`yaml:"http_server"`
// }
type Config struct{
	Env 		string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	ProjectId	string`yaml:"project_id"`
	APIKey    string `yaml:"api_key" env-required:"true"`
	HTTPServer	`yaml:"http_server"`
}



func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("config","","Path to configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set.")
		}
	}

	if _,err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config file dose not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cant read config file: %s", err.Error())
	}

	return &cfg
}