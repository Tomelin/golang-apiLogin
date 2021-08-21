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

RUN cp -aR /app /usr/local/go/src/apiLogin
RUN go get github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .