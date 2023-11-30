# gRPC Routes App

## Overview

This gRPC-based application is designed to fetch the route information for a given IP destination of a Server and return it to the client. It's written in Go, utilizing the gRPC framework and Protocol Buffers.

## Purpose

The primary goal of this application is to demonstrate the usage of gRPC for communication between a client and server. The server calculates and returns the route details, including the interface, gateway, and source information for the specified IP destination.

## Prerequisites

Before running the application, ensure that you have the following installed:

- [Go](https://golang.org/doc/install)
- [Protocol Buffers](https://developers.google.com/protocol-buffers) (protoc)
- [gRPC Go](https://pkg.go.dev/google.golang.org/grpc)
- [Go netroute](https://pkg.go.dev/github.com/libp2p/go-netroute)

## How to Run

### Server

1. Navigate to the `route_server` directory.
2. Run the following commands to build and start the server:

    ```bash
    go build
    ./route_server
    ```

   The server will start listening on the specified port.

### Client

1. Navigate to the `route_client` directory.
2. Run the following commands to build and execute the client:

    ```bash
    go build
    ./route_client -addr localhost:8088 -destination 8.8.8.8
    ```

   Adjust the `-addr` and `-destination` flags as needed.

## gRPC Protobuf Definitions

The communication between the client and server is defined using Protocol Buffers (proto3). The `routes.proto` file specifies the service and message definitions.

### Protocol Buffers Compilation

Before running the application, ensure you have the protocol buffers compiled. Run the following commands included in the txt file:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routes/routes.proto
```

## Output example

Server output example:

```css
albert@264YL route_server % ./route_server
2023/11/30 17:16:31 server listening at [::]:8088
2023/11/30 17:16:39 Received route request to: 0.0.0.0
2023/11/30 17:16:49 Received route request to: 8.8.8.8
```

Client output example:
```css
albert@264YL route_client % ./route_client -addr localhost:8088 -destination 0.0.0.0
2023/11/30 17:16:39 Sending route information request for route: 0.0.0.0
2023/11/30 17:16:39 Response: Interface: en0 Gateway: 192.168.31.1 Source: 192.168.31.223
albert@264YL route_client % ./route_client -addr localhost:8088 -destination 8.8.8.8
2023/11/30 17:16:49 Sending route information request for route: 8.8.8.8
2023/11/30 17:16:49 Response: Interface: en0 Gateway: 192.168.31.1 Source: 192.168.31.223
```
