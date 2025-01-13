package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/database/repository"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func NewConsumeOltpService(
	redisClient *redis.Client,
	config shared.ConfigInterface,
	voteRepository repository.VoteRepositoryInterface,
) ConsumeOltpServiceInterface {
	return &ConsumeOltpService{
		redisClient:    redisClient,
		config:         config,
		voteRepository: voteRepository,
	}
}

type ConsumeOltpService struct {
	redisClient    *redis.Client
	config         shared.ConfigInterface
	voteRepository repository.VoteRepositoryInterface
}

func (c *ConsumeOltpService) Run(id string) {
	fmt.Println("ConsumeOltpService Run")

	ctx := context.Background()
	for {
		messages, err := c.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.config.GetOltpStreamGroupName(),
			Consumer: fmt.Sprintf("consumer-oltp-%s", id),
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
					id, err := c.voteRepository.InsertOne(ctx, c.config.GetMongoDbDatabaseName(), "oltpcoll", vote)
					if err != nil {
						log.Println(err)
					}
					fmt.Println("Oltp VoteId:", id)
				}
			}
		}
	}

}
