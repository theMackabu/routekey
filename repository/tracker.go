package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"routekey/models"
	"gorm.io/gorm"
)

type Tracker interface {
	GenerateTracker(ctx *gin.Context, track *models.Tracker) error
	GetTrackers(ctx *gin.Context) ([]models.Tracker, error)
	GetTracker(ctx *gin.Context, track models.Tracker) (models.Tracker, error)
	Update(ctx *gin.Context, track *models.Tracker) error
	Delete(ctx *gin.Context, track *models.Tracker) error
}

type tracker struct {
	db gorm.DB
}

func (repo *tracker) GenerateTracker(ctx *gin.Context, track *models.Tracker) error {
	if track.ID == "" {
		return errors.New("id is empty")
	}

	existingLink := models.Tracker{}
	result := repo.db.Find(&existingLink, "id = ?", track.ID)
	if result.Error != nil {
		return result.Error
	}
	if existingLink.ID != "" {
		return errors.New("tracker already exists")
	}

	tx := repo.db.Begin()
	if tx.Create(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (repo *tracker) GetTrackers(ctx *gin.Context) ([]models.Tracker, error) {
	trackers := []models.Tracker{}
	if err := repo.db.Find(&trackers).Error; err != nil {
		return nil, err
	}
	return trackers, nil
}

func (repo *tracker) GetTracker(ctx *gin.Context, track models.Tracker) (models.Tracker, error) {
	var tracker models.Tracker
	if err := repo.db.Find(&tracker, "id = ?", track.ID).Error; err != nil {
		return tracker, err
	}
	return tracker, nil
}

func (repo *tracker) Update(ctx *gin.Context, track *models.Tracker) error {
	if track.ID == "" {
		return errors.New("id is empty")
	}

	tx := repo.db.Begin()
	if tx.Save(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (repo *tracker) Delete(ctx *gin.Context, track *models.Tracker) error {
	if track.ID == "" {
		return errors.New("id is empty")
	}

	tx := repo.db.Begin()
	if tx.Delete(&track).Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func NewTracker(db *gorm.DB) Tracker {
	return &tracker{db: *db}
}
