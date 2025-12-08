package repository

import (
	"context"

	"time"

	"anoa.com/telkomalumiforum/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	Create(ctx context.Context, attachment *model.Attachment) error
	UpdateThreadID(ctx context.Context, attachmentIDs []uint, threadID uuid.UUID, userID uuid.UUID) error
	FindOrphans(ctx context.Context, cutoffTime time.Time) ([]model.Attachment, error)
	Delete(ctx context.Context, id uint) error
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

func (r *attachmentRepository) Create(ctx context.Context, attachment *model.Attachment) error {
	return r.db.WithContext(ctx).Create(attachment).Error
}

func (r *attachmentRepository) UpdateThreadID(ctx context.Context, attachmentIDs []uint, threadID uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&model.Attachment{}).
		Where("id IN ? AND user_id = ?", attachmentIDs, userID).
		Update("thread_id", threadID).Error
}

func (r *attachmentRepository) FindOrphans(ctx context.Context, cutoffTime time.Time) ([]model.Attachment, error) {
	var attachments []model.Attachment
	err := r.db.WithContext(ctx).
		Where("thread_id IS NULL AND post_id IS NULL AND created_at < ?", cutoffTime).
		Find(&attachments).Error
	return attachments, err
}

func (r *attachmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Attachment{}, id).Error
}
