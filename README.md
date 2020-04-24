
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
- [x] external span propagation example
    - [x] db tracer example
    - [ ] http client/server example - `WIP`
    - [ ] http client/middleware/server example
    - [ ] grpc example
    - [ ] broker example
    - [ ] `TBD` ask google example

---

## Links

### OpenTracing
- https://opentracing.io/
- https://github.com/opentracing/opentracing-go
- https://github.com/opentracing-contrib/examples/

### Jaeger
- https://www.jaegertracing.io/
- https://github.com/jaegertracing/jaeger-client-go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/config/example_test.go
    - https://github.com/jaegertracing/jaeger-client-go/blob/master/metrics/prometheus/metrics_test.go

### Internal propagation
- https://docs.lightstep.com/docs/go-add-spans

### External propagation
- https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e

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
- https://github.com/jfeng45/grpcservice

### OpenTelemetry
- https://opentelemetry.io/
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a
