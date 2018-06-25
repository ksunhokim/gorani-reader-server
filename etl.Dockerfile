FROM golang:alpine
WORKDIR /go/src/github.com/sunho/gorani-reader-server
ADD . .
RUN make install-etl
WORKDIR /home
CMD ["etl"]
