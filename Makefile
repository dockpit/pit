debug:
	go-bindata -o=server/ui/bin/assets_dev.go -tags=debug -pkg=uibin -ignore=server/ui/bin -debug=true server/ui/...
	go run -tags=debug main.go start -bind=:9000 

release:
	go-bindata -o=server/ui/bin/assets_prod.go -tags=!debug -pkg=uibin -ignore=server/ui/bin server/ui/...
	go build -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" -o $(GOPATH)/bin/pit main.go

publish-release:
	github-release release \
    --user dockpit \
    --repo pit \
    --tag v$(shell cat VERSION) \
    --pre-release

publish-upload:
	github-release upload \
    --user dockpit \
    --repo pit \
    --tag v$(shell cat VERSION) \
    --name "pit-v$(shell cat VERSION)-osx-amd64" \
    --file $(GOPATH)/bin/pit

publish-tag:
	git tag v$(shell cat VERSION)
	git push --tags

publish: release publish-tag publish-release publish-upload