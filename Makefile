XC_ARCH = "386 amd64 arm"
XC_OS = "linux darwin windows freebsd openbsd"

#cross-compile
build: 
	rm -fr bin/*
	mkdir -p bin/
	@echo "Building..."
	gox \
	    -os=$(XC_OS) \
	    -arch=$(XC_ARCH) \
	    -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}" \
	    ./...

#build for local testing
dev:
	gox \
		-os="`go env GOOS`" \
		-arch="`go env GOARCH`" \
		-output "${GOPATH}/bin/pit" \
		./...