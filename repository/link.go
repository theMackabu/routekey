package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"routekey/models"
	"routekey/utils"
	"gorm.io/gorm"
)

type Link interface {
	Create(ctx *gin.Context, link *models.Link) error
	Read(ctx *gin.Context, link models.Link) (models.Link, error)
	ReadAll(ctx *gin.Context) ([]models.Link, error)
	Update(ctx *gin.Context, link *models.Link) error
	Delete(ctx *gin.Context, link *models.Link) error
	GetStats(ctx *gin.Context, link models.Link) (models.Stat, error)
	GenQRCode(ctx *gin.Context, link models.Link) (models.QRCode, error)
}

type link struct {
	db gorm.DB
}

func (repo *link) Create(ctx *gin.Context, link *models.Link) error {
	if link.Address == nil {
		return errors.New("address is empty")
	}
	if link.Target == nil {
		return errors.New("target is empty")
	}
	if link.Link == nil {
		return errors.New("link is empty")
	}
	var l models.Link
	result := repo.db.Find(&l, "address = ?", link.Address)
	if result.Error != nil {
		return result.Error
	}
	if l.ID != "" {
		return errors.New("link already exists")
	}
	existingLink := models.Link{}
	result = repo.db.Find(&existingLink, "link = ?", link.Link)
	if result.Error != nil {
		return result.Error
	}
	if link.Link == existingLink.Link {
		return errors.New("link already exists")
	}

	tx := repo.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Create(&link).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *link) Read(ctx *gin.Context, link models.Link) (models.Link, error) {
	var l models.Link
	if link.ID == "" && link.Target == nil {
		return l, errors.New("id is empty")
	}
	if link.ID == "" {
		result := repo.db.Find(&l, "target = ?", link.Target)
		if result.Error != nil {
			return l, result.Error
		}
	} else {
		result := repo.db.Find(&l, "id = ?", link.ID)
		if result.Error != nil {
			return l, result.Error
		}
	}

	return l, nil
}

func (repo *link) ReadAll(ctx *gin.Context) ([]models.Link, error) {
	var links []models.Link
	if err := repo.db.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (repo *link) Update(ctx *gin.Context, link *models.Link) error {
	if link.ID == "" {
		return errors.New("id is empty")
	}

	var l models.Link
	result := repo.db.Find(&l, "id = ?", link.ID)
	if result.Error != nil {
		return result.Error
	}

	if link.Address != nil {
		l.Address = link.Address
	}
	if link.Target != nil {
		l.Target = link.Target
	}
	if link.Link != nil {
		l.Link = link.Link
	}
	if link.Banned != nil {
		l.Banned = link.Banned
	}
	if link.Password != nil {
		l.Password = link.Password
	}
	if link.Description != nil {
		l.Description = link.Description
	}
	if link.ExpireAt.IsZero() {
		l.ExpireAt = link.ExpireAt
	}
	if link.Reusable != nil {
		l.Reusable = link.Reusable
	}

	if link.CreatedAt.IsZero() {
		l.CreatedAt = link.CreatedAt
	}
	if link.UpdatedAt.IsZero() {
		l.UpdatedAt = link.UpdatedAt
	}

	tx := repo.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Save(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *link) Delete(ctx *gin.Context, link *models.Link) error {
	if link.ID == "" {
		return errors.New("id is empty")
	}
	tx := repo.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Delete(&link).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *link) GetStats(ctx *gin.Context, link models.Link) (models.Stat, error) {
	var stat models.Stat
	result := repo.db.Find(&stat, "link = ?", link.Link)
	if result.Error != nil {
		return stat, result.Error
	}
	return stat, nil
}

func (repo *link) GenQRCode(ctx *gin.Context, link models.Link) (models.QRCode, error) {
	var qrCode models.QRCode
	var err error
	result := repo.db.Find(&link, "id = ?", link.ID)
	if result.Error != nil {
		return qrCode, result.Error
	}
	addr := ctx.Request.Host + "/" + *link.Link
	qrCode.Image, err = utils.GenerateQRCode(addr)
	if err != nil {
		return qrCode, err
	}
	return qrCode, nil
}

func NewLink(db *gorm.DB) Link {
	return &link{db: *db}
}
