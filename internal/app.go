package internal

import (
	"github.com/PrakharSrivastav/artist-service-grpc/internal/backend"
	"github.com/opentracing/opentracing-go"
)

// Application is a wrapper for all the backend infrastructure
type Application struct {
	GrpcServer *backend.Server
}

// NewApplication returns an instance of Application
func NewApplication(tracer opentracing.Tracer) *Application {
	app := Application{}
	app.GrpcServer = backend.NewGrpcServer(tracer)
	return &app
}
