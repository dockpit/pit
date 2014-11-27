# cross compile for now only works for darwin (cgo stuff as usual)
# @see http://dominik.honnef.co/posts/2013/12/cross_compiling_go_and_hidden_uses_of_cgo/
# XC_ARCH = "386 amd64 arm"
# XC_OS = "linux darwin windows freebsd openbsd"

XC_ARCH = "amd64"
XC_OS = "darwin"
BT_API_KEY = $$BINTRAY_API_KEY

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
	docker build -t dockpit/pit:`cat VERSION` .
	docker tag dockpit/pit:`cat VERSION` dockpit/pit:latest

#build for local testing
dev:
	gox \
		-os="`go env GOOS`" \
		-arch="`go env GOARCH`" \
		-output "${GOPATH}/bin/pit" \
		./...

#publish version to bintray
publish:
	rm -fr bin/dist
	mkdir -p bin/dist
	for FOLDER in ./bin/*_* ; do \
		NAME=`basename $$FOLDER`_`CAT VERSION` ; \
		ARCHIVE=bin/dist/$$NAME.zip ; \
		echo Zipping: $$FOLDER... ; \
		zip bin/dist/$$NAME.zip $$FOLDER/* ; \
		curl -T $$ARCHIVE -uadvanderveer:$(BT_API_KEY) https://api.bintray.com/content/dockpit/pit/pit/`cat VERSION`/$$NAME.zip ; \
	done 

#run all unit tests
test:
	go test ./... 