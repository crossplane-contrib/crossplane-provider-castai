# syntax=docker/dockerfile:1

ARG GO_VERSION=1.26.10

# Rebuild the Terraform CLI and the CAST AI provider plugin from their upstream
# release tags with a Go toolchain that carries the eMudhra CA distrust
# (CVE-2026-111943). Same versions and build flags as upstream.
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS tf-builder

RUN apk --no-cache add git

ARG TARGETOS
ARG TARGETARCH
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY

ENV CGO_ENABLED=0
ENV GOFLAGS="-trimpath -buildvcs=false"
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}

# Terraform CLI (MPL-2.0)
RUN git clone -c advice.detachedHead=false --depth 1 --branch "v${TERRAFORM_VERSION}" \
      https://github.com/hashicorp/terraform /src/terraform \
  && cd /src/terraform \
  && go build -ldflags "-s -w -X github.com/hashicorp/terraform/version.dev=no" -o /out/terraform . \
  && install -D -m 0644 LICENSE /out/licenses/terraform/LICENSE

# CAST AI Terraform provider plugin (MIT)
RUN git clone -c advice.detachedHead=false --depth 1 --branch "v${TERRAFORM_PROVIDER_VERSION}" \
      "https://github.com/castai/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}" /src/provider \
  && cd /src/provider \
  && go build -ldflags "-s -w -X main.version=${TERRAFORM_PROVIDER_VERSION} -X main.commit=$(git rev-parse HEAD)" \
       -o "/out/${TERRAFORM_NATIVE_PROVIDER_BINARY}" . \
  && install -D -m 0644 LICENSE "/out/licenses/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}/LICENSE"

FROM alpine:3.22.7
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH
ARG CROSSPLANE_PROVIDER_VERSION

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532
ENV CASTAI_ADDITIONAL_USER_AGENT crossplane-provider-castai/${CROSSPLANE_PROVIDER_VERSION}

# Setup Terraform environment

## Provider-dependent configuration
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

## End of - Provider-dependent configuration

ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0

COPY --from=tf-builder /out/terraform /usr/local/bin/terraform
COPY --from=tf-builder "/out/${TERRAFORM_NATIVE_PROVIDER_BINARY}" "${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}"
COPY --from=tf-builder /out/licenses /usr/local/share/licenses
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

RUN chmod +x /usr/local/bin/terraform ${PLUGIN_DIR}/* \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION ${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE ${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION ${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
