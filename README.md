# Go tutorial - gRPC

This is a simple example of a gRPC server and client in Go. It is based on a course from Udemy.

## Generate the code

To generate the code, run the following command:

```bash
protoc -I ./greet/proto --go_out=. --go_opt=module=github.com/vergissberlin/go-tutorial-grcp/greet/proto --go-grpc_out=. --go-grpc_opt=module=github.com/vergissberlin/go-tutorial-grcp/greet/proto ./greet/proto/dummy.proto
```
