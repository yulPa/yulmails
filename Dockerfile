FROM golang:1.9.2-alpine3.6 as builder

RUN apk add git --update curl
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/yulPa/yulmails
COPY . ./
RUN dep ensure -vendor-only; \
  GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.6
WORKDIR /etc/yulmails
COPY --from=builder /go/src/github.com/yulPa/yulmails/main .
COPY --from=builder /go/src/github.com/yulPa/yulmails/conf/ .
RUN chmod +x main
CMD ["./main"]
