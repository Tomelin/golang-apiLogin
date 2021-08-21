FROM golang:1.16-alpine AS build

WORKDIR /app

RUN apk add --no-cache git
RUN mkdir /app/bin
ADD .env /app
ADD go.mod /app
ADD go.sum /app
RUN go mod download
RUN go mod verify
ADD main.go /app
ADD src /app
RUN cd /app/bin
RUN export GOROOT=/app
RUN go get github.com/gorilla/mux
RUN go build


FROM alpine:3.9 

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=build /app/build_api /app/api

# This container exposes port 8080 to the outside world
EXPOSE 5002

# Run the binary program produced by `go install`
CMD ["/app/api"]