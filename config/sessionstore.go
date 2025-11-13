package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

func CreateSessionStore() redis.Store {
	redisAddress := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sessionSecret := []byte(os.Getenv("SESSION_SECRET"))
	store, err := redis.NewStore(10, "tcp", redisAddress, "", "", sessionSecret)
	if err != nil {
		log.Fatal("Redis Store 初始化失敗", err)
	}

	return store
}
