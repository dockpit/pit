FROM google/golang:1.3

#we usegodep
RUN go get github.com/tools/godep
RUN mkdir -p /gopath/src/github.com/dockpit/pit

WORKDIR /gopath/src/github.com/dockpit/pit
ADD . /gopath/src/github.com/dockpit/pit
RUN godep go build -o /gopath/bin/pit -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`"

CMD []
EXPOSE 8000
ENTRYPOINT ["/gopath/bin/pit"]