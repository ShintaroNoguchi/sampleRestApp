FROM golang:1.14.4-alpine3.12

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]