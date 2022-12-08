package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

var rdb *redis.Client
var ctx = context.TODO()

type RPIP struct {
	Token string `json:"token"`
}

func Init() {
	if rdb != nil {
		rdbErr := rdb.Ping(ctx)
		if rdbErr != nil {
			ConnectRedis()
		}
	} else {
		ConnectRedis()
	}
}

func ConnectRedis() {
	RedisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       RedisDB,
	})
}

func SetAccess(UserId int, Username string, ExpiresAt time.Time, Token string) error {
	Init()
	at := ExpiresAt
	now := time.Now()

	errAccess := rdb.Set(ctx, Username+":"+Token, UserId, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}

func GetAccess(Username string, Token string) error {
	Init()
	_, err := rdb.Get(ctx, Username+":"+Token).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InvokeAccess(Username string, Token string) error {
	Init()
	iter := rdb.Scan(ctx, 0, Username+":"+Token, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		rdb.Del(ctx, key)
	}
	return nil
}

func InvokeAllAccess(Username string) error {
	Init()
	iter := rdb.Scan(ctx, 0, Username+":*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		rdb.Del(ctx, key)
	}
	return nil
}

func Close() {
	err := rdb.Close()
	if err != nil {
		return
	}
}
