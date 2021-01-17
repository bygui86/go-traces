
# VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

start-minikube-with-kubemq :		## Start Minikube and deploy KubeMQ
	minikube start --vm-driver=docker --cpus=4 --memory=8192 --disk-size=20g -p kubemq
	@sleep 30s
	kubectl apply -f https://get.kubemq.io/deploy
	@sleep 30s
	kubectl get kubemqclusters,kubemqdashboards -n kubemq

stop-minikube :		## Stop Minikube and KubeMQ
	minikube stop -p kubemq

delete-minikube :		## Stop Minikube and KubeMQ
	minikube delete -p kubemq

proxy-kubemq :		## Proxy KubeMQ
	kubemqctl set cluster proxy

open-kubemq-ui :		## Open KubeMQ UI in browser
	kubemqctl get dashboard


## application

build :		## Build application
	go build

run :		## Run application from source code
	godotenv -f local.env go run main.go


## containerisation

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build : __check-container-tag		## Build container
	docker build . -t bygui86/kubemq-consumer:$(CONTAINER_TAG) --no-cache

container-push : __check-container-tag		## Push container to Docker hub
	docker push bygui86/kubemq-consumer:$(CONTAINER_TAG)


## kubernetes

deploy :		## Deploy on Kubernetes
	kustomize build kube/ | kubectl apply -f -

delete :		## Delete from Kubernetes
	kustomize build kube/ | kubectl delete -f -

logs :		## Show logs from Kubernetes pod
	kubectl logs -l app=kubemq-producer -n apps --tail 10000 -f


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
