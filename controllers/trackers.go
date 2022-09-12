package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"routekey/models"
	"routekey/repository"
	"routekey/utils"
)

type Trackers interface {
	GenerateTracker(ctx *gin.Context)
	GetTrackers(ctx *gin.Context)
	GetTracker(ctx *gin.Context)
	QRCode(ctx *gin.Context)
	Status(ctx *gin.Context)
	DeleteTracker(ctx *gin.Context)
}

type trackers struct {
	track repository.Tracker
}

func (t *trackers) GenerateTracker(ctx *gin.Context) {
	tracker := models.Tracker{}
	var err error

	tracker.ID = utils.GenerateShortURL()
	ip := ctx.ClientIP()
	tracker.IP = &ip
	hostname := ctx.Request.Host
	url := "http://" + hostname + "/api/v1/trackers/" + tracker.ID + "/qr.png"
	tracker.URL = url

	err = t.track.GenerateTracker(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to save tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, tracker)
}

func (t *trackers) GetTrackers(ctx *gin.Context) {
	trackers, err := t.track.GetTrackers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get trackers",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, trackers)
}

func (t *trackers) GetTracker(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, tracker)
}

func (t *trackers) QRCode(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}

	tracker.VisitCount++
	err = t.track.Update(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to update tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	image := utils.GenerateTrackerImage()
	ctx.Data(http.StatusOK, "image/png", image)
}

func (t *trackers) Status(ctx *gin.Context) {
	var trackerStatus models.TrackerStatus
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	trackerStatus.ID = tracker.ID
	trackerStatus.URL = tracker.URL
	trackerStatus.Message = "not seen"
	if tracker.VisitCount > 0 {
		trackerStatus.Seen = true
		trackerStatus.Message = "seen"
	}

	ctx.JSON(http.StatusOK, trackerStatus)
}

func (t *trackers) DeleteTracker(ctx *gin.Context) {
	id := ctx.Param("id")
	tracker := models.Tracker{
		ID: id,
	}
	tracker, err := t.track.GetTracker(ctx, tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to get tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}

	if tracker.ID == "" {
		ctx.JSON(http.StatusNotFound, models.Error{
			Error: models.ServiceError{
				Type:   "not_found",
				Title:  "not_found",
				Detail: "tracker not found",
				Status: http.StatusNotFound,
			},
		})
		return
	}
	err = t.track.Delete(ctx, &tracker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ServiceError{
				Type:   "service_error",
				Title:  "service_error",
				Detail: "failed to delete tracker",
				Status: http.StatusInternalServerError,
			},
		})
		return
	}
	msg := models.Message{
		Code:    http.StatusOK,
		Message: "tracker deleted",
	}
	ctx.JSON(http.StatusOK, msg)
}

func NewTrackers(tracker repository.Tracker) Trackers {
	return &trackers{
		track: tracker,
	}
}
