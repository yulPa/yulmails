FROM golang:1.9.2-alpine3.6 as builder
RUN apk add git --update
WORKDIR /go/src/github.com/yulPa/yulmails
COPY . ./
RUN go get ./... && GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.6
WORKDIR /etc/yulmails
COPY --from=builder /go/src/github.com/yulPa/yulmails/main .
RUN chmod +x main
CMD ["./main"]
