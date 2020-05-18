module github.com/bygui86/go-traces/http-server

go 1.14

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	contrib.go.opencensus.io/integrations/ocsql v0.1.5
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.3.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.5.1
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.opencensus.io v0.22.3
	go.uber.org/zap v1.14.1
)
