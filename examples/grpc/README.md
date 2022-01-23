```
go get google.golang.org/grpc
```

```
cd ..
protoc  grpc/tx.proto --go_out=plugins=grpc:grpc
```

## Run the server

```
cd server
go run server.go
```



## Run the Client

```
cd client
go run client.go
```