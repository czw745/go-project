package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

var Client *redis.Client

//REDISConfig ...
type REDISConfig struct {
	Host string
	Port string
}

//BuildREDISConfig ... build redis config
func BuildREDISConfig() *REDISConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisConfig := REDISConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
	return &redisConfig
}

//REDISURL ... redis url
func REDISURL(redis *REDISConfig) string {
	return fmt.Sprintf("%s:%s", redis.Host, redis.Port)
}

//REDISConnection ... redis connection
func REDISConnection() {
	//Initializing redis
	Client = redis.NewClient(&redis.Options{
		Addr: REDISURL(BuildREDISConfig()), //redis port
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
