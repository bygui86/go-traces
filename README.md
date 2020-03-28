
# Go traces

Go sample project to expose traces

## Run

1. start PostgreSQL in a container
```shell script
make run-postgres
```

2. run application
```shell script
make run
```

---

## REST endpoints

- `GET /products` > Fetch a list of products in response to a valid 
- `GET /products/{id}` > Fetch a product in response to a valid 
- `POST /products` > Create a new product in response to a valid 
- `PUT /products/{id}` > Update a product in response to a valid 
- `DELETE /products/{id}` > Delete a product in response to a valid 

---

## Links

### OpenTracing
- https://opentracing.io/
- https://github.com/opentracing/opentracing-go
- https://github.com/opentracing-contrib/examples/

### OpenTelemetry
- https://opentelemetry.io/
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/open-telemetry/opentelemetry-collector

### TBD
- https://github.com/alextanhongpin/go-jaeger-trace
- https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e
    - https://github.com/jfeng45/grpcservice

### Kubernetes
- https://medium.com/opentracing/opentracing-on-kubernetes-get-yer-tracing-for-free-7a69cca03c8a
