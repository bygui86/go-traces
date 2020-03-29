
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

## Links

### OpenTracing
- https://opentracing.io/
- https://github.com/opentracing/opentracing-go
- https://github.com/opentracing-contrib/examples/

### Jaeger
- https://www.jaegertracing.io/
- https://github.com/jaegertracing/jaeger-client-go
- https://github.com/jaegertracing/jaeger-client-go/blob/master/config/example_test.go
- https://github.com/alextanhongpin/go-jaeger-trace

### TBD
- https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e
    - https://github.com/jfeng45/grpcservice

### OpenTelemetry
- https://opentelemetry.io/
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a
