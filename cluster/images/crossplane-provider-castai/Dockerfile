# syntax=docker/dockerfile:1.6
#
# CVE test image:
# - Newer Alpine for OpenSSL/base CVEs
# - Optional OpenTofu CLI (MPL) instead of HashiCorp Terraform 1.5.7 (BSL-safe)
#   Build with: CLI_FLAVOR=opentofu
# - Default CLI_FLAVOR=terraform keeps current Terraform 1.5.7 behavior

ARG CLI_FLAVOR=terraform
ARG OPENTOFU_VERSION=1.9.4

# ---------------------------------------------------------------------------
# Optional OpenTofu source stage (only used when CLI_FLAVOR=opentofu)
# ---------------------------------------------------------------------------
FROM ghcr.io/opentofu/opentofu:${OPENTOFU_VERSION}-minimal AS opentofu

# ---------------------------------------------------------------------------
# Runtime image
# ---------------------------------------------------------------------------
FROM alpine:3.22.1

RUN apk --no-cache upgrade \
  && apk --no-cache add ca-certificates bash openssl

ARG TARGETOS
ARG TARGETARCH
ARG CROSSPLANE_PROVIDER_VERSION
ARG CLI_FLAVOR=terraform
ARG OPENTOFU_VERSION=1.9.4

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532
ENV CASTAI_ADDITIONAL_USER_AGENT crossplane-provider-castai/${CROSSPLANE_PROVIDER_VERSION}

# Setup Terraform/OpenTofu environment
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

ENV PLUGIN_DIR=/terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0
ENV CLI_FLAVOR=${CLI_FLAVOR}

RUN mkdir -p ${PLUGIN_DIR}

# CAST AI Terraform provider plugin (shared by both CLI flavors)
ADD https://github.com/castai/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}/releases/download/v${TERRAFORM_PROVIDER_VERSION}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/provider.zip
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

# Install CLI:
# - terraform: HashiCorp Terraform ${TERRAFORM_VERSION} (default 1.5.7)
# - opentofu: OpenTofu binary installed as `terraform` so Upjet keeps working
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/terraform.zip
COPY --from=opentofu /usr/local/bin/tofu /tmp/tofu

RUN set -eux; \
  unzip /tmp/provider.zip -d "${PLUGIN_DIR}"; \
  chmod +x "${PLUGIN_DIR}"/*; \
  rm /tmp/provider.zip; \
  if [ "${CLI_FLAVOR}" = "opentofu" ]; then \
    install -m 0755 /tmp/tofu /usr/local/bin/terraform; \
    ln -sf /usr/local/bin/terraform /usr/local/bin/tofu; \
    echo "Installed OpenTofu ${OPENTOFU_VERSION} as /usr/local/bin/terraform"; \
  else \
    unzip /tmp/terraform.zip -d /usr/local/bin; \
    chmod +x /usr/local/bin/terraform; \
    echo "Installed HashiCorp Terraform ${TERRAFORM_VERSION}"; \
  fi; \
  rm -f /tmp/terraform.zip /tmp/tofu; \
  terraform version; \
  chown -R ${USER_ID}:${USER_ID} /terraform

# Provider controller needs these environment variables at runtime
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
