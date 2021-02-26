FROM golang:1.15-alpine

RUN mkdir -p /go/src/github.com/coretrix/skeleton

WORKDIR /go/src/github.com/coretrix/skeleton

ADD services/docker-entrypoint.sh /usr/bin/docker-entrypoint
RUN chmod +x /usr/bin/docker-entrypoint

ENTRYPOINT ["docker-entrypoint"]