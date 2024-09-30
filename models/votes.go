package models

import (
	"encoding/json"
	"time"
)

type VoteRow struct {
	ID        *int64          `json:"id"`
	Vote      VoteType        `json:"vote"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt *time.Time      `json:"created_at"`
	SessionID string          `json:"session_id"`
}

type Metadata struct {
}

type VoteTotalResult struct {
	Vote  VoteType `json:"vote"`
	Total int      `json:"total"`
}
