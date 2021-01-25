
# Go traces

## Description

Some simple applications to explore how to push traces in various context:

- http
- grpc
- broker

This is still a work-in-progress, [here](TODOs.md) the TODO list.

---

## How to start

### 1. Kubernetes cluster (Minikube)

```bash
make start-minikube
```

### 2. Infrastructure

```bash
make deploy-all-infra
```

### 3. Applications

```bash
make deploy-all-apps
```

### 4. Observe

In a terminal port-forward Grafana

```bash
make port-forw-grafana
```

Go to `http://localhost:3000` in the browser to access Grafana dashboards

Credentials:

`username`: admin
`password`: secret

### 5. Generate HTTP applications traces

1. Open port forwarding to `http-client`

1. Use the [Postman](https://www.postman.com/) [provided](postman/) to make some REST requests

### 6. Cleanup

```bash
make stop-minikube delete-minikube
```

---

## Observability

### Tracing

Tracing technologies:

- `Jaeger`, working as expected
- `Zipkin`, tested only locally, not in Kubernetes
- `GrafanaTempo`, not able to receive traces from same applications using Jaeger library

All applications support both Jaeger and Zipkin.

Tracing configurations can be set through environment variables:

| EnvVar | Default | Available values |
| --- | --- | --- |
| ENABLE_TRACING | true | true, false |

All default Jaeger environment variables are fully supported transparently.

### Monitoring

Monitoring technologies:

- `node-exporter`
- `kube-state-metrics`
- `prometheus-adapter`
- `Prometheus`
- `Grafana`

All applications expose `:9090/metrics` endpoint.

Monitoring configurations can be set through environment variables:

| EnvVar | Default | Available values |
| --- | --- | --- |
| ENABLE_MONITORING | true | true, false |
| MONITOR_HOST | 0.0.0.0 | - |
| MONITOR_PORT | 9090 | - |

### Logging

Logging technologies:

- (`TBD`) `Promtail` or `Vector`
- `GrafanaLoki`

All applications use `go.uber.org/zap` as logging library.

All applications log per default to stdout.

Logging configurations can be set through environment variables:

| EnvVar | Default | Available values |
| --- | --- | --- |
| LOG_ENCODING | console | console, json |
| LOG_LEVEL | info | trace, debug, info, warn, error, fatal |

---

## Links

### OpenCensus
- https://github.com/opencensus-integrations/ocsql
- https://opencensus.io/stats/
- https://opencensus.io/exporters/supported-exporters/go/prometheus/

### OpenTracing
- https://opentracing.io/
- https://opentracing.io/guides/golang/
- https://github.com/opentracing/opentracing-go
- https://github.com/opentracing-contrib/examples/
- https://www.reddit.com/r/golang/comments/cyahcp/help_wanted_with_opentracing_inject_extract_using/

### OpenTelemetry
- https://opentelemetry.io/
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector

### Jaeger
- https://www.jaegertracing.io/
- https://www.jaegertracing.io/docs/1.21/troubleshooting/
- https://github.com/jaegertracing/jaeger-client-go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/config/example_test.go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/metrics/prometheus/metrics_test.go

### Zipkin
- https://github.com/openzipkin/zipkin-go
- https://github.com/openzipkin/zipkin-go/blob/master/example_httpserver_test.go
- https://github.com/openzipkin-contrib/zipkin-go-opentracing
- https://github.com/openzipkin/zipkin-dependencies
- https://medium.com/devthoughts/instrumenting-a-go-application-with-zipkin-b79cc858ac3e

### Internal propagation
- https://docs.lightstep.com/docs/go-add-spans

### Grafana Tempo
- https://grafana.com/docs/tempo/latest/

#### DB
- https://github.com/ExpansiveWorlds/instrumentedsql - `ADOPTED SOLUTION`
- https://medium.com/@bas.vanbeek/opencensus-and-go-database-sql-322a26be5cc5
- https://opencensus.io/codelabs/gosqlguide/#0
- https://github.com/opencensus-integrations/ocsql - `NOT WORKING`
- https://github.com/gchaincl/sqlhooks - `NOT TESTED`
- https://github.com/luna-duclos/instrumentedsql - `NOT TESTED`

#### HTTP tracing
- https://docs.lightstep.com/docs/go-add-spans
- https://github.com/alextanhongpin/go-jaeger-trace
- https://medium.com/opentracing/tracing-http-request-latency-in-go-with-opentracing-7cc1282a100a
- https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
- https://medium.com/@marcus.olsson/adding-context-and-options-to-your-go-client-package-244c4ad1231b

#### gRPC tracing
- https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e
- https://github.com/jfeng45/grpcservice
- https://github.com/opentracing-contrib/go-grpc

#### Broker tracing

##### Kafka
- https://github.com/confluentinc/confluent-kafka-go
- https://github.com/jaegertracing/jaeger/blob/master/pkg/kafka/producer/config.go (shopify/sarama kafka library)
##### KubeMQ
- https://github.com/kubemq-io/kubemq-go
- https://github.com/kubemq-io/kubemq-go/blob/master/event_store.go
- https://github.com/kubemq-io/kubemq-go/blob/master/trace.go

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a

### Tools

- https://grafana.com/blog/2019/11/22/kubecon-demo-a-preview-of-grafana-jaeger/
- https://grafana.com/go/introduction-to-distributed-tracing

### Integration with logs (Grafana Loki)

- https://grafana.com/docs/grafana/latest/datasources/loki/#derived-fields
- https://grafana.com/docs/tempo/latest/guides/loki-derived-fields/
