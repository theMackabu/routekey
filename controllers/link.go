package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"routekey/models"
	"routekey/repository"
	"routekey/utils"
)

type Link interface {
	GetLinks(c *gin.Context)
	AddLink(c *gin.Context)
	AddLinkAdmin(c *gin.Context)
	GetLink(c *gin.Context)
	UpdateLink(c *gin.Context)
	DeleteLink(c *gin.Context)
	GetLinkStats(c *gin.Context)
	GenQRCode(c *gin.Context)
}

type link struct {
	link repository.Link
}

func (l *link) GetLinks(c *gin.Context) {
	links, err := l.link.ReadAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, links)
}

func (l *link) AddLinkAdmin(c *gin.Context) {
	var link models.LinkBody
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: err.Error(),
				Status: http.StatusBadRequest,
			},
		})
		return
	}
	if link.Target == "" {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: "Target is empty",
				Status: http.StatusBadRequest,
			},
		})
		return
	}

	newLink := models.Link{
		Target: &link.Target,
	}

	_, err := l.link.Read(c, newLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	newLink.ID = utils.GenerateUUID()
	newLink.Target = &link.Target
	if link.CustomURL != "" {
		url := c.Request.Host
		newLink.Link = &link.CustomURL
		url = url + "/" + link.CustomURL
		newLink.Address = &url
	} else {
		shortURL := utils.GenerateShortURL()
		newLink.Link = &shortURL
		url := c.Request.Host
		url = url + "/" + shortURL
		newLink.Address = &url
	}
	if link.Reusable != nil {
		newLink.Reusable = link.Reusable
	}
	if link.Description != "" {
		newLink.Description = &link.Description
	}
	if link.Password != "" {
		newLink.Password = &link.Password
	}
	if link.ExpireIn != "" {
		expireAt := utils.GetExpireAt(link.ExpireIn)
		if !expireAt.IsZero() {
			newLink.ExpireAt = &expireAt
		}
	} else {
		expireAt := utils.GetExpireAt("5 minutes")
		newLink.ExpireAt = &expireAt
	}
	
	ip := c.ClientIP()
	newLink.IP = &ip

	if err = l.link.Create(c, &newLink); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "Internal Error",
				Detail: "An internal error has occurred",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, newLink)
}

func (l *link) AddLink(c *gin.Context) {
	var link models.LinkBody
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: err.Error(),
				Status: http.StatusBadRequest,
			},
		})
		return
	}
	if link.Target == "" {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: "Target is empty",
				Status: http.StatusBadRequest,
			},
		})
		return
	}

	newLink := models.Link{
		Target: &link.Target,
	}

	_, err := l.link.Read(c, newLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	newLink.ID = utils.GenerateUUID()
	newLink.Target = &link.Target
	shortURL := utils.GenerateShortURL()
	newLink.Link = &shortURL
	url := c.Request.Host
	url = url + "/" + shortURL
	newLink.Address = &url
	
	
	if link.ExpireIn == "300" {
		expireAt := utils.GetExpireAt("5 minutes")
		newLink.ExpireAt = &expireAt
	} else if link.ExpireIn == "900" {
		expireAt := utils.GetExpireAt("15 minutes")
		newLink.ExpireAt = &expireAt
	} else if link.ExpireIn == "1800" {
		expireAt := utils.GetExpireAt("30 minutes")
		newLink.ExpireAt = &expireAt
	} else if link.ExpireIn == "3600" {
		expireAt := utils.GetExpireAt("60 minutes")
		newLink.ExpireAt = &expireAt
	} else {
		expireAt := utils.GetExpireAt("5 minutes")
		newLink.ExpireAt = &expireAt
	}
	
	ip := c.ClientIP()
	newLink.IP = &ip

	if err = l.link.Create(c, &newLink); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "Internal Error",
				Detail: "An internal error has occurred",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, newLink)
}

func (l *link) GetLink(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	c.JSON(http.StatusOK, link)
}

func (l *link) UpdateLink(c *gin.Context) {
	id := c.Param("id")
	var link models.LinkBody
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Type:   "bad_request",
				Title:  "Bad Request",
				Detail: err.Error(),
				Status: http.StatusBadRequest,
			},
		})
		return
	}

	existingLink, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if existingLink.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	if link.Target != "" {
		existingLink.Target = &link.Target
	}
	if link.CustomURL != "" {
		existingLink.Link = &link.CustomURL
		url := c.Request.Host + "/" + link.CustomURL
		existingLink.Address = &url
	}
	if link.Description != "" {
		existingLink.Description = &link.Description
	}
	if link.Password != "" {
		existingLink.Password = &link.Password
	}
	if link.ExpireIn != "" {
		expireAt := utils.GetExpireAt(link.ExpireIn)
		if !expireAt.IsZero() {
			existingLink.ExpireAt = &expireAt
		}
	}
	if link.Reusable != nil {
		existingLink.Reusable = link.Reusable
	}

	if err := l.link.Update(c, &existingLink); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, existingLink)
}

func (l *link) DeleteLink(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	if err := l.link.Delete(c, &link); err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, link)
}

func (l *link) GetLinkStats(c *gin.Context) {
	id := c.Param("id")
	stats, err := l.link.GetStats(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (l *link) GenQRCode(c *gin.Context) {
	id := c.Param("id")
	link, err := l.link.Read(c, models.Link{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	if link.ID == "" {
		c.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "Not Found",
				Detail: "Link not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	qrCode, err := l.link.GenQRCode(c, link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "internal_error",
				Title:  "An internal error has occurred",
				Detail: err.Error(),
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	c.Data(http.StatusOK, "image/png", qrCode.Image)
}

func NewLink(linkRepo repository.Link) Link {
	return &link{
		link: linkRepo,
	}
}
