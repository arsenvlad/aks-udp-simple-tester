FROM golang:1.10.1-alpine

# Add all source code
COPY . ./src/udp/

# Indicate the binary as our entrypoint
ENTRYPOINT go run ./src/udp/server.go

# Expose your port
EXPOSE 10001