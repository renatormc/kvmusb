package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	ExeDir    string
	Port      string
	SecretKey []byte
}

func GetConfig() Config {
	return config
}

func InitWithExeDir(exeDir string) {
	config.ExeDir = exeDir
	envFile := filepath.Join(config.ExeDir, ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Port = os.Getenv("PORT")
	config.SecretKey = []byte(os.Getenv("SECRETKEY"))

}

func Init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	InitWithExeDir(filepath.Dir(ex))
}
