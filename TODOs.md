
# TODOs

## BUGs

- [ ] jaeger service monitor not visible in prom/targets
- [ ] grpc-apps scrape failing

## Apps

| App             | code | metrics | logs | traces | dockerfile | k8s manifests | k8s probes | status |
|-----------------|------|---------|------|--------|------------|---------------|------------|--------|
| standalone      | ok   | todo    | todo | ok     | ok         | ok            | todo       | ready  |
| grpc-protobuf   | ok   | -       | -    | -      | -          | -             | -          | ready  |
| grpc-server     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| grpc-client     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-server-db  | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-server     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| http-client     | ok   | ok      | ok   | ok     | ok         | ok            | todo       | ready  |
| kubemq-producer | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kubemq-consumer | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kafka-producer  | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |
| kafka-consumer  | ok   | ok      | ok   | ok     | todo       | todo          | todo       | wip    |

### Special tasks

- [x] ~~split http-server into in-memory-db and db~~
- [ ] add k8s probes on each app
- [ ] replace all `host = locahost` with `host = 0.0.0.0` as default config 

## Infrastructure

- [x] ~~makefile~~

- [x] ~~namespaces~~
  - [x] kustomize
  - [x] manifests
  - [x] labels

- [x] ~~prometheus~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics
  - [x] dashboard

- [x] ~~loki~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics
  - [x] dashboard

- [ ] promtail
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] metrics
  - [ ] dashboard

- [ ] vector
  - [ ] kustomize
  - [ ] manifests
  - [ ] labels
  - [ ] resources
  - [ ] storage
  - [ ] affinity/tolerations
  - [ ] rbac/security-ctx
  - [ ] config
  - [ ] metrics
  - [ ] dashboard

- [x] ~~tempo~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics
  - [x] dashboard

- [x] ~~jaeger~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics
  - [x] dashboard

- [x] ~~grafana~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] dashboards

- [x] ~~kubemq~~
  - [x] kustomize
  - [x] manifests
  - [x] labels
  - [x] resources
  - [x] storage
  - [x] affinity/tolerations
  - [x] rbac/security-ctx
  - [x] config
  - [x] metrics
  - [x] dashboard
