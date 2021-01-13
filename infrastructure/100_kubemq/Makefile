
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

deploy-crds :		## Deploy CRDs
	kustomize build 10_crds/ | kubectl apply -f -

deploy-operator :		## Deploy KubeMQ Operator
	kustomize build 20_operator/ | kubectl apply -f -

deploy-cluster :		## Deploy KubeMQ cluster
	kustomize build 30_cluster/ | kubectl apply -f -

deploy-conn :		## Deploy KubeMQ connectors
	kustomize build 40_connectors/ | kubectl apply -f -

deploy-dash :		## Deploy KubeMQ dashboard
	kustomize build 50_dashboard/ | kubectl apply -f -

deploy-all : deploy-crds deploy-operator deploy-cluster		## Deploy whole KubeMQ environment


## delete

delete-crds :		## Delete CRDs
	kustomize build 10_crds/ | kubectl delete -f -

delete-operator :		## Delete KubeMQ Operator
	kustomize build 20_operator/ | kubectl delete -f -

delete-cluster :		## Delete KubeMQ cluster
	kustomize build 30_cluster/ | kubectl delete -f -

delete-conn :		## Delete KubeMQ connectors
	kustomize build 40_connectors/ | kubectl delete -f -

delete-dash :		## Delete KubeMQ dashboard
	kustomize build 50_dashboard/ | kubectl delete -f -

delete-all : delete-cluster delete-operator delete-crds		## Delete whole KubeMQ environment


## ops

port-forw-prom :		## Open port forwarding to KubeMQ
	@kubectl port-forward svc/kubemq 32000 -n kubemq


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
	@echo ""
