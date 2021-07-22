VER=`cat VERSION`

# Build
build:
	go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=false' -X 'github.com/ainsleyclark/verbis/api/version.Version=$(VER)'"
.PHONY: build

# Set Verbis up when cloned.
setup:
	go mod tidy
	cd admin && npm install
	cd admin && npm run build
.PHONY: setup

# Builds and serves
serve:
	$(MAKE) build && ./verbisexec start
.PHONY: serve

# Builds Verbis for production
build-prod:
	go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=true' -X 'github.com/ainsleyclark/verbis/api/version.Version=$(VER)'"
.PHONY: build-prod

# Creates and build dist folder
dist:
	goreleaser release --rm-dist --snapshot
.PHONY: dist

# Echo the current versnion
version:
	echo $(VER)
.PHONY: version

# Release verbis, expects commit message
release:
	./bin/release.sh
.PHONY: release

# Run gofmt in ./api...
format:
	go fmt ./api/...
.PHONY: format

# Test uses race and coverage
test:
	go clean -testcache && go test -race $$(go list ./... | grep -v /res/ | grep -v /api/mocks/ | grep -v /build/ | grep -v /api/test | grep -v /api/importer) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

# Test with -v
test-v:
	go clean -testcache && go test -race -v $$(go list ./... | grep -v /res/ | grep -v /api/mocks/ | grep -v /build/ | grep -v /api/test | grep -v /api/importer) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.out
.PHONY: cover

# Github Actions
ci:
	rm -rf admin/dist
	mkdir admin/dist
	touch admin/dist/.gitkeep
	$(MAKE) format
	$(MAKE) test
.PHONY: ci

# Make mocks keeping directory tree
mock:
	cd api && rm -rf mocks && mockery --all --keeptree
.PHONY: mock

# Run linter
lint:
	golangci-lint run ./api/...
.PHONY: lint

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

# Make format, lint and test
all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
.PHONY: all
