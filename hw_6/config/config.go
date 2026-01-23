package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port      string
	JWTTTL    time.Duration
	SecretKey []byte
	DB        DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func (cfg *DBConfig) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

func GetConfig() *AppConfig {
	godotenv.Load(".env", "../.env")

	jwtTTL, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		log.Fatalf("Invalid JWT_TTL: %v", err)
	}

	secretKey := decodeBase64(os.Getenv("SECRET_KEY"))

	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("Invalid POSTGRES_PORT: %v", err)
	}

	cfg := &AppConfig{
		Port:      os.Getenv("PORT"),
		JWTTTL:    jwtTTL,
		SecretKey: secretKey,
		DB: DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     dbPort,
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		panic(err)
	}

	return cfg
}

func decodeBase64(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("Failed to decode base64: %v", err)
	}
	return data
}
