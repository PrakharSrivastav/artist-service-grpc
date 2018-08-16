FROM golang:alpine as builder
RUN apk update && apk add --no-cache git \
	&& apk add --no-cache ca-certificates \
	&& apk add --update --no-cache sqlite \
	&& apk add --no-cache build-base \
	&& adduser -D -g '' appuser \
	&& mkdir /go/src/github.com \
	&& mkdir /go/src/github.com/PrakharSrivastav \
	&& mkdir /go/src/github.com/PrakharSrivastav/artist-service-grpc
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