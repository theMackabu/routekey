package controllers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"routekey/database"
)

const (
	StatusOK                 string = "OK"
	StatusPartiallyAvailable string = "Partially Available"
	StatusUnavailable        string = "Unavailable"
	StatusTimeout            string = "Timeout during health check"
)

type (
	Check struct {
		Status string `json:"status"`
		Timestamp string `json:"timestamp"`
		StartUp string `json:"startup"`
		Uptime string `json:"uptime"`
		Databases map[string]string `json:"databases"`
		System `json:"system"`
	}

	System struct {
		Version string `json:"version"`
		GoroutinesCount int `json:"goroutines_count"`
		TotalAllocBytes int `json:"total_alloc_bytes"`
		HeapObjectsCount int `json:"heap_objects_count"`
		AllocBytes int `json:"alloc_bytes"`
	}
)

func NewSystemMetrics() System {
	s := runtime.MemStats{}
	runtime.ReadMemStats(&s)

	return System{
		Version:          runtime.Version(),
		GoroutinesCount:  runtime.NumGoroutine(),
		TotalAllocBytes:  int(s.TotalAlloc),
		HeapObjectsCount: int(s.HeapObjects),
		AllocBytes:       int(s.Alloc),
	}
}

func NewCheck(status string, databases map[string]string, startTime time.Time, bootTime time.Duration) Check {
	return Check{
		Status:    status,
		Timestamp: time.Now().Format(time.RFC3339),
		StartUp:   bootTime.String(),
		Uptime:    time.Since(startTime).String(),
		Databases:  databases,
		System:    NewSystemMetrics(),
	}
}

type HealthCheck interface {
	HealthCheck(ctx *gin.Context, startTime time.Time, bootTime time.Duration)
}

type healthCheck struct {
}
func (m *healthCheck) HealthCheck(ctx *gin.Context, startTime time.Time, bootTime time.Duration) {
	status := StatusUnavailable
	databases := make(map[string]string)

	if database.IsConnected {
		status = StatusOK
		if database.IsSQLite {
			databases["sqlite"] = "using sqlite database"
		}
	} else {
		status = StatusPartiallyAvailable
		databases["sqlite"] = "failed during sqlite health check"
	}

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.JSON(http.StatusOK, NewCheck(
		status,
		databases,
		startTime,
		bootTime,
	))
}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}
