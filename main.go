package main

import (
	"fmt"

	"github.com/PrakharSrivastav/artist-service-grpc/internal"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create a new application
	app := internal.NewApplication()

	// Create a new db connection
	connection, err := sqlx.Open("sqlite3", "./artist.db")
	checkErr(err)

	// Inject connection to service
	internal := service.New(connection)

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
