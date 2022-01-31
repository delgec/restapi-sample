package model

import (
	"time"

	"github.com/gocql/gocql"
)

// Feedback struct for cassandra
type Feedback struct {
	Id           gocql.UUID
	SenderId     gocql.UUID
	Subject      string
	Message      string
	Device       string
	Attachment   string
	AdminComment string
	CreatedAt    time.Time
	// Replied        bool
	// RepliedBy      gocql.UUID
	// RepliedAt      time.Time
	// RepliedMessage string
}
