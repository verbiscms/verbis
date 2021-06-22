VER=`cat VERSION`

build:
	go build -o verbisexec
.PHONY: build

serve:
	$(MAKE) build && ./verbisexec start
.PHONY: serve

build-prod:
	go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=true' -X 'github.com/ainsleyclark/verbis/api/version.Version=$(VER)'"
.PHONY: build-prod

version:
	echo $(VER)
.PHONY: version

release:
	./bin/release.sh
.PHONY: release

format:
	go fmt ./api/...
.PHONY: format

mock:
	cd api && rm -rf mocks && mockery --all --keeptree
.PHONY: mock

lint:
	golangci-lint run ./api/...
.PHONY: lint

test:
	go clean -testcache && go test -race $$(go list ./... | grep -v /res/ | grep -v /api/mocks/ | grep -v /build/ | grep -v /api/test)
.PHONY: test

test-v:
	go clean -testcache && go test -race $$(go list ./... | grep -v /res/ | grep -v /api/mocks/) -v
.PHONY: test-v

# Show to-do items per file.
todo:
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=node_modules \
		--exclude=Makefile \
		--exclude="*.map" \
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow' .
.PHONY: todo

all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
.PHONY: all