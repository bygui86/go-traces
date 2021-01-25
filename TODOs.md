
# TODOs

## Goals

- [ ] integrations
  - [ ] `integrate traces from jaeger and logs from loki in grafana`
- [ ] vector
  - [ ] grafana dashboard
- [ ] kubemq
  - [ ] grafana dashboard

## Apps

| App             | code | metrics | logs | traces | dockerfile | k8s manifests | k8s probes | status |
|-----------------|------|---------|------|--------|------------|---------------|------------|--------|
| standalone      | ok   | todo    | todo | ok     | ok         | ok            | todo       | ready  |
| grpc-server     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| grpc-client     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-server-db  | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-server     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-client     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| kubemq-producer | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| kubemq-consumer | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| kafka-producer  | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| kafka-consumer  | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |

### Special tasks

- [ ] add k8s probes on each app

## Infrastructure

- [ ] loki + grafana
  - [ ] config traces integration with jaeger
    - https://grafana.com/docs/grafana/latest/datasources/jaeger/#linking-trace-id-from-logs
    - https://grafana.com/docs/grafana/latest/datasources/loki/#derived-fields

- [ ] kubemq
  - [ ] grafana dashboard

- [ ] vector
  - [ ] grafana dashboard
