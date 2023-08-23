package routes

import (
	"routekey/client"
	"routekey/config"
	"routekey/controllers"
	"routekey/helpers"
	"routekey/middlewares"
	"routekey/services"

	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	StartTime time.Time
	BootTime  time.Duration
)

func Setup() *gin.Engine {
	router := gin.Default()
	svc := services.NewServices()

	// ------- word updater code - blobbybilb -------
	wordupdaterpassword := "changethis123"
	updater := router.Group("/updater" + wordupdaterpassword)

	updater.GET("/addword/:word", func(c *gin.Context) {
		config.AddWord(c.Param("word"))
	})

	updater.GET("/removeword/:word", func(c *gin.Context) {
		config.RemoveWord(c.Param("word"))
	})

	updater.GET("/listwords", func(c *gin.Context) {
		words := config.ReadWords()
		c.JSON(http.StatusOK, words)
	})
	// ------- END word updater code -------

	router.GET("/:link", func(c *gin.Context) {
		svc.URLService().Redirect(c)
	})

	router.GET("/:link/qrcode", func(c *gin.Context) {
		svc.URLService().GenQR(c)
	})

	staticFs := helpers.EmbedFolder(client.DistDir, "dist")
	staticServer := static.Serve("/", staticFs)

	router.Use(staticServer)
	router.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet && !strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	api := router.Group("/api")
	api.Use(middlewares.CORSMiddleware())

	v2 := api.Group("/v2")
	public := api.Group("/public")

	api.GET("/health", func(c *gin.Context) {
		svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
	})

	links := v2.Group("/links")
	domains := v2.Group("/domains")
	tracker := v2.Group("/trackers")

	links.GET("", middlewares.JWTAuth(), func(c *gin.Context) {
		svc.LinkService().GetLinks(c)
	})
	links.POST("", func(c *gin.Context) {
		svc.LinkService().AddLink(c)
	})
	links.POST("/custom", middlewares.JWTAuth(), func(c *gin.Context) {
		svc.LinkService().AddLinkAdmin(c)
	})
	links.GET("/:id", func(c *gin.Context) {
		svc.LinkService().GetLink(c)
	})
	links.GET("/:id/qrcode", func(c *gin.Context) {
		svc.LinkService().GenQRCode(c)
	})
	links.PATCH("/:id", middlewares.JWTAuth(), func(c *gin.Context) {
		svc.LinkService().UpdateLink(c)
	})
	links.DELETE("/:id", middlewares.JWTAuth(), func(c *gin.Context) {
		svc.LinkService().DeleteLink(c)
	})
	links.GET("/:id/stats", func(c *gin.Context) {
		svc.LinkService().GetLinkStats(c)
	})
	domains.GET("", func(c *gin.Context) {
		svc.DomainService().GetDomains(c)
	})
	tracker.GET("", func(c *gin.Context) {
		svc.TrackerService().GetTrackers(c)
	})
	tracker.GET("/gen", func(c *gin.Context) {
		svc.TrackerService().GenerateTracker(c)
	})
	tracker.GET("/:id", func(c *gin.Context) {
		svc.TrackerService().GetTracker(c)
	})
	tracker.GET("/:id/qr.png", func(c *gin.Context) {
		svc.TrackerService().QRCode(c)
	})
	tracker.GET("/:id/status", func(c *gin.Context) {
		svc.TrackerService().Status(c)
	})
	tracker.DELETE("/:id", func(c *gin.Context) {
		svc.TrackerService().DeleteTracker(c)
	})

	public.POST("/login", controllers.Login)
	public.POST("/signup", controllers.Signup)

	protected := api.Group("/protected").Use(middlewares.JWTAuth())
	protected.GET("/profile", controllers.Profile)

	return router
}
