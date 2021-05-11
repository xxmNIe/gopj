FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/my



EXPOSE 8000
COPY .  $GOPATH/src/my
WORKDIR $GOPATH/src/my
RUN go build .
ENTRYPOINT ["./example-web"]
