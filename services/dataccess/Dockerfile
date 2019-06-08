FROM golang:1.12-alpine3.9 as builder
RUN apk add --update \
	git
WORKDIR /opt
ENV GO111MODULE=on

COPY . .
RUN CG0_ENABLED=0 GOOS=linux go build -o dataccess

FROM alpine:3.9

WORKDIR /opt

COPY --from=builder /opt/dataccess .

CMD ["/opt/dataccess"]
