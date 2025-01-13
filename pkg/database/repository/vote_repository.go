package repository

import (
	"github.com/danielchiovitti/ballot-box/pkg/database/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var voteRepositoryInstance *VoteRepository[entity.VoteEntity]
var lockVote sync.Mutex

type VoteRepository[T any] struct {
	client *mongo.Client
	BaseRepository[T]
}

func NewVoteRepository(client *mongo.Client) *VoteRepository[entity.VoteEntity] {
	if voteRepositoryInstance == nil {
		lockVote.Lock()
		defer lockVote.Unlock()
		if voteRepositoryInstance == nil {
			voteRepositoryInstance = &VoteRepository[entity.VoteEntity]{
				client: client,
				BaseRepository: BaseRepository[entity.VoteEntity]{
					Client: client,
				},
			}
		}
	}
	return voteRepositoryInstance
}
