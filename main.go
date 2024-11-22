// main.go
package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"context"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Example of setting a value
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}

	// Example of getting a value
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Println("key:", val)

	// Example of using a stream
	streamName := "mystream"
	err = rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{"field1": "value1"},
	}).Err()
	if err != nil {
		fmt.Println("Error adding to stream:", err)
		return
	}

	// Reading from the stream
	streams, err := rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{streamName, "0"},
		Count:   1,
		Block:   0,
	}).Result()
	if err != nil {
		fmt.Println("Error reading from stream:", err)
		return
	}
	for _, stream := range streams {
		for _, message := range stream.Messages {
			fmt.Println("Stream message:", message.Values)
		}
	}
}

