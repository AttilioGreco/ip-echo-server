# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./


RUN go build -o /server

## Deploy
FROM alpine:3.16

WORKDIR /app
ADD ./templates/* /app/templates/
COPY --from=build /server /app

EXPOSE 8080

USER guest:nogroup

ENTRYPOINT ["/app/server"]