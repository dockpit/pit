XC_ARCH = "amd64"
XC_OS = "darwin linux windows"

debug:
	go-bindata -o=server/ui/bin/assets_dev.go -tags=debug -pkg=uibin -ignore=server/ui/bin -ignore=server/ui/node_modules -debug=true server/ui/...
	go run -tags=debug main.go start 

release:
	rm -fr bin/*
	mkdir -p bin/
	@echo "Building..."
	go-bindata -o=server/ui/bin/assets_prod.go -tags=!debug -pkg=uibin -ignore=server/ui/bin server/ui/...
	gox \
	    -os=$(XC_OS) \
	    -arch=$(XC_ARCH) \
	    -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" \
	    -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}" \
	    .

publish-zip:
	rm -fr bin/dist
	mkdir -p bin/dist
	for FOLDER in ./bin/*_* ; do \
		NAME=`basename $$FOLDER`_`CAT VERSION` ; \
		ARCHIVE=bin/dist/$$NAME.zip ; \
		pushd $$FOLDER ; \
		echo Zipping: $$FOLDER... `pwd` ; \
		zip ../dist/$$NAME.zip ./* ; \
		popd ; \
	done 

publish-release:
	github-release release \
    --user dockpit \
    --repo pit \
    --tag v$(shell cat VERSION) \
    --pre-release

publish-upload:
	for FOLDER in ./bin/*_* ; do \
		NAME=`basename $$FOLDER`_`CAT VERSION` ; \
		ARCHIVE=bin/dist/$$NAME.zip ; \
		github-release upload \
	    --user dockpit \
	    --repo pit \
	    --tag v$(shell cat VERSION) \
	    --name $$NAME.zip \
	    --file $$ARCHIVE ; \
	done 

publish-tag:
	git tag v$(shell cat VERSION)
	git push --tags