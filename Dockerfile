FROM golang:1.13.1-alpine3.10 AS builder

WORKDIR /go/src/gitlab.com/tortuemat/yulmails

COPY . .

RUN go build -o yulctl

FROM alpine:3.10

COPY --from=builder /go/src/gitlab.com/tortuemat/yulmails/yulctl /usr/local/bin

