FROM golang:1.13.1-alpine3.10 AS builder

WORKDIR /yulmails
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /go/src/gitlab.com/tortuemat/yulmails/plugins/spamassassin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o spamassassin-plugin

FROM alpine:3.10

RUN apk add --update spamassassin
RUN addgroup spamassassin && \
adduser -S \
       -D -G spamassassin \
       -h /var/lib/spamassassin/ \
   spamassassin && \
   sa-update
EXPOSE 12800
COPY --from=builder /go/src/gitlab.com/tortuemat/yulmails/plugins/spamassassin/spamassassin-plugin /usr/local/bin


