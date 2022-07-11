package model

import "time"

type Option struct {
	ID        int       `json:"id"`
	PollID    int       `json:"poll_id"`
	Name      string    `json:"name"`
	Vote      int       `json:"vote"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
