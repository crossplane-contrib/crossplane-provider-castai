# Provider CastAI

`provider-castai` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
CastAI API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/haarchri/provider-castai):
```
up ctp provider install haarchri/provider-castai:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-castai
spec:
  package: haarchri/provider-castai:v0.1.0
EOF
```

# phase1:
```
apiVersion: castai.upbound.io/v1alpha1
kind: EksCluster
metadata:
  name: sample-cluster
  labels:
    cast-ai-cluster: sample
spec:
  forProvider:
    accountId: "123456789101"
    name: sample-cluster
    region: eu-central-1
```

# phase2:
```
apiVersion: castai.upbound.io/v1alpha1
kind: EksCluster
metadata:
  name: sample-cluster
  labels:
    cast-ai-cluster: sample
spec:
  forProvider:
    accountId: "123456789101"
    assumeRoleArn: arn:aws:iam::123456789101:role/cast-eks-sample-cluster-cluster-role
    deleteNodesOnDisconnect: false
    name: sample-cluster
    region: eu-central-1
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/haarchri/provider-castai).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/haarchri/provider-castai/issues).
