# Provider crossplane-provider-castai

`crossplane-provider-castai` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
CastAI API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/crossplane-provider-castai):
```
up ctp provider install crossplane-contrib/crossplane-provider-castai:v0.39.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-castai
spec:
  package: xpkg.upbound.io/crossplane-contrib/crossplane-provider-castai:v0.39.0
EOF
```

To run the provider with debug logging, create a `DeploymentRuntimeConfig` and
reference it from the `Provider` via `spec.runtimeConfigRef`:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: debug-config
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --debug
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-castai
spec:
  package: xpkg.upbound.io/crossplane-contrib/crossplane-provider-castai:v0.39.0
  runtimeConfigRef:
    name: debug-config
EOF
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/crossplane-provider-castai).

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

### Updating to a New Terraform Provider Version

To regenerate the Crossplane provider against a new version of the
[Terraform CAST AI provider](https://github.com/castai/terraform-provider-castai):

1. **Update the version in the `Makefile`** — change the two version references:

   ```makefile
   export TERRAFORM_PROVIDER_VERSION ?= <NEW_VERSION>        # e.g. 8.39.2
   export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-castai_v<NEW_VERSION>
   ```

2. **Initialize submodules** (first time only):

   ```console
   make submodules
   ```

3. **Generate the provider schema, pull docs, and transform the schema**:

   ```console
   make generate.init
   ```

   This downloads the Terraform provider binary, runs `terraform providers schema`
   to produce `config/schema.json`, clones the upstream docs, and runs
   `hack/transform-schema.py` to convert `nested_type` attributes for Upjet v1
   compatibility.

4. **Run the Upjet code-generation pipeline**:

   ```console
   go run cmd/generator/main.go "$PWD"
   ```

5. **Regenerate deepcopy helpers, CRDs, and other artifacts**:

   ```console
   make generate
   ```

6. **If the new version introduces new Terraform resources**, you also need to:
   - Add an external name entry in `config/external_name.go`.
   - Create a resource config package under `config/<resource>/config.go`.
   - Register the `Configure` function in `config/provider.go`.

7. **Build and verify**:

   ```console
   go build ./...
   ```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/crossplane-provider-castai/issues).
