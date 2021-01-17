
# TODOs

## Goals

- [ ] logging
  - [ ] send logs to loki using promtail
  - [ ] send logs to loki using vector
  - [ ] see logs in grafana from loki
- [ ] brokers
  - [x] ~~deploy kubemq~~
  - [x] ~~scrape kubemq metrics~~
  - [ ] kubemq grafana dashboard
- [ ] integrations
  - [ ] `integrate traces from jaeger and logs from loki in grafana`

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

- [x] ~~split http-server into in-memory-db and db~~
- [x] ~~replace all `host = locahost` with `host = 0.0.0.0` as default config~~
- [ ] add k8s probes on each app

## Infrastructure

- [ ] kubemq
  - [ ] grafana dashboard

- [ ] promtail
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] prometheus metrics
  - [ ] grafana dashboard

- [ ] vector
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] prometheus metrics
  - [ ] grafana dashboard

- [ ] loki + grafana
  - [ ] config traces integration with jaeger
  - [ ] config traces integration with tempo

- [ ] tempo
  - [ ] retry sending traces to tempo
  - [ ] integration with grafana
  - [ ] integration with loki
  - [ ] grafana dashboards check
