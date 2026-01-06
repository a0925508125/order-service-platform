PROTO_DIRS = /proto order-service/proto product-service/proto user-service/proto

proto:
	@for dir in $(PROTO_DIRS); do \
		echo "Compiling $$dir ..."; \
		protoc --go_out=. --go-grpc_out=. -I=/proto -I=$$dir $$dir/*.proto; \
	done
