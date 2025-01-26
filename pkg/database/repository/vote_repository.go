package repository

import (
	"github.com/danielchiovitti/ballot-box/pkg/database/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var voteRepositoryInstance *VoteRepository
var lockVote sync.Mutex

type VoteRepository struct {
	client *mongo.Client
	BaseRepository[entity.VoteEntity]
}

func NewVoteRepository(client *mongo.Client) VoteRepositoryInterface {
	if voteRepositoryInstance == nil {
		lockVote.Lock()
		defer lockVote.Unlock()
		if voteRepositoryInstance == nil {
			voteRepositoryInstance = &VoteRepository{
				client: client,
				BaseRepository: BaseRepository[entity.VoteEntity]{
					Client: client,
				},
			}
		}
	}
	return voteRepositoryInstance
}
