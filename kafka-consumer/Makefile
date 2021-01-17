
# VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-kafka :		## Run Apache Zookeeper and Apache Kafka in containers
	docker swarm init
	docker stack deploy kafka -c kafka_docker-compose.yml

stop-kafka :		## Stop Apache Zookeeper and Apache Kafka in containers
	docker stack rm kafka
	@sleep 30
	docker swarm leave --force


## application

build :		## Build application
	go build

run :		## Run application from source code
	@export KAFKA_BOOTSTRAP_SERVERS=$$(ifconfig en0 | grep inet | grep -v inet6 | awk '{print $$2}') && \
	godotenv -f local.env go run main.go


## containerisation

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build : __check-container-tag		## Build container
	docker build . -t bygui86/kafka-consumer:$(CONTAINER_TAG)

container-push : __check-container-tag		## Push container to Docker hub
	docker push bygui86/kafka-consumer:$(CONTAINER_TAG)


## kubernetes

deploy :		## Deploy on Kubernetes
	kustomize build kube/ | kubectl apply -f -

delete :		## Delete from Kubernetes
	kustomize build kube/ | kubectl delete -f -

logs :		## Show logs from Kubernetes pod
	kubectl logs -l app=kafka-consumer -n apps --tail 10000 -f


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
