
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

deploy-ns :		## Deploy namespaces
	kustomize build 10_namespaces/ | kubectl apply -f -

deploy-prom :		## Deploy Prometheus
	cd 20_prometheus && make deploy-all

deploy-loki :		## Deploy Loki
	kustomize build 30_loki/ | kubectl apply -f -

# deploy-promtail :		## Deploy Promtail
# 	kustomize build 31_promtail/ | kubectl apply -f -

# deploy-vector :		## Deploy Vector
# 	kustomize build 31_vector/ | kubectl apply -f -

deploy-jaeger :		## Deploy Jaeger
	cd 40_jaeger && make deploy-all

deploy-tempo :		## Deploy Tempo
	kustomize build 40_tempo/ | kubectl apply -f -

deploy-grafana :		## Deploy Grafana
	kustomize build 50_grafana/ | kubectl apply -f -

deploy-kubemq :		## Deploy KubeMQ
	cd 100_kubemq && make deploy-all

deploy-all : deploy-ns deploy-prom deploy-loki deploy-jaeger deploy-grafana		## Deploy all infrastructure


## delete

delete-ns :		## Delete namespaces
	kustomize build 10_namespaces/ | kubectl delete -f -

delete-prom :		## Delete Prometheus
	cd 20_prometheus && make delete-all

delete-loki :		## Delete Loki
	kustomize build 30_loki/ | kubectl delete -f -

# delete-promtail :		## Delete Promtail
# 	kustomize build 31_promtail/ | kubectl delete -f -

# delete-vector :		## Delete Vector
# 	kustomize build 31_vector/ | kubectl delete -f -

delete-jaeger :		## Delete Jaeger
	cd 40_jaeger && make delete-all

delete-tempo :		## Delete Tempo
	kustomize build 40_tempo/ | kubectl delete -f -

delete-grafana :		## Delete Grafana
	kustomize build 50_grafana/ | kubectl delete -f -

delete-kubemq :		## Delete KubeMQ
	cd 100_kubemq && make delete-all

delete-all : delete-grafana delete-jaeger delete-loki delete-prom		## Delete all infrastructure
	sleep 60
	@make delete-ns


## ops

port-forw-graf :		## Open port forwarding to Grafana
	@kubectl port-forward svc/grafana 3000 -n monitoring

port-forw-prom :		## Open port forwarding to Prometheus
	@kubectl port-forward svc/prometheus 9090 -n monitoring

port-forw-jaeger :		## Open port forwarding to Jaeger
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
