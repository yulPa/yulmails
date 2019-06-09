FROM golang:1.12-alpine3.9 as builder
RUN apk add --update \
	git
WORKDIR /opt
ENV GO111MODULE=on

# download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build project
COPY . .
RUN CG0_ENABLED=0 GOOS=linux go build -o conservation

FROM alpine:3.9

WORKDIR /opt

COPY --from=builder /opt/conservation .

CMD ["/opt/conservation"]
