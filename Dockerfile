#build stage
FROM golang:1.16 AS build


RUN mkdir /verbis
WORKDIR /verbis

# Add Maintainer Info
LABEL maintainer="Ainsley Clark <info@ainsleyclark.com>"

# Copy go mod and sum files
#COPY go.mod go.sum ./

ADD . /verbis

# Copy the source from the current directory to the Working Directory inside the container
#COPY . .

RUN go test -v ./api/...

ENV CGO_ENABLED=0

RUN GOOS=darwin GOARCH=amd64 go build -o "test"
