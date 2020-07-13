# hello-world-api

## Prerequisites

1. Go version 1.14
2. Docker version 19.03.8
3. Port 8080 open on host machine

## Build Instructions
Using the provided Dockerfile to build `hello-world:latest`, run the command in the project root directory:
```
docker build . -t hello-world
```
## Run Instructions
Run using the `docker run` command below:
```
docker run -itd -p 8080:8080 --name hello-world-api hello-world
```

## Debug Logging
Debug logs will look like the following `2020/07/13 01:50:43 - localhost:8080/`.

Debug logging can be enabled upon container startup.  To enable logging use the following command:
```
docker run -itd -p 8080:8080 --name hello-world-api hello-world api -d
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
