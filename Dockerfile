FROM golang:1.12.5-alpine3.9

# installing sqlite
RUN apk add sqlite
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN rm -rf /var/cache/apk/*

#installing gcc
RUN apk add build-base

# building go app
EXPOSE 8080
ENV GOPATH /go
RUN mkdir -p $GOPATH/bin
RUN mkdir -p $GOPATH/src/github.com/vbhv007/travel-log-api
ADD . $GOPATH/src/github.com/vbhv007/travel-log-api/
WORKDIR $GOPATH/src/github.com/vbhv007/travel-log-api/
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $GOPATH/bin/main-linux
WORKDIR $GOPATH/bin/
CMD ["./main-linux"]