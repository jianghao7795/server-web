FROM golang:alpine as builder

WORKDIR /go/src/server-web
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="978561120@qq.com"

WORKDIR /go/src/server-web

COPY --from=0 /go/src/server-web/server ./
COPY --from=0 /go/src/server-web/resource ./resource/
COPY --from=0 /go/src/server-web/config.docker.yaml ./

EXPOSE 3100
ENTRYPOINT ./server -c config.docker.yaml
