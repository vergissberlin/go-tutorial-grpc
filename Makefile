generate:
	protoc \
		-I greet/proto \
		--go_out=. \
		--go_opt=module=github.com/vergissberlin/go-tutorial-grpc \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/vergissberlin/go-tutorial-grpc \
		greet/proto/dummy.proto

default: generate
