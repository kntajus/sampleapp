# syntax=docker/dockerfile:1

FROM golang:1.17.5-alpine3.15 as build
ENV CGO_ENABLED="0"
WORKDIR /app
COPY api api
COPY protos protos
COPY go.mod go.sum ./
WORKDIR /app/api
RUN go build -o api .

FROM scratch
COPY --from=build /app/api/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
