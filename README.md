
# Go traces

Go sample project to expose traces

## Run

1. start PostgreSQL
```shell script
make run-postgres
```

2. start Jaeger
```shell script
make run-jaeger
```

3. start application
```shell script
make run
```

4. play a bit with [Postman](https://www.postman.com/) loading the [prepared collection](postman/postman_collection.json)

---

## Servers

### Monitoring

URL: `localhost:9090/metrics`

### Products endpoints

Root URL: `localhost:8080/`

| Method | URL | Description
| --- | --- | --- |
| GET | /products | Fetch list of products |
| GET | /products/{id} | Fetch a product by ID |
| POST | /products | Create a new product |
| PUT | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID |

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
- [ ] http client/server example
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

### Internal span propagation
- https://docs.lightstep.com/docs/go-add-spans

### HTTP tracing
- https://docs.lightstep.com/docs/go-add-spans
- https://github.com/alextanhongpin/go-jaeger-trace
- https://medium.com/opentracing/tracing-http-request-latency-in-go-with-opentracing-7cc1282a100a

### gRPC tracing
- https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e
    - https://github.com/jfeng45/grpcservice

### OpenTelemetry
- https://opentelemetry.io/
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a
