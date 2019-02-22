# Gorm / mux sample rest api

## Instructions

### Dependencies

run `go get github.com/davide-c/rest-sample-b`

### Run the server

cd to project's folder then:

run `docker-compose up`
run `go run main.go`

It listens on port `3000`

### Hit the api with curl

run `curl 127.0.0.1:3000/properties/`

### Available endpoints:

- `/properties/ [GET]`
- `/property/ [POST]`
- `/property/{title} [GET]`
- `/assets/ [GET]`
- `/asset/ [POST]`
