FROM golang:alpine as builder
RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/PrakharSrivastav
RUN mkdir /go/src/github.com/PrakharSrivastav/artist-service-grpc
COPY . /go/src/github.com/PrakharSrivastav/artist-service-grpc
WORKDIR  /go/src/github.com/PrakharSrivastav/artist-service-grpc
RUN go build -o app

FROM alpine
RUN mkdir /application
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/PrakharSrivastav/artist-service-grpc/app /application
WORKDIR /application
USER appuser
EXPOSE 6565
ENTRYPOINT ["/application/app"]