help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: ## Generate grpc files
	protoc --go_out=. --go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		proto/hello.proto

deploy-hello-server: ## Builds grpc server as Docker Image on port 50051
	docker build . -f server/Dockerfile -t tlgevers/grpc-hello-server
	docker push tlgevers/grpc-hello-server:latest

deploy-hello-client: ## Builds grpc client as Docker Image on port 50051
	docker build . -f client/Dockerfile -t tlgevers/grpc-hello-client
	docker push tlgevers/grpc-hello-client:latest
