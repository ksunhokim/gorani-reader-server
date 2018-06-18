FROM golang:stretch

WORKDIR /go/src/github.com/sunho/gorani-reader-server/go

COPY ./Gopkg.lock ./Gopkg.toml ./

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -vendor-only
