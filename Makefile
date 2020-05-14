
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

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

# TODO

### postgres

run-postgres :		## Run PostgreSQL in a container
	cd http-server/ && make run-postgres

stop-postgres :		## Stop PostgreSQL in container
	cd http-server/ && make stop-postgres

### kafka

run-kafka :		## Run Apache Zookeeper and Apache Kafka in containers
	cd kafka-producer/ && make run-kafka

stop-kafka :		## Stop Apache Zookeeper and Apache Kafka in containers
	cd kafka-producer/ && make stop-kafka

### kubemq

run-kubemq :		## Run Minikube and deploy KubeMQ
	cd kubemq-producer/ && make run-kubemq

stop-kubemq :		## Stop Minikube and KubeMQ
	cd kubemq-producer/ && make stop-kubemq

proxy-kubemq :		## Proxy KubeMQ
	cd kubemq-producer/ && make proxy-kubemq

open-kubemq-ui :		## Open KubeMQ UI in browser
	cd kubemq-producer/ && make open-kubemq-ui

## applications

### http

run-http-server :		## Run HTTP server application
	cd http-server/ && make run

run-http-client :		## Run HTTP client application
	cd http-client/ && make run

### kafka

run-kafka-consumer :		## Run Kafka consumer application
	cd kafka-consumer/ && make run

run-kafka-producer :		## Run Kafka producer application
	cd kafka-producer/ && make run

### kubemq

run-kubemq-consumer :		## Run KubeMQ consumer application
	cd kubemq-consumer/ && make run

run-kubemq-producer :		## Run KubeMQ producer application
	cd kubemq-producer/ && make run

### grpc

# TODO

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
