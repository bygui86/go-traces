
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

deploy-crds :		## Deploy CRDs
	kustomize build 10_crds/ | kubectl apply -f -

deploy-operator :		## Deploy Jaeger Operator
	kustomize build 20_operator/ | kubectl apply -f -

deploy-jaeger :		## Deploy Jaeger
	kustomize build 30_jaeger/ | kubectl apply -f -

deploy-monitor :		## Deploy Jaeger monitoring
	kustomize build 40_monitoring/ | kubectl apply -f -

deploy-all : deploy-crds deploy-operator deploy-jaeger deploy-monitor		## Deploy whole Jaeger environment


## delete

delete-crds :		## Delete CRDs
	kustomize build 10_crds/ | kubectl delete -f -

delete-operator :		## Delete Jaeger Operator
	kustomize build 20_operator/ | kubectl delete -f -

delete-jaeger :		## Delete Jaeger
	kustomize build 30_jaeger/ | kubectl delete -f -

delete-monitor :		## Delete Jaeger monitoring
	kustomize build 40_monitoring/ | kubectl delete -f -

delete-all : delete-monitor delete-jaeger delete-operator delete-crds		## Delete whole Jaeger environment


## ops

port-forw-prom :		## Open port forwarding to Jaeger
	@kubectl port-forward svc/jaeger-query 16686 -n tracing


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
