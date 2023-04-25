FROM golang:latest as builder

WORKDIR /go/src/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/server

COPY --from=builder /go/src/server/server ./
COPY --from=builder /go/src/server/resource ./resource/
COPY --from=builder /go/src/server/config.docker.yaml ./
COPY --from=builder /go/src/server/rbac_model.conf ./

EXPOSE 3100
ENTRYPOINT ./server -c config.docker.yaml
