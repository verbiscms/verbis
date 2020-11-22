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

mock:
	cd api && rm -rf mocks && mockery --all --keeptree

install:
	go install cms