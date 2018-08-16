# Artist Service
This project provides the server side implementations for artist services defined in the project [gql-grpc-defintions](https://github.com/PrakharSrivastav/gql-grpc-defintions). Please have a look at [artist.proto](https://github.com/PrakharSrivastav/gql-grpc-defintions/blob/master/schema/artist.proto) for protobuf definitions. This application also communicates with `track-service` and `album-service` via grpc.

This service also provides the kubernetes configurations under `k8s/` directory. 

## Project structure
```
.
├── artist.db               // sqlite database
├── docker-compose.yaml     // docker-compose
├── Dockerfile              // Dockerfile
├── internal                
│   ├── app.go              // Wrapper for application
│   ├── backend
│   │   └── server.go       // GRPC server settings
│   ├── client
│   │   ├── artist.go       // client calls to other grpc services
│   │   ├── client.go       // GRPC client settings
│   │   └── rpcClient.go    // GRPC connection infra
│   ├── model
│   │   └── artist.go       // Database entity representations
│   ├── rpc
│   │   └── artist.go       // Implementation for grpc stubs
│   └── service
│       ├── impl.go         // Internal service implementations
│       ├── repository.go   // database layer
│       └── service.go      // Internal service definitions
├── k8s
│   ├── deployment.yml      // kubernetes deployment config
│   └── service.yml         // kubernetes service config
├── main.go                 // Entrypoint
└── README.md
```