
# Go traces

Go sample project to expose traces

## Run

start Jaeger
```shell script
make run-jaeger
```

### HTTP applications

1. start PostgreSQL
    ```shell script
    make run-postgres
    ```

2. start server
    ```shell script
    make run-http-server
    ```

3. start client
    ```shell script
    make run-http-client
    ```

4. play a bit with [Postman](https://www.postman.com/) loading the prepared collections
    - [HTTP server collection](http-server/postman/postman_collection.json)
    - [HTTP client collection](http-client/postman/postman_collection.json)

### gRPC applications

`TODO`

### Broker applications

`TODO`

---

## Examples

- [x] rest server with database
- [x] logging adoption
- [x] monitoring adoption
- [x] tracing adoption
- [x] simple traces example
- [x] traces with logging
- [x] traces with monitoring
- [x] internal span propagation example
- [ ] external span propagation example
    - [ ] db tracer example - `NOT WORKING`
    - [x] http client/server example
    - [ ] kafka example - `WIP`
    - [ ] kubemq example
    - [ ] grpc example
    - [ ] `TBD` http client/middleware/server example
    - [ ] `TBD` ask google example
- [x] send spans to jaeger
- [ ] send spans to zipkin

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
- https://github.com/jaegertracing/jaeger-client-go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/config/example_test.go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/metrics/prometheus/metrics_test.go

### Zipkin
- https://github.com/openzipkin/zipkin-go

### Internal propagation
- https://docs.lightstep.com/docs/go-add-spans

#### DB
- https://medium.com/@bas.vanbeek/opencensus-and-go-database-sql-322a26be5cc5
- https://opencensus.io/codelabs/gosqlguide/#0
- https://github.com/opencensus-integrations/ocsql
- https://github.com/gchaincl/sqlhooks

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
`TO BE SEARCHED`

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a
