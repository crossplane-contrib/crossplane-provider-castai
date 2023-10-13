# Provider CastAI

`crossplane-provider-castai` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
CastAI API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/castai/crossplane-provider-castai):
```
up ctp provider install castai/crossplane-provider-castai:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-castai
spec:
  package: xpkg.upbound.io/castai/crossplane-provider-castai:v0.3.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/castai/crossplane-provider-castai).

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
open an [issue](https://github.com/castai/crossplane-provider-castai/issues).
