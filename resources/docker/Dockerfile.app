FROM golang:1.15

ENV GOPATH /go
WORKDIR $GOPATH/src/app
ENV GO111MODULE=on
ENV PORT=8080
# COPY . .
# RUN chmod +x s3manager