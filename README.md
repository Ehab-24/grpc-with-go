# grpc-with-go
A simple grpc implementation in golang

## Prerequisites
 - You must have any of the [three latest releases](https://go.dev/doc/devel/release) of Go.
 - You must have the [proto compiler](https://grpc.io/docs/protoc-installation/) installed.

## Running Locally
1. Clone the project:
```
git clone https://github.com/Ehab-24/grpc-with-go.git
```
2. Optionally run:
```
go mod tidy
```
3. Start the server:
```
go run server/server.go
```
4. Start the client in another terminal:
```
go run client/client.go
```

## Hitting the API with Postman:
If you prefer, you can invoke any of the implemented functions using **Postman**.
1. Open a new tab in Postman and switch the _request type_ to _grpc_.
2. Import the `inventory/inventory.proto` into Postman.
4. Select a function to invoke from the URL bar
5. Press **ctrl+Enter** or click the **invoke** button.

## Changing the .proto file
If you wish to change the `inventory.proto` file, you must regenerate `inventory_grpc.pb.go` and `inventory.pb.go`. To do this:
1. Update your PATH so that _protoc_ compiler knows about the plugins:
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
2. Make and save changes to the `.proto` file.
3. Regenerate `.pb.go` files. While still in `grpc-with-go/`, run:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    inventory/inventory.proto
```
