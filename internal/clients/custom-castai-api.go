/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/dkb-bank/provider-castai/apis/v1beta1"
)

const (
	errNoProviderConfigRef            = "no providerConfigRef is given"
	errCannotGetProvider              = "cannot get referenced Provider"
	errCannotTrackProviderConfigUsage = "cannot track ProviderConfig usage"
	errOnlySecretSourceAllowed        = "only Secret supported as Source"
	errExtractSecret                  = "cannot extract credentials from secret"
	errExtractSecretKey               = "cannot extract secret key"
	errGetCredentialsSecret           = "cannot get credentials secret"
	errInvalidSecretData              = "'%s' is required in secret data"
)

// GetConfig loads config for custom castai api
func GetConfig(ctx context.Context, c client.Client, mg resource.Managed) (*CustomClient, error) {
	switch {
	case mg.GetProviderConfigReference() != nil:
		return useProviderConfig(ctx, c, mg)
	default:
		return nil, errors.New(errNoProviderConfigRef)
	}
}

// CustomClient is the struct used to
type CustomClient struct {
	client  http.Client
	baseURL string
	headers map[string]string
}

// Get calls a get endpoint with default host and headers
func (c *CustomClient) Get(ctx context.Context, url string) (*http.Response, error) {
	fullURL := c.baseURL + url
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	return c.client.Do(req)

}

func useProviderConfig(ctx context.Context, c client.Client, mg resource.Managed) (*CustomClient, error) { // nolint:gocyclo
	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, errCannotGetProvider)
	}

	t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errCannotTrackProviderConfigUsage)
	}

	if pc.Spec.Credentials.Source != xpv1.CredentialsSourceSecret {
		return nil, errors.New(errOnlySecretSourceAllowed)
	}

	credentials, credsErr := extractCredentialsFromSecret(ctx, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if credsErr != nil {
		return nil, errors.Wrap(credsErr, errExtractSecret)
	}

	customClient := CustomClient{
		baseURL: credentials.APIUrl,
		headers: map[string]string{
			"X-API-Key": credentials.APIToken,
		},
	}

	return &customClient, nil
}

type providerCredentials struct {
	APIToken string `json:"api_token"`
	APIUrl   string `json:"api_url"`
}

func extractCredentialsFromSecret(ctx context.Context, client client.Client, s xpv1.CommonCredentialSelectors) (*providerCredentials, error) {
	if s.SecretRef == nil {
		return nil, errors.New(errExtractSecretKey)
	}
	secret := &corev1.Secret{}
	if err := client.Get(ctx, types.NamespacedName{Namespace: s.SecretRef.Namespace, Name: s.SecretRef.Name}, secret); err != nil {
		return nil, errors.Wrap(err, errGetCredentialsSecret)
	}

	credentialJSON := secret.Data[s.SecretRef.Key]

	if credentialJSON == nil {
		return nil, errors.New(fmt.Sprintf(errInvalidSecretData, s.SecretRef.Key))
	}

	var credentials *providerCredentials
	err := json.Unmarshal(credentialJSON, &credentials)

	if err != nil {
		return nil, err
	}

	return credentials, nil
}
