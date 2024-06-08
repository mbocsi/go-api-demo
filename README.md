# Go API Demo

A RESTful API made using Go and the standard net/http package. This is a demo project for gathering experience in Go and backend development. Works with the api-demo-frontend project.

## Running the server (dev)

`go run cmd/api/main.go`

## Building a docker image

`docker build --platform linux/amd64 -t mbocsi/goapi .`

## Running the docker container

`docker run -p 8080:8080 mbocsi/goapi`
