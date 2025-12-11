package model

import (
	"time"

	"github.com/google/uuid"
)

type ThreadLike struct {
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid" json:"user_id"`
	ThreadID  uuid.UUID `gorm:"primaryKey;type:uuid" json:"thread_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type PostLike struct {
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid" json:"user_id"`
	PostID    uuid.UUID `gorm:"primaryKey;type:uuid" json:"post_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
