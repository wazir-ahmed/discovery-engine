CURDIR=$(shell pwd)

.PHONY: build
build:
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/config/config.proto
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/worker/worker.proto
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/consumer/consumer.proto
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/analyzer/analyzer.proto
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/insight/insight.proto
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative v1/observability/observability.proto

