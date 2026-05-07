//go:build generate
// +build generate

/*
Copyright 2021 Upbound Inc.
*/

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files
//go:generate bash -c "find ./cluster -iname 'zz_*' ! -iname 'zz_generated.managed*.go' -delete"
//go:generate bash -c "find ./namespaced -iname 'zz_*' ! -iname 'zz_generated.managed*.go' -delete"
//go:generate bash -c "find ./cluster -type d -empty -delete"
//go:generate bash -c "find ./namespaced -type d -empty -delete"
//go:generate bash -c "find ../internal/controller/cluster -iname 'zz_*' -delete 2>/dev/null || true"
//go:generate bash -c "find ../internal/controller/namespaced -iname 'zz_*' -delete 2>/dev/null || true"
//go:generate bash -c "find ../internal/controller/cluster -type d -empty -delete 2>/dev/null || true"
//go:generate bash -c "find ../internal/controller/namespaced -type d -empty -delete 2>/dev/null || true"
//go:generate rm -rf ../examples-generated

// Generate documentation from Terraform docs.
//go:generate go run github.com/crossplane/upjet/v2/cmd/scraper -n ${TERRAFORM_PROVIDER_SOURCE} -r ../.work/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_DOCS_PATH} -o ../config/provider-metadata.yaml

// Run Upjet generator
//go:generate go run ../cmd/generator/main.go ..

// Generate deepcopy methodsets and CRD manifests for cluster scoped types
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./cluster/... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate deepcopy methodsets and CRD manifests for namespaced scoped types
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./namespaced/... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc) for cluster scoped types
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./cluster/...

// Generate crossplane-runtime methodsets (resource.Claim, etc) for namespaced scoped types
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./namespaced/...

package apis

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck
)
