FROM golang:alpine

RUN mkdir -p /go/src/github.com/erning/hello-echo
WORKDIR /go/src/github.com/erning/hello-echo

EXPOSE 5000
