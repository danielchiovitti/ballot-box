package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func NewConsumeOlapService(
	redisClient *redis.Client,
	config shared.ConfigInterface,
) ConsumeOlapServiceInterface {
	return &ConsumeOlapService{
		redisClient: redisClient,
		config:      config,
	}
}

type ConsumeOlapService struct {
	redisClient *redis.Client
	config      shared.ConfigInterface
}

func (c *ConsumeOlapService) Run() {
	fmt.Println("ConsumeOlapService Run")

	ctx := context.Background()
	for {
		messages, err := c.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.config.GetOlapStreamGroupName(),
			Consumer: "consumer-olap",
			Streams:  []string{c.config.GetOlapStreamName(), ">"},
			Count:    1,
			Block:    0,
		}).Result()

		if err != nil {
			log.Println(err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, stream := range messages {
			for _, message := range stream.Messages {
				if vt, exists := message.Values["vote"]; exists {
					var vote model.Vote
					_ = json.Unmarshal([]byte(vt.(string)), &vote)
					_ = c.redisClient.XAck(ctx, c.config.GetOlapStreamName(), c.config.GetOlapStreamGroupName(), message.ID).Err()
					fmt.Println("Olap Vote:", vote)
				}
			}
		}
	}

}
