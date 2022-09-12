package repository

import (
	"errors"
	"time"

	"routekey/models"
	"routekey/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type URL interface {
	Redirect(ctx *gin.Context, link models.Link) (models.Link, error)
	GenQR(ctx *gin.Context, qr models.QRCode) (models.QRCode, error)
}

type url struct {
	db gorm.DB
}

func (repo *url) Redirect(ctx *gin.Context, link models.Link) (models.Link, error) {
	if link.Link == nil {
		return link, errors.New("link is empty")
	}

	result := repo.db.Find(&link, "link = ?", link.Link)
	if result.Error != nil {
		return link, result.Error
	}
	if link.ID == "" {
		return link, errors.New("link not found")
	}
	link.VisitCount++
	link.UpdatedAt = time.Now()

	tx := repo.db.Begin()
	if tx.Error != nil {
		return link, tx.Error
	}
	if err := tx.Save(&link).Error; err != nil {
		tx.Rollback()
		return link, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return link, err
	}

	return link, nil
}

func (u *url) GenQR(ctx *gin.Context, qr models.QRCode) (models.QRCode, error) {
	var err error
	qr.Image, err = utils.GenerateQRCode(qr.Content)
	if err != nil {
		return qr, err
	}
	return qr, nil
}

func NewURLRepo(db *gorm.DB) URL {
	return &url{db: *db}
}
