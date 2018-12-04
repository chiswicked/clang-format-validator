FROM golang:1.11.2-alpine3.8
RUN apk update && apk add git gcc make musl-dev
ADD . /go/src/github.com/chiswicked/clang-format-validator
WORKDIR /go/src/github.com/chiswicked/clang-format-validator
RUN make clean install test build