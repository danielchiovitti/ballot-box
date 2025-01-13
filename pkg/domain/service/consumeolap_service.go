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

func NewConsumeOlapService(
	redisClient *redis.Client,
	config shared.ConfigInterface,
	voteRepository repository.VoteRepositoryInterface,
) ConsumeOlapServiceInterface {
	return &ConsumeOlapService{
		redisClient:    redisClient,
		config:         config,
		voteRepository: voteRepository,
	}
}

type ConsumeOlapService struct {
	redisClient    *redis.Client
	config         shared.ConfigInterface
	voteRepository repository.VoteRepositoryInterface
}

func (c *ConsumeOlapService) Run(id string) {
	fmt.Println("ConsumeOlapService Run")

	ctx := context.Background()
	for {
		messages, err := c.redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.config.GetOlapStreamGroupName(),
			Consumer: fmt.Sprintf("consumer-olap-%s", id),
			Streams:  []string{c.config.GetOlapStreamName(), ">"},
			Count:    20,
			Block:    2000,
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

					_, err := c.voteRepository.InsertOne(ctx, c.config.GetMongoDbDatabaseName(), "olapcoll", vote)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}

}
