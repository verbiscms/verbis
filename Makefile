build:
	rm -rf verbisexec && go build -o verbisexec

serve:
	$(MAKE) build && ./verbisexec start

build-prod:
	rm -rf verbisexec && go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=true'" -tags prod

run:
	go run ./main.go

live-serve:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run serve

live-test:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run test

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
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow' .
.PHONY: todo

all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test