FROM golang:onbuild
//https://blog.golang.org/docker
ADD . /go/src/github.com/irom77/go-public/pan2influx
RUN go install github.com/irom77/go-public/pan2influx
ENTRYPOINT /go/bin/pan2influx -h