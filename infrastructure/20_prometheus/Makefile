
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

deploy-crds :		## Deploy CRDs
	kustomize build 10_crds/ | kubectl apply -f -

deploy-operator :		## Deploy Prometheus Operator
	kustomize build 20_operator/ | kubectl apply -f -

deploy-kube-state :		## Deploy Kube-State-Metrics
	kustomize build 30_kube-state-metrics/ | kubectl apply -f -

deploy-node-exp :		## Deploy Node-Exporter
	kustomize build 40_node-exporter/ | kubectl apply -f -

deploy-prom-adapt :		## Deploy Prometheus-Adapter
	kustomize build 50_prometheus-adapter/ | kubectl apply -f -

deploy-ns-rbac :		## Deploy RBAC for all namespaces
	kustomize build 60_ns-rbac/ | kubectl apply -f -

deploy-prom :		## Deploy Prometheus
	kustomize build 70_prometheus/ | kubectl apply -f -

deploy-all : deploy-crds deploy-operator deploy-kube-state deploy-node-exp deploy-prom-adapt deploy-ns-rbac deploy-prom		## Deploy whole Prometheus environment


## delete

delete-crds :		## Delete CRDs
	kustomize build 10_crds/ | kubectl delete -f -

delete-operator :		## Delete Prometheus Operator
	kustomize build 20_operator/ | kubectl delete -f -

delete-prom-adapt :		## Delete Prometheus-Adapter
	kustomize build 50_prometheus-adapter/ | kubectl delete -f -

delete-node-exp :		## Delete Node-Exporter
	kustomize build 40_node-exporter/ | kubectl delete -f -

delete-kube-state :		## Delete Kube-State-Metrics
	kustomize build 30_kube-state-metrics/ | kubectl delete -f -

delete-ns-rbac :		## Delete RBAC for all namespaces
	kustomize build 60_ns-rbac/ | kubectl delete -f -

delete-prom :		## Delete Prometheus
	kustomize build 70_prometheus/ | kubectl delete -f -

delete-all : delete-prom delete-ns-rbac delete-kube-state delete-node-exp delete-prom-adapt delete-operator delete-crds		## Delete whole Prometheus environment


## pvc

delete-pvc :		## Delete Kubernetes PersistentVolumeClaim for Prometheus environment
	kubectl delete pvc -l app=prometheus


## ops

port-forw-prom :		## Open port forwarding to Prometheus
	@kubectl port-forward svc/prometheus 9090 -n monitoring


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
