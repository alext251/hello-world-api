FROM golang:1.14

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go mod download \
 && go build -o api . \
 && cp api /go/bin

# Entrypoint is to start a new shell 
# with this either call 'api' for default api behavior
# or call 'api -d' to include debug loggdoing
CMD ["api"]
