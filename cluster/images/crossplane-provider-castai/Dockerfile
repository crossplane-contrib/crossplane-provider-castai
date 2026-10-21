FROM alpine:3.20.10
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH
ARG CROSSPLANE_PROVIDER_VERSION

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532
ENV CASTAI_ADDITIONAL_USER_AGENT crossplane-provider-castai/${CROSSPLANE_PROVIDER_VERSION}

# Setup Terraform environment

## Provider-dependent configuration
ARG OPENTOFU_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

## End of - Provider-dependent configuration

ENV PLUGIN_DIR /terraform/provider-mirror/registry.opentofu.org/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://github.com/opentofu/opentofu/releases/download/v${OPENTOFU_VERSION}/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD https://github.com/opentofu/opentofu/releases/download/v${OPENTOFU_VERSION}/tofu_${OPENTOFU_VERSION}_SHA256SUMS /tmp/tofu_SHA256SUMS
ADD https://github.com/castai/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}/releases/download/v${TERRAFORM_PROVIDER_VERSION}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

# Symlink created because Upjet runs the CLI as `terraform`.
RUN cd /tmp && grep " tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip$" tofu_SHA256SUMS | sha256sum -c - \
  && unzip /tmp/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip tofu -d /usr/local/bin \
  && chmod +x /usr/local/bin/tofu \
  && ln -s tofu /usr/local/bin/terraform \
  && rm /tmp/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/tofu_SHA256SUMS \
  && unzip /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION ${OPENTOFU_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE ${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION ${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
