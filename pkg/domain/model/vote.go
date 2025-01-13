package model

import "time"

type Vote struct {
	Wall      int       `json:"wall" bson:"wall,omitempty"`
	Candidate int       `json:"candidate" bson:"candidate,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}
