
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## K8S

### infra

start-minikube :		## Start Minikube
	minikube start --vm-driver=docker --cpus=6 --memory=12288 --disk-size=30g -p go-traces

stop-minikube :		## Stop Minikube
	minikube stop -p go-traces

delete-minikube :		## Stop Minikube
	minikube delete -p go-traces

deploy-all-infra :		## Deploy all infrastructure
	cd infrastructure/ && make deploy-all

delete-all-infra :		## Delete all infrastructure
	cd infrastructure/ && make delete-all

### apps

build-all-apps :		## Build all applications
	cd standalone/ && make build
	cd grpc-client/ && make build
	cd grpc-server/ && make build
	cd http-client/ && make build
	cd http-server/ && make build
	cd http-server-db/ && make build
	# cd kafka-consumer/ && make build
	# cd kafka-producer/ && make build
	cd kubemq-consumer/ && make build
	cd kubemq-producer/ && make build

__check-container-tag-all-apps :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build-all-apps : build-all-apps __check-container-tag-all-apps		## Build container for all applications
	cd standalone/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-client/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-server/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-client/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server-db/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	# cd kafka-consumer/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	# cd kafka-producer/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd kubemq-consumer/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd kubemq-producer/ && make container-build CONTAINER_TAG=$(CONTAINER_TAG)

container-push-all-apps : __check-container-tag-all-apps		## Push container of all applications to Docker hub
	cd standalone/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-client/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-server/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-client/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server-db/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	# cd kafka-consumer/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	# cd kafka-producer/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd kubemq-consumer/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd kubemq-producer/ && make container-push CONTAINER_TAG=$(CONTAINER_TAG)

deploy-all-apps :		## Deploy all applications to Kubernetes
	cd standalone/ && make deploy
	cd grpc-client/ && make deploy
	cd grpc-server/ && make deploy
	cd http-client/ && make deploy
	cd http-server/ && make deploy
	# cd http-server-db/ && make deploy
	cd kubemq-consumer/ && make deploy
	cd kubemq-producer/ && make deploy
	# cd kafka-consumer/ && make deploy
	# cd kafka-producer/ && make deploy

delete-all-apps :		## Deploy all applications from Kubernetes
	cd standalone/ && make delete
	cd grpc-client/ && make delete
	cd grpc-server/ && make delete
	cd http-client/ && make delete
	cd http-server/ && make delete
	# cd http-server-db/ && make delete
	cd kubemq-consumer/ && make delete
	cd kubemq-producer/ && make delete
	# cd kafka-consumer/ && make delete
	# cd kafka-producer/ && make delete

### ops

port-forw-http-client :		## Open port forwarding to http-client app
	kubectl port-forward svc/http-client 8080 -n apps

port-forw-grafana :		## Open port forwarding to Grafana
	cd infrastructure/ && make port-forw-grafana

port-forw-jaeger :		## Open port forwarding to Jaeger
	cd infrastructure/ && make port-forw-jaeger



## DOCKER

### jaeger

run-jaeger :		## Run Jaeger (all-in-one) in a container
	docker run -d --rm --name jaeger \
		-p 5775:5775/udp \
		-p 5778:5578 \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 9411:9411 \
		-p 14268:14268 \
		-p 14269:14269 \
		-p 16686:16686 \
		jaegertracing/all-in-one

stop-jaeger :		## Stop Jaeger (all-in-one) in container
	docker stop jaeger

open-jaeger-ui :		## Open Jaeger UI in browser
	open http://localhost:16686

### zipkin

run-zipkin :		## Run Zipkin in a container
	docker run -d --rm --name zipkin \
		-p 9411:9411 \
		openzipkin/zipkin

stop-zipkin :		## Stop Zipkin in container
	docker stop zipkin

open-zipkin-ui :		## Open Zipkin UI in browser
	open http://localhost:9411



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
