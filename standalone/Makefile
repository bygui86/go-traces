
# VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-jaeger :		## Start Jaeger
	docker run -d --rm --name jaeger \
		-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
		-p 5775:5775/udp \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 14268:14268 \
		-p 14250:14250 \
		-p 9411:9411 \
		jaegertracing/all-in-one:1.21.0

stop-jaeger :		## Stop Jaeger
	docker container stop jaeger


## application

build :		## Build application
	go build

run :		## Run application from source code
	godotenv -f local.env go run main.go


## containerisation

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build : __check-container-tag		## Build container
	docker build . -t bygui86/standalone:$(CONTAINER_TAG)

container-push : __check-container-tag		## Push container to Docker hub
	docker push bygui86/standalone:$(CONTAINER_TAG)


## kubernetes

deploy :		## Deploy on Kubernetes
	kustomize build kube/ | kubectl apply -f -

delete :		## Delete from Kubernetes
	kustomize build kube/ | kubectl delete -f -

logs :		## Show logs from Kubernetes pod
	kubectl logs -l app=standalone -n apps --tail 10000 -f


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
