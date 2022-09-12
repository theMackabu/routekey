package repository

import (
	"github.com/gin-gonic/gin"
	"routekey/models"
	"gorm.io/gorm"
)

type Domain interface {
	Create(ctx *gin.Context, domain models.Domain) (models.Domain, error)
	Read(ctx *gin.Context, domain models.Domain) (models.Domain, error)
	ReadAll(ctx *gin.Context) ([]models.Domain, error)
	Update(ctx *gin.Context, domain *models.Domain) error
	Delete(ctx *gin.Context, domain *models.Domain) error
}

type domain struct {
	db gorm.DB
}

func (repo *domain) Create(ctx *gin.Context, domain models.Domain) (models.Domain, error) {
	return domain, nil
}

func (repo *domain) Read(ctx *gin.Context, domain models.Domain) (models.Domain, error) {
	return domain, nil
}

func (repo *domain) ReadAll(ctx *gin.Context) ([]models.Domain, error) {
	return nil, nil
}

func (repo *domain) Update(ctx *gin.Context, domain *models.Domain) error {
	return nil
}

func (repo *domain) Delete(ctx *gin.Context, domain *models.Domain) error {
	return nil
}

func NewDomain(db *gorm.DB) Domain {
	return &domain{db: *db}
}
