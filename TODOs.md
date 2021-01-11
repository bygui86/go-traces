
# TODOs

## Apps

| App             | code | metrics | logs | traces | dockerfile | k8s manifests | k8s probes | status |
|-----------------|------|---------|------|--------|------------|---------------|------------|--------|
| grpc-protobuf   | ok   | -       | -    | -      | -          | -             | -          | ready  |
| grpc-server     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| grpc-client     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-server     | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| http-client     | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kubemq-producer | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kubemq-consumer | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kafka-producer  | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kafka-consumer  | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |

### Specific

- [ ] http-server - remove postgres dependency
- [ ] add k8s probes on each app

## Infrastructure

- [x] makefile

- [x] namespaces
  - [x] kustomize
  - [x] manifests
  - [x] labels

- [ ] prometheus
  - [ ] kustomize
  - [x] manifests
  - [x] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [x] metrics/dashboard

- [x] loki
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics/dashboard

- [ ] promtail
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] metrics/dashboard

- [ ] vector
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] metrics/dashboard

- [x] tempo
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics/dashboard

- [ ] jaeger
  - [ ] kustomize
  - [x] manifests
  - [ ] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [ ] config
  - [ ] metrics/dashboard

- [x] grafana
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] dashboards

- [ ] kubemq
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] metrics/dashboard
