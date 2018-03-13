FROM alpine
MAINTAINER Eduardo Oliveira <eduardoliveira2009@gmail.com>

RUN apk update && apk add go ca-certificates git  sqlite build-base && rm -rf /var/cache/apk/*

RUN mkdir -p /root/go/src/github.com/EduardoOliveira/go-lorem-image
WORKDIR /root/go/src/github.com/EduardoOliveira/go-lorem-image
RUN git clone https://github.com/EduardoOliveira/go-lorem-image.git
WORKDIR /root/go/src/github.com/EduardoOliveira/go-lorem-image
RUN go get
RUN go build

EXPOSE 9999
ENTRYPOINT ./go-lorem-image