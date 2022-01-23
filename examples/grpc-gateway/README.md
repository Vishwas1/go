## gRPC gateway


> The gRPC-Gateway is a plugin of the Google protocol buffers compiler protoc. It reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC. 

![](https://github.com/grpc-ecosystem/grpc-gateway/raw/master/docs/assets/images/architecture_introduction_diagram.svg)


- https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/






## Compile proto

```
cd ..
protoc  grpc-gateway/tx.proto --go_out=plugins=grpc:grpc-gateway --go-grpc_out grpc-gateway

protoc -I grpc-gateway --go_out=plugins=grpc:grpc-gateway --go-grpc_out grpc-gateway --grpc-gateway_out grpc-gateway --grpc-gateway_opt paths=source_relative grpc-gateway/tx.proto 

```

## Run the server

```
cd server
go run server.go
```

## 

Ref: https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/

Now that we’ve got a working Go gRPC server, we need to add the gRPC-Gateway annotations.