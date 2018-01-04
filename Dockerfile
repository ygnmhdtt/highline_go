FROM golang:1.9.2-alpine

ENV GOPATH=/golang

WORKDIR /golang/highline-go

ADD . /golang/highline-go
