build:
	go build -o verbis

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

test:
	go test $$(go list ./... | grep -v /res/ | grep -v /api/mocks/)

test-v:
	go test $$(go list ./... | grep -v /res/ | grep -v /api/mocks/) -v

install:
	go install verbis