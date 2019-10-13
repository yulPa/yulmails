FROM golang:1.13.1-alpine3.10 AS builder

WORKDIR /yulmails
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /go/src/gitlab.com/tortuemat/yulmails
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yulctl

FROM alpine:3.10

COPY --from=builder /go/src/gitlab.com/tortuemat/yulmails/yulctl /usr/local/bin

