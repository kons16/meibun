FROM golang:1.14-alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/kons16/meibun/api-server

COPY . .
RUN apk add --no-cache
   && go get github.com/pilu/fresh
