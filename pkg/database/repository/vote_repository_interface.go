package repository

import "github.com/danielchiovitti/ballot-box/pkg/database/entity"

type VoteRepositoryInterface interface {
	BaseRepositoryInterface[entity.VoteEntity]
}
