# kube-bouncer
KubeBouncer is a set of Kubernetes admission controller that denies resource deployment in specified namespaces, enforcing security policies and best practices in Kubernetes.

## Getting Started
Currently there are two sets of validation webhooks available as part of the kubebouncer:
#### NamespaceBouncer `/validate-namespace`
1. Checks if any requested resource action is done on a namespace provided in the webhook deployment environment variable `DENY_NAMESPACES`, and denies the request/action. To create a NamespaceBouncer webhook, you can use the provided YAML configuration and replace the variables with your own values.
  ```yaml
  ---
  apiVersion: admissionregistration.k8s.io/v1
  kind: ValidatingWebhookConfiguration
  metadata:
    name: nsbouncer-webhook
  webhooks:
  - name: ${FQDN_OF_THE_SERVICE}
    sideEffects: None
    admissionReviewVersions: ["v1"]
    timeoutSeconds: 5
    clientConfig:
      service:
        name: ${SERVICE_NAME}
        namespace: ${NAMESPACE}
        path: "/validate-namespace"
      caBundle: ${CA_BUNDLE} # Replace it with the ca.pem file which is used to generate the certificates and keys for the webhook
    rules:
      # rules can be configured based on the user requirements
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
  ```
#### PodBouncer `/validate-pods`
1. Checks if the `readinessProbes` or `livenessProbes` are present in the pod spec, and denies the request if they are not. To create a PodBouncer webhook, you can use the provided YAML configuration and replace the variables with your own values.
  ```yaml
  ---
  apiVersion: admissionregistration.k8s.io/v1
  kind: ValidatingWebhookConfiguration
  metadata:
    name: podbouncer-webhook
  webhooks:
  - name: ${FQDN_OF_THE_SERVICE}
    sideEffects: None
    admissionReviewVersions: ["v1"]
    timeoutSeconds: 5
    clientConfig:
      service:
        name: ${SERVICE_NAME}
        namespace: ${NAMESPACE}
        path: "/validate-pods"
      caBundle: ${CA_BUNDLE} # Replace it with the ca.pem file which is used to generate the certificates and keys for the webhook
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
  ```
## Deployment

### Using Manifests
#### Prerequisites
To deploy this application, you will need the following:
- A Kubernetes cluster with version `1.20` or higher
- `kubectl` installed and configured to communicate with your cluster

#### Steps
1. Clone the kube-bouncer repo and `cd` into the `./manifests` directory
2. Run the following commands to install the 01_deployment.yaml, 02_service.yaml
  ```bash
  kubectl create -f <filename.yaml>
  ```
3. Now, configure the `03_webhook.yaml` manifest based on the requirement and run the above command


**Note:** *More checks and bouncers will be added in the future, and contributions are welcome. If you have any issues or suggestions, please feel free to open an [issue](https://github.com/Kavinraja-G/kube-bouncer/issues).*
