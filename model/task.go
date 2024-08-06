package model

import (
	"time"

	"github.com/LukaGiorgadze/gonull"
)

type Task struct {
	Id           int                        `json:"id"`
	Title        string                     `json:"title"`
	Description  gonull.Nullable[string]    `json:"description"`
	CreatedAt    time.Time                  `json:"created_at"`
	CompletedAt  gonull.Nullable[time.Time] `json:"completed_at"`
	IsCompleted  bool                       `json:"is_completed"`
	RewardInSats int                        `json:"rewards_in_sats"`
}
