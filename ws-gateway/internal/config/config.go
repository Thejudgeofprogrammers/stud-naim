package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	secret        string // Такой же как в gateway, нужно было логику разделить
	Exp           int
	Ref_time      int
}

func LoadEnv() *Config {
	rootDir, _ := os.Getwd()
	nameEnv := os.Getenv("CONFIG_FILE")
	log.Println("env:", nameEnv)
	if nameEnv == "" {
		nameEnv = ".env.dev"
	}
	path := filepath.Join(rootDir, nameEnv)

	if _, err := os.Stat(path); err != nil {
		log.Fatal("Error file env not exists")
	}

	err := godotenv.Load(path)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		Port:          getenv("PORT", "8001"),
		secret:        getenv("SECRET", "1984"),
		Exp:           getenvInt("EXPIRE_JWT", "600"),
		Ref_time:      getenvInt("REFRESH_TIME_JWT", "604800"),
	}

	return config
}

func getenvInt(k string, v string) int {
	e := os.Getenv(k)
	if e == "" {
		num, _ := strconv.Atoi(v)
		return num
	}
	num, err := strconv.Atoi(e)
	if err != nil {
		num, _ = strconv.Atoi(v)
		fmt.Printf("key: %s=%s not a number, default=%s", k, e, v)
		return num
	}
	return num
}

func getenv(k string, v string) string {
	e := os.Getenv(k)
	if e == "" {
		return v
	}
	return e
}

func (c *Config) GetSecret() string {
	return c.secret
}
