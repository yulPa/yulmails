FROM golang:1.9.2-alpine3.6 as builder
RUN apk add git --update
WORKDIR /go/src/github.com/check_mails
COPY . ./
RUN go get ./... && GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.6
WORKDIR /root/
COPY --from=builder /go/src/github.com/check_mails/main .
RUN chmod +x main
CMD ["./main"]
