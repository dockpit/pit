XC_ARCH = "amd64"
XC_OS = "darwin linux windows"

debug:
	go-bindata \
		-o=server/ui/bin/assets_dev.go \
		-tags=debug \
		-pkg=uibin \
		-ignore=server/ui/bin \
		-ignore=server/ui/node_modules \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/node_modules \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/src \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/examples \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/tasks \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/test \
		-debug=true server/ui/...
	go run -tags=debug -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" main.go start 

#1. build release binaries
release:
	rm -fr bin/*
	mkdir -p bin/
	@echo "Building..."
	go-bindata \
		-o=server/ui/bin/assets_prod.go \
		-tags=!debug \
		-pkg=uibin \
		-ignore=server/ui/bin \
		-ignore=server/ui/node_modules \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/node_modules \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/src \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/examples \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/tasks \
		-ignore=server/ui/vendor/semantic-ui-1.11.5/test \
		server/ui/...
	gox \
	    -os=$(XC_OS) \
	    -arch=$(XC_ARCH) \
	    -ldflags "-X main.Version `cat VERSION` -X main.Build `date -u +%Y%m%d%H%M%S`" \
	    -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}" \
	    .

#2. tag git commit
publish-tag:
	git tag v$(shell cat VERSION)
	git push --tags

#3 create github release
publish-release:
	github-release release \
    --user dockpit \
    --repo pit \
    --tag v$(shell cat VERSION) \
    --pre-release

#4 zip binaries
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

#5 create checksums of zip archives
publish-checksums:
	cd bin/dist && shasum -a256 * > ./dockpit_$(shell cat VERSION)_SHA256SUMS

#6 upload zip and checksums
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
	github-release upload \
	    --user dockpit \
	    --repo pit \
	    --tag v$(shell cat VERSION) \
	    --name dockpit_$(shell cat VERSION)_SHA256SUMS \
	    --file bin/dist/dockpit_$(shell cat VERSION)_SHA256SUMS

