package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VoteEntity struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Wall       int                `bson:"wall,omitempty"`
	Candidate  int                `bson:"candidate,omitempty"`
	BaseEntity `bson:",inline"`
}
