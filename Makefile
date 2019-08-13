build:
	protoc -I. consignment-service/proto/consignment/consignment.proto --micro_out=. --go_out=.
