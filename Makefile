build:
	go build -o verbisexec

serve:
	$(MAKE) build && ./verbisexec start

build-prod:
	go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=true' -X 'github.com/ainsleyclark/api/version.Version=v0.0.3'" -tags prod

release:
	./bin/release.sh

format:
	go fmt ./api/...

mock:
	cd api && rm -rf mocks && mockery --all --keeptree

lint:
	golangci-lint run ./api/...

test:
	go clean -testcache && go test -race $$(go list ./... | grep -v /res/ | grep -v /api/mocks/ | grep -v /build/ | grep -v /api/test )

test-v:
	go clean -testcache && go test -race $$(go list ./... | grep -v /res/ | grep -v /api/mocks/) -v

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