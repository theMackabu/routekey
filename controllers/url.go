package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"routekey/models"
	"routekey/repository"
)

type URL interface {
	Redirect(ctx *gin.Context)
	Track(ctx *gin.Context)
	GenQR(ctx *gin.Context)
}

type url struct {
	url repository.URL
}

func (u *url) Redirect(ctx *gin.Context) {
	fmt.Println(1234)
	url := ctx.Param("link")
	link, err := u.url.Redirect(ctx, models.Link{Link: &url})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	if time.Since(*link.ExpireAt).Seconds() > 0 {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: "link has expired",
				Status: http.StatusInternalServerError,
			},
		})
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	raterFile, _ := os.ReadFile("./rater/rate.html")
	raterString := string(raterFile)
	raterString = strings.ReplaceAll(raterString, "__routekeyname__", *link.Link)
	raterString = strings.ReplaceAll(raterString, "__routekeylink__", *link.Target)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(raterString))

	// target := *link.Target
	// ctx.Redirect(http.StatusFound, target)
}

func (u *url) Track(ctx *gin.Context) {
	url := ctx.Param("link")

	ctx.JSON(http.StatusOK, models.Link{
		Link: &url,
	})
}

func (u *url) GenQR(ctx *gin.Context) {
	var qr models.QRCode
	qr.Content = ctx.Param("link")
	hostname := ctx.Request.Host
	qr.Content = "http://" + hostname + "/" + qr.Content
	qr, err := u.url.GenQR(ctx, qr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	ctx.Data(http.StatusOK, "image/png", qr.Image)
}

func NewURL(urlRepo repository.URL) URL {
	return &url{
		url: urlRepo,
	}
}
