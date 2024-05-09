package model

import "time"

// RollResult secret santa partners list/draw.
type RollResult struct {
	Id          string               // hash or id
	ResultMap   map[string][2]string // map[id]partnerIds
	Description string               // short intro
	CreatedAt   time.Time
}
