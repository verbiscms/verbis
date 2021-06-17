build:
	rm -rf verbisexec && go build -o verbisexec -tags dev

build-run:
	$(MAKE) bld && ./verbisexec start

build-prod:
	rm -rf verbisexec && go build -o verbisexec -ldflags="-X 'github.com/ainsleyclark/verbis/api.ProductionString=true'" -tags prod

run:
	go run ./main.go

serve:
	go run ./main.go serve

install-serve:
	go install cms && cms serve

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

install:
	go install verbis

all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test