package services

type ServiceContainer interface {
	RegisterImageDrawingService(service ImageDrawingService)
	ImageDrawingService() ImageDrawingService
}

type serviceContainer struct {
	imageDrawingService ImageDrawingService
}

func NewServiceContainer() ServiceContainer {
	return &serviceContainer{}
}

func (c *serviceContainer) RegisterImageDrawingService(service ImageDrawingService) {
	c.imageDrawingService = service
}
func (c *serviceContainer) ImageDrawingService() ImageDrawingService {
	return c.imageDrawingService
}
