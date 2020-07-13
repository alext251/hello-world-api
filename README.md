# hello-world-api

## Prerequisites
This section contains the list of prerequisites as well as installation information for each.

1. Go version 1.14.  To install download the latest Go 1.14 package for your machine's OS at https://golang.org/dl/.  Then follow the steps for your OS and installation method at https://golang.org/doc/install to unpackage the binary and make executable on your machine.
2. Docker version 19.03.  Follow the instructions at https://docs.docker.com/get-docker/ to install docker for your machine's OS.  


## Build Instructions
Using the provided Dockerfile to build image `hello-world:latest`, run the command in the project root directory:
```
docker build . -t hello-world
```

## Run Instructions
Run the API container with `docker run`.  The API listens on port 8080 in the container.  To access the endpoints you need to specify a port on your host machine to forward to port 8080 on the container by adding `-p <HOST PORT>:8080` to the `docker run` command. From here on out, the instructions assume port 8080 is also being used on the host machine.   
```
docker run -itd -p 8080:8080 --name hello-world-api hello-world
```

## Debug Logging
Debug logs will look like the following `2020/07/13 01:50:43 - localhost:8080/`.

Debug logging can be enabled upon container startup.  To enable logging use the following command:
```
docker run -itd -p 8080:8080 --name hello-world-api hello-world api --debug
```

To access the logs run: 
```
docker logs hello-world-api
```

## Testing
Here are some example curls to test and their expected results:
1. `curl localhost:8080/` returns `<p>Hello, World</p>`
2. `curl -H 'Accept: application/json' localhost:8080/` returns `{"message": "Hello, World"}`
3. `curl -H 'Accept: image/webp' localhost:8080/` returns `<p>Hello, World</p>`
4. `curl -X POST localhost:8080/` returns `<p>Hello, World</p>`

To run the go API tests in `web/web_test.go` the options are to call:
1. `go test -v ./...` from the project root directory
2. `go test -v ./web` from the project root directory
3. `go test -v` from the `web/` directory

## Teardown Instructions
To tear down API container execute command:
```
    docker stop hello-world-api && docker rm hello-world-api
```
