FROM golang:1.9.2-alpine3.6 as builder

RUN apk add git --update curl
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/yulPa/yulmails
COPY . ./
RUN dep ensure -vendor-only; \
  GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.6
COPY --from=builder /go/src/github.com/yulPa/yulmails/main .
RUN mv main /usr/bin/yulmails && chmod +x /usr/bin/yulmails
CMD ["yulmails", \
  "api", \
  "--tls-crt-file", "/etc/yulmails/conf/yulmails.local.tld/yulmails.local.tld.crt", \
  "--tls-key-file", "/etc/yulmails/conf/yulmails.local.tld/yulmails.local.tld.key" \
]
