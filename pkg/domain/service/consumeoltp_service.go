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

func NewConsumeOltpService(
	redisClient *redis.Client,
	config shared.ConfigInterface,
) ConsumeOltpServiceInterface {
	return &ConsumeOltpService{
		redisClient: redisClient,
		config:      config,
	}
}

type ConsumeOltpService struct {
	redisClient *redis.Client
	config      shared.ConfigInterface
}

func (c *ConsumeOltpService) Run() {
	fmt.Println("ConsumeOltpService Run")

	ctx := context.Background()
	for {
		messages, err := c.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.config.GetOltpStreamGroupName(),
			Consumer: "consumer-oltp",
			Streams:  []string{c.config.GetOltpStreamName(), ">"},
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
					_ = c.redisClient.XAck(ctx, c.config.GetOltpStreamName(), c.config.GetOltpStreamGroupName(), message.ID).Err()
					fmt.Println("Oltp Vote:", vote)
				}
			}
		}
	}

}
