package model

import "time"

type Vote struct {
	Wall      int       `json:"wall"`
	Candidate int       `json:"candidate"`
	CreatedAt time.Time `json:"created_at"`
}
