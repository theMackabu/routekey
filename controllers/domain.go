package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"routekey/models"
)

type Domain interface {
	GetDomains(ctx *gin.Context)
}

type domain struct {
}

func (d *domain) GetDomains(ctx *gin.Context) {
	var domains []models.Domain
	ctx.JSON(http.StatusNotImplemented, domains)
}

func NewDomain() Domain {
	return &domain{}
}
