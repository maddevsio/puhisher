package model

import "time"

type (
	// Standup model used for serialization/deserialization stored standups
	Standup struct {
		ID       int64     `db:"id" json:"id"`
		Created  time.Time `db:"created" json:"created"`
		Modified time.Time `db:"modified" json:"modified"`
		Username string    `db:"username" json:"userName"`
		Comment  string    `db:"comment" json:"comment"`
		GroupID  int64     `db:"groupid" json:"groupid"`
	}

	// Intern rerpesents intern
	Intern struct {
		ID       int64  `db:"id"`
		Username string `db:"username"`
		Lives    int    `db:"lives"`
		GroupID  int64  `db:"groupid" json:"groupid"`
	}
)
