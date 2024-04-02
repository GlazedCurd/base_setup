# syntax=docker/dockerfile:1

FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /testservice

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /testservice /testservice

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/testservice"]