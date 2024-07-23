# go-base-service


## Description

## Generating compiled GRPC files
Compiled grpc files should be generated using the following command:

[//]: # (```protoc --go_out=. --go-grpc_out=. proto/*.proto```)
```protoc --proto_path=. --micro_out=. --go_out=. proto/*.proto```

They should be generated everytime a new proto client file is added, or everytime the proto file of the exposed GRPC endpoints of this service is modified.

