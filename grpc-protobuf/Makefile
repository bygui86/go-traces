
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

	# protoc --proto_path=./grpc-protobuf/ --go_out=plugins=grpc:./grpc-protobuf ./grpc-protobuf/*
build :		## Compile protobuf
	rm -f ./*.pb.go
	rm -f ../grpc-server/grpc_interface/*.pb.go
	rm -f ../grpc-client/grpc_interface/*.pb.go
	protoc -I=./ --go_out=plugins=grpc:./ ./*.proto
	cp ./*.pb.go ../grpc-server/grpc_interface/
	cp ./*.pb.go ../grpc-client/grpc_interface/


## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
