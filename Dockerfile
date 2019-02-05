FROM golang:latest

ENV DEBCONF_NOWARNINGS yes
ENV GOPATH $GOPATH:/go/src

ARG PACKAGE="github.com/hiromisuzuki/clean-arch-example"
RUN mkdir -p /go/src/${PACKAGE}
WORKDIR /go/src/${PACKAGE}
COPY . .

RUN apt-get update && \
  apt-get upgrade -y && \
  pwd && \
  ls -lta && \
  go get github.com/lib/pq && \
  go get github.com/pilu/fresh && \
  go get github.com/alecthomas/template && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/golang/dep/cmd/dep && \
  go get github.com/swaggo/swag/cmd/swag && \
  go get bitbucket.org/liamstask/goose/cmd/goose && \
  dep ensure

CMD fresh -c config/fresh.conf main.go

EXPOSE 8080