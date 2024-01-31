package application

import (
	ip_ebus "github.com/octoposprime/op-be-user/internal/application/infrastructure/port/ebus"
	ip_repo "github.com/octoposprime/op-be-user/internal/application/infrastructure/port/repository"
	ip_service "github.com/octoposprime/op-be-user/internal/application/infrastructure/port/service"
	ds "github.com/octoposprime/op-be-user/internal/domain/service"
)

// Service is an application service.
// It manages the business logic of the application.
type Service struct {
	*ds.Service
	ip_repo.DbPort
	ip_repo.RedisPort
	ip_ebus.EBusPort
	ip_service.ServicePort
}

// NewService creates a new *Service.
func NewService(domainService *ds.Service, dbRepository ip_repo.DbPort, redisRepository ip_repo.RedisPort, eBus ip_ebus.EBusPort, internalService ip_service.ServicePort) *Service {
	service := &Service{
		domainService,
		dbRepository,
		redisRepository,
		eBus,
		internalService,
	}
	service.DbPort.SetLogger(service.Log)
	service.EBusPort.SetLogger(service.Log)
	service.EventListen()
	service.Migrate()
	return service
}
