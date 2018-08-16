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
│   │   └── server.go       // GRPC server settings (host,port) for this service
│   ├── client
│   │   ├── artist.go       // client calls to other grpc services
│   │   ├── client.go       // GRPC client settings (host,port) to connect to other services
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

## Kubernetes config
Kubernetes configurations are provided in k8s directory. If you are publishing the docker images to your repository, please make changes in `k8s/deployment.yml` and point to your repo.

## Docker configuration
Check `Dockerfile` for image configuration and `docker-compose.yml` for repo and tagging details. Please point to your repository if you want to use different docker-hub repo.

## Deployment instructions
Please follow the below instructions for development and deployment.
- Make the code changes.
- Change the docker image repo and tag in `docker-compose.yml`
- login to your docker hub via command-line, if you need to push to docker-hub by running `docker login`. It will ask you for username and password.
- run `docker-compose build`
- run `docker-compose push`
- change docker image tag in `k8s/deployment.yml`
- run `kubectl apply -f k8s/` from root of project to deploy to kubernetes.

