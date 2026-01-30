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
	Port         string             `validate:"required"`
	DB           DBConfig           `validate:"required"`
	RefreshToken RefreshTokenConfig `validate:"required"`
	JWT          JWTConfig          `validate:"required"`
}

type DBConfig struct {
	Host     string `validate:"required"`
	Port     int    `validate:"gt=0"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
}

type RefreshTokenConfig struct {
	Size int           `validate:"gt=0"`
	Type string        `validate:"required"`
	TTL  time.Duration `validate:"gt=0"`
}

type JWTConfig struct {
	SecretKey []byte        `validate:"required"`
	TTL       time.Duration `validate:"gt=0"`
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

	refresh_token_size, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_SIZE"))
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_SIZE: %v", err)
	}

	refresh_token_ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_TTL: %v", err)
	}

	cfg := &AppConfig{
		Port: os.Getenv("PORT"),
		JWT: JWTConfig{
			TTL:       jwtTTL,
			SecretKey: secretKey,
		},
		DB: DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     dbPort,
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
		RefreshToken: RefreshTokenConfig{
			Size: refresh_token_size,
			Type: os.Getenv("REFRESH_TOKEN_TYPE"),
			TTL:  refresh_token_ttl,
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
