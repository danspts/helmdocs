| Name | Title | Type | Required | Default |
|------|-------|------|----------|---------|
| `metrics` | Prometheus metrics details | **object** | false |  |
| `metrics.serviceMonitor` |  | **object** | false |  |
| `metrics.serviceMonitor.enabled` | Create Prometheus Operator ServiceMonitor | **boolean** | false |  |
| `metrics.enabled` | Create Prometheus metrics exporter | **boolean** | false |  |
| `architecture` | Redis architecture | **string** *[enum]* <details>"standalone", "replication"</details> | false |  |
| `auth` | Authentication configuration | **object** | false |  |
| `auth.enabled` | Use password authentication | **boolean** | false |  |
| `auth.password` | Redis password | **string** | false |  |
| `master` | Master replicas settings | **object** | false |  |
| `master.kind` | Workload Kind | **string** *[enum]* <details>"Deployment", "StatefulSet", "DaemonSet"</details> | false |  |
| `master.persistence` | Persistence for master replicas | **object** | false |  |
| `master.persistence.enabled` | Enable persistence | **boolean** | false |  |
| `master.persistence.size` | Persistent Volume Size | **string**(slider "1*Gi*\" ≤ "x*Gi*" ≤ "100*Gi*") | false |  |
| `replica` | Redis replicas settings | **object** | false |  |
| `replica.replicaCount` | Number of Redis replicas | **integer** | false |  |
| `replica.persistence` | Persistence for Redis replicas | **object** | false |  |
| `replica.persistence.enabled` | Enable persistence | **boolean** | false |  |
| `replica.persistence.size` | Persistent Volume Size | **string**(slider "1*Gi*\" ≤ "x*Gi*" ≤ "100*Gi*") | false |  |
| `replica.kind` | Workload Kind | **string** *[enum]* <details>"DaemonSet", "StatefulSet"</details> | false |  |
| `volumePermissions` |  | **object** | false |  |
| `volumePermissions.enabled` | Enable Init Containers | **boolean** | false |  |
