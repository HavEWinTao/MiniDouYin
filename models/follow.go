package models

import "time"

type Follow struct {
	ID         int64
	FolloweeID int64
	FollowerID int64
	IsFollow   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
