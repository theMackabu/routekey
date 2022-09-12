package services

import (
	"routekey/controllers"
	"routekey/database"
	"routekey/repository"
)

type Services interface {
	HealthCheckService() controllers.HealthCheck
	LinkService() controllers.Link
	URLService() controllers.URL
	DomainService() controllers.Domain
	TrackerService() controllers.Trackers
}

type services struct {
	healthCheck controllers.HealthCheck
	link        controllers.Link
	url         controllers.URL
	domain      controllers.Domain
	trackers    controllers.Trackers
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

func (svc *services) LinkService() controllers.Link {
	return svc.link
}

func (svc *services) URLService() controllers.URL {
	return svc.url
}

func (svc *services) DomainService() controllers.Domain {
	return svc.domain
}

func (svc *services) TrackerService() controllers.Trackers {
	return svc.trackers
}

func NewServices() Services {
	db := database.Initialize()
	return &services{
		healthCheck: controllers.NewHealthCheck(),
		link: controllers.NewLink(
			repository.NewLink(db),
		),
		url: controllers.NewURL(
			repository.NewURLRepo(db),
		),
		domain: controllers.NewDomain(),
		trackers: controllers.NewTrackers(
			repository.NewTracker(db),
		),
	}
}
