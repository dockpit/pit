FROM google/golang:1.3

RUN go get -d github.com/docker/docker/...
WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" app

CMD []
EXPOSE 8000
ENTRYPOINT ["/gopath/bin/app"]