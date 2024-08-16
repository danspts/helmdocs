# Helmdocs

The idea behind this generator is to create `readme` and template `values.yaml` files from the `values.schema.json` of your helm charts

Other generators use the `values.yaml` files to generate the rest. I think this approach is wrong because it is always best to start with strong typed definitions and then generate weaker typed files. In this case, values.yaml is constricted by object, basic types and hooks, while `values.schema.json` defines a.... schema.

I hope this is of some use and happy helming to everyone!

## How to use it?

Let's use some values.schema.json (all rights to this [file](https://github.com/bitnami/charts/blob/a53f6eb/bitnami/redis/values.schema.json)) 

```yaml
{
    "$schema": "http://json-schema.org/schema#",
    "type": "object",
    "properties": {
      "architecture": {
        "type": "string",
        "title": "Redis architecture",
        "form": true,
        "description": "Allowed values: `standalone` or `replication`",
        "enum": ["standalone", "replication"]
      },
      "auth": {
        "type": "object",
        "title": "Authentication configuration",
        "form": true,
        "properties": {
          "enabled": {
            "type": "boolean",
            "form": true,
            "title": "Use password authentication"
          },
          "password": {
            "type": "string",
            "title": "Redis password",
            "form": true,
            "description": "Defaults to a random 10-character alphanumeric string if not set",
            "hidden": {
              "value": false,
              "path": "auth/enabled"
            }
          }
        }
      },
      "master": {
        "type": "object",
        "title": "Master replicas settings",
        "form": true,
        "properties": {
          "kind": {
            "type": "string",
            "title": "Workload Kind",
            "form": true,
            "description": "Allowed values: `Deployment`, `StatefulSet` or `DaemonSet`",
            "enum": ["Deployment", "StatefulSet", "DaemonSet"]
          },
          "persistence": {
            "type": "object",
            "title": "Persistence for master replicas",
            "form": true,
            "properties": {
              "enabled": {
                "type": "boolean",
                "form": true,
                "title": "Enable persistence",
                "description": "Enable persistence using Persistent Volume Claims"
              },
              "size": {
                "type": "string",
                "title": "Persistent Volume Size",
                "form": true,
                "render": "slider",
                "sliderMin": 1,
                "sliderMax": 100,
                "sliderUnit": "Gi",
                "hidden": {
                  "value": false,
                  "path": "master/persistence/enabled"
                }
              }
            }
          }
        }
      },
      "replica": {
        "type": "object",
        "title": "Redis replicas settings",
        "form": true,
        "hidden": {
          "value": "standalone",
          "path": "architecture"
        },
        "properties": {
          "kind": {
            "type": "string",
            "title": "Workload Kind",
            "form": true,
            "description": "Allowed values: `DaemonSet` or `StatefulSet`",
            "enum": ["DaemonSet", "StatefulSet"]
          },
          "replicaCount": {
            "type": "integer",
            "form": true,
            "title": "Number of Redis replicas"
          },
          "persistence": {
            "type": "object",
            "title": "Persistence for Redis replicas",
            "form": true,
            "properties": {
              "enabled": {
                "type": "boolean",
                "form": true,
                "title": "Enable persistence",
                "description": "Enable persistence using Persistent Volume Claims"
              },
              "size": {
                "type": "string",
                "title": "Persistent Volume Size",
                "form": true,
                "render": "slider",
                "sliderMin": 1,
                "sliderMax": 100,
                "sliderUnit": "Gi",
                "hidden": {
                  "value": false,
                  "path": "replica/persistence/enabled"
                }
              }
            }
          }
        }
      },
      "volumePermissions": {
        "type": "object",
        "properties": {
          "enabled": {
            "type": "boolean",
            "form": true,
            "title": "Enable Init Containers",
            "description": "Use an init container to set required folder permissions on the data volume before mounting it in the final destination"
          }
        }
      },
      "metrics": {
        "type": "object",
        "form": true,
        "title": "Prometheus metrics details",
        "properties": {
          "enabled": {
            "type": "boolean",
            "title": "Create Prometheus metrics exporter",
            "description": "Create a side-car container to expose Prometheus metrics",
            "form": true
          },
          "serviceMonitor": {
            "type": "object",
            "properties": {
              "enabled": {
                "type": "boolean",
                "title": "Create Prometheus Operator ServiceMonitor",
                "description": "Create a ServiceMonitor to track metrics using Prometheus Operator",
                "form": true,
                "hidden": {
                  "value": false,
                  "path": "metrics/enabled"
                }
              }
            }
          }
        }
      }
    }
  }
```

### Docker

#### To generate a readme use:
```
docker run -v $(pwd):/opt/tests ghcr.io/danspts/helmdocs:1.6.1 generate readme -schema-path /opt/tests/tests/redis/values.schema.json  -output  /opt/tests/tests/redis/README.md
```

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



#### To generate values use:
```
docker run -v $(pwd):/opt/tests ghcr.io/danspts/helmdocs:1.6.1 generate values -schema-path /opt/tests/tests/redis/values.schema.json  -output  /opt/tests/tests/redis/values.yaml
```

```yaml
architecture: <string> #  *[enum]* <details>"standalone", "replication"</details>

auth:
  enabled: <boolean>
  password: <string>

master:
  kind: <string> #  *[enum]* <details>"Deployment", "StatefulSet", "DaemonSet"</details>
  persistence:
    enabled: <boolean>
    size: <string> # (slider "1*Gi*\" ≤ "x*Gi*" ≤ "100*Gi*")

replica:
  kind: <string> #  *[enum]* <details>"DaemonSet", "StatefulSet"</details>
  replicaCount: <integer>
  persistence:
    enabled: <boolean>
    size: <string> # (slider "1*Gi*\" ≤ "x*Gi*" ≤ "100*Gi*")

volumePermissions:
  enabled: <boolean>

metrics:
  enabled: <boolean>
  serviceMonitor:
    enabled: <boolean>

```

### GHA

You can create a workflow:

```yaml
name: Generate Docs

on:
  workflow_dispatch:

jobs:
  generate:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Generate Values
        uses: ./.github/actions/generate-docs
        with:
          command: values
          schema_path: "tests/redis/values.schema.json"
          output_file: "tests/redis/values.yaml"
          omit_default: "false"
      
      - name: Generate README
        uses: ./.github/actions/generate-docs
        with:
          command: readme
          schema_path: "tests/redis/values.schema.json"
          output_file: "tests/redis/README.md"

```

If you want to fail the tests if things were not generated, you can run:

```yaml
- name: Fail if there are uncommitted changes
  run: git diff --exit-code
```