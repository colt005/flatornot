package models

import "encoding/json"

type VoteRow struct {
	ID        int64           `json:"id"`
	Vote      VoteType        `json:"vote"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt string          `json:"created_at"`
}

type Metadata struct {
}

type VoteTotalResult struct {
	Vote  VoteType `json:"vote"`
	Total int      `json:"total"`
}
