package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MYSQL        MySQLConfig
	JWT          JWTConfig
	SERVER       ServerConfig
	CLOUDSTORAGE CloudStorageConfig
	MIDTRANS     MidtransConfig
}

type (
	MySQLConfig struct {
		MYSQL_USER string
		MYSQL_PASS string
		MYSQL_HOST string
		MYSQL_PORT string
		MYSQL_NAME string
	}

	ServerConfig struct {
		SERVER_HOST string
		SERVER_PORT string
	}

	JWTConfig struct {
		JWT_SECRET string
	}

	CloudStorageConfig struct {
		AWS_ACCESS_KEY_ID     string
		AWS_SECRET_ACCESS_KEY string
		AWS_REGION            string
		AWS_BUCKET_NAME       string
	}

	MidtransConfig struct {
		MIDTRANS_SERVER_KEY string
		MIDTRANS_CLIENT_KEY string
	}
)

func LoadConfig() (*Configuration, error) {

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	} else {
		fmt.Println("warning: .env file not found. make sure environment variables are set")
	}

	return &Configuration{
		MYSQL: MySQLConfig{
			MYSQL_USER: os.Getenv("MYSQL_USER"),
			MYSQL_PASS: os.Getenv("MYSQL_PASS"),
			MYSQL_HOST: os.Getenv("MYSQL_HOST"),
			MYSQL_PORT: os.Getenv("MYSQL_PORT"),
			MYSQL_NAME: os.Getenv("MYSQL_NAME"),
		},
		SERVER: ServerConfig{
			SERVER_HOST: os.Getenv("SERVER_HOST"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
		JWT: JWTConfig{
			JWT_SECRET: os.Getenv("JWT_SECRET"),
		},
		CLOUDSTORAGE: CloudStorageConfig{
			AWS_ACCESS_KEY_ID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			AWS_SECRET_ACCESS_KEY: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			AWS_REGION:            os.Getenv("AWS_REGION"),
			AWS_BUCKET_NAME:       os.Getenv("AWS_BUCKET_NAME"),
		},
		MIDTRANS: MidtransConfig{
			MIDTRANS_SERVER_KEY: os.Getenv("MIDTRANS_SERVER_KEY"),
			MIDTRANS_CLIENT_KEY: os.Getenv("MIDTRANS_CLIENT_KEY"),
		},
	}, nil
}
