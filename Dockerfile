FROM alpine
MAINTAINER Eduardo Oliveira <eduardoliveira2009@gmail.com>

RUN apk update && apk add go ca-certificates git build-base && rm -rf /var/cache/apk/*

RUN mkdir -p /root/go/src/github.com/EduardoOliveira/goloremimage
WORKDIR /root/go/src/github.com/EduardoOliveira/goloremimage

RUN git clone https://github.com/EduardoOliveira/go-lorem-image.git .
RUN go get
RUN go build

EXPOSE 9999
ENTRYPOINT ./goloremimage