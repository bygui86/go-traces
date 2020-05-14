
# VARIABLES
# -


# ENVIRONMENT VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

#run-kubemq :		## Run Minikube and deploy KubeMQ
# TODO start minikube
# TODO wait 30sec
# TODO deploy kubemq

#stop-kubemq :		## Stop Minikube and KubeMQ
# TODO stop minikube

#proxy-kubemq :		## Proxy KubeMQ
# TODO proxy kubemq

## application

build :		## Build application
	go build

run :		## Run application from source code
	godotenv -f local.env go run main.go

run-binary : build		## Run application from binary
	source local-binary.env && ./kubemq-consumer

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
