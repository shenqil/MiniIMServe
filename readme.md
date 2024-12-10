# 1. protobuf 编译命令

```
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  protocol\protobuf\messages.proto
```