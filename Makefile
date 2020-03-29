
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

build :		## Build application
	export GO111MODULE=on && \
	go build

run :		## Run application from source code
	export GO111MODULE=on && \
	godotenv -f local.env go run main.go

run-binary : build		## Run application from binary
	source local-binary.env && \
	./go-traces

run-postgres :		## Run PostgreSQL in a container
	docker run -d --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres

run-jaeger :		## Run Jaeger in a container
	docker run -d --name jaeger \
		-p 5775:5775/udp \
		-p 5778:5578 \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 9411:9411 \
		-p 14268:14268 \
		-p 16686:16686 \
		jaegertracing/all-in-one

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
