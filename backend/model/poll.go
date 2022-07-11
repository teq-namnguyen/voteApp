package model

import "time"

type Poll struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OptionInPoll struct {
	Poll
	Options []Option `json:"options"`
}
