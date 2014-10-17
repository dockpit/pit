XC_ARCH = "386 amd64 arm"
XC_OS = "linux darwin windows freebsd openbsd"

#cross-compile binaries
build: 
	rm -fr bin/*
	mkdir -p bin/
	@echo "Building..."
	gox \
	    -os=$(XC_OS) \
	    -arch=$(XC_ARCH) \
	    -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" \
	    -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}" \
	    ./...

#build for local testing
dev:
	gox \
		-os="`go env GOOS`" \
		-arch="`go env GOARCH`" \
		-output "${GOPATH}/bin/pit" \
		./...

#run all unit tests
test:
	godep go test ./...