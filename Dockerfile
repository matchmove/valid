FROM golang:1.8

ENV APPDIR $GOPATH/src/github.com/matchmove/valid

ADD . $APPDIR
WORKDIR $APPDIR

RUN go get ./...
