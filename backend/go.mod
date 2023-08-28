module github.com/DanielYevelkin/clutch-custom-gateway/backend

go 1.16

require (
	github.com/bufbuild/buf v0.56.0
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/fullstorydev/grpcurl v1.8.7
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/lyft/clutch/backend v0.0.0-20230821181314-3719e34d057e
	github.com/lyft/clutch/tools v0.0.0-20230821181314-3719e34d057e
	github.com/shurcooL/vfsgen v0.0.0-20230704071429-0000e147ea92
	github.com/stretchr/testify v1.8.4
	github.com/uber-go/tally/v4 v4.1.6
	go.uber.org/zap v1.25.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230815205213-6bfd019c3878
	google.golang.org/grpc v1.57.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
)
