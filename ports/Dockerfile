# syntax=docker/dockerfile:1

FROM golang:1.17.5-alpine3.15 as build
ENV CGO_ENABLED="0"
WORKDIR /app
COPY ports ports
COPY protos protos
COPY go.mod go.sum ./
WORKDIR /app/ports
RUN go build -o ports .

FROM scratch
COPY --from=build /app/ports/ports /ports
ENTRYPOINT ["/ports"]
