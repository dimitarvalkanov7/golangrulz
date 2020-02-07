FROM golang

ADD . /go/src/github.com/dimitarvalkanov7/golangrulz

RUN go install github.com/dimitarvalkanov7/golangrulz

ENTRYPOINT /go/bin/golangrulz

EXPOSE 8080