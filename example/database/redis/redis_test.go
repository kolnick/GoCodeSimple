package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var client redis.Client
var ctx = context.Background()

func init() {
	client = *redis.NewClient(&redis.Options{
		Addr: "192.168.1.5:6379",
		//		Addr: "localhost:6379",
	})
}

func TestPing(t *testing.T) {
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pong)
}

func TestSet(t *testing.T) {
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("err", err)
	}
}

func TestGet(t *testing.T) {
	result, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("key :", result)
}

// 带有过期时间的键值对
func TestSetKeyExpiryTTL(t *testing.T) {
	key := "key_expiry"
	value := "123456"
	duration := 10 * time.Second
	err := client.Set(ctx, key, value, duration).Err()
	if err != nil {
		fmt.Println("err", err)
	}
}

func TestGeyKeyExpiry(t *testing.T) {
	// 获取带有过期时间的键值对
	val_expiry, err := client.Get(ctx, "key_expiry").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key_expiry:", val_expiry)

	// 等待一段时间，以验证过期功能
	time.Sleep(11 * time.Second)

	// 再次获取带有过期时间的键值对
	val_expiry, err = client.Get(ctx, "key_expiry").Result()
	if err == redis.Nil {
		fmt.Println("key_expiry 已过期")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("key_expiry:", val_expiry)
	}
}
