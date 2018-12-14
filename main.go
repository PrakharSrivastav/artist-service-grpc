package main

import (
	"fmt"
	"github.com/PrakharSrivastav/artist-service-grpc/internal"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	checkErr(err)
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(sameSpan),
	)

	// Create a new application
	app := internal.NewApplication(tracer)

	// Create a new db connection
	connection, err := sqlx.Open("sqlite3", "./artist.db")
	checkErr(err)

	// Inject connection to service
	internal := service.New(connection,tracer)

	// Initial data migration
	checkErr(internal.CleanupAndInit())

	// Inject internal service to artistService
	artistService := rpc.New(internal)

	// Inject artist service to grpc
	app.GrpcServer.Add(artistService)

	// Start grpc server
	checkErr(app.GrpcServer.Start())
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

const (
	// Our service name.
	serviceName = "artist-service"

	// Host + port of our service.
	hostPort = "127.0.0.1:6565"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://zipkin:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false
)
