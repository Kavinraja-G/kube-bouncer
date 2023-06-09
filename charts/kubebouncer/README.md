# kubebouncer

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.1.0](https://img.shields.io/badge/AppVersion-0.1.0-informational?style=flat-square)

A Helm chart for Kubebouncer

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | PodAffinities if any |
| envVars | list | `[{"name":"DENY_NAMESPACES","value":""}]` | Environment vars used in the webhook server deployment |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"ghcr.io/kavinraja-g/kubebouncer"` | Image repository to be used for kubebouncer |
| image.tag | string | `"0.1.0"` | Overrides the image tag whose default is the chart appVersion. |
| imagePullSecrets | list | `[]` | ImagePullSecrets if the image is from a private registry |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` | NodeSelectors used for Pod placements |
| podAnnotations | object | `{}` | Annotations for the pods |
| replicaCount | int | `1` | Replica Count used for the kubebouncer deployment |
| resources | object | `{}` | Resource `requests` & `limits` for the server |
| service | object | `{"port":8443,"type":"ClusterIP"}` | Configure service to be used for webhooks |
| service.port | int | `8443` | port defaults to `8443` must be the same as the server listens |
| tlsSecret | string | `"kubebouncer-tls-secret"` | Secret with TLS certificates used for webhooks |
| tolerations | list | `[]` | Tolerations used for Pod placements in the node if they have taints configured |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
