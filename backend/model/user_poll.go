package model

type UserPoll struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PollID int `json:"poll_id"`
}
