/*
Copyright 2022 Upbound Inc.
*/

package awsuserarn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tjcontroller "github.com/upbound/upjet/pkg/controller"
	ujresource "github.com/upbound/upjet/pkg/resource"

	v1alpha1 "github.com/castai/crossplane-provider-castai/apis/castai/v1alpha1"
	"github.com/castai/crossplane-provider-castai/internal/clients"
)

const (
	errNotAWSUserARN = "managed resource is not a AWSUserARN custom resource"
	errUpdate        = "cannot update AWSUserARN"
	errStatusUpdate  = "cannot update status of AWSUserARN"
)

// Setup adds a controller that reconciles AWSUserARN.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.AWSUserARNGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK))
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.AWSUserARN{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.AWSUserARNGroupVersionKind),
			managed.WithExternalConnecter(&connector{
				kube: mgr.GetClient(),
			}),
			managed.WithPollInterval(time.Minute*1),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
			managed.WithConnectionPublishers(cps...),
			managed.WithInitializers()))
}

type connector struct {
	kube client.Client
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	customClient, err := clients.GetConfig(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}

	return &external{
			kube:   c.kube,
			client: *customClient,
		},
		nil
}

type external struct {
	kube   client.Client
	client clients.CustomClient
}

func (e *external) Observe(_ context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	cr, ok := mg.(*v1alpha1.AWSUserARN)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotAWSUserARN)
	}
	if meta.WasDeleted(cr) {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	if cr.Status.AtProvider.ARN == nil {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	if cr.Status.AtProvider.ClusterID != cr.Spec.ForProvider.ClusterID {
		return managed.ExternalObservation{
			ResourceUpToDate: false,
		}, nil
	}

	cr.Status.SetConditions(xpv1.Available())
	ujresource.SetUpToDateCondition(mg, true)

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	err := e.createOrUpdate(ctx, mg)
	return managed.ExternalCreation{}, err
}

func (e *external) createOrUpdate(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.AWSUserARN)
	if !ok {
		return errors.New(errNotAWSUserARN)
	}
	clusterID := cr.Spec.ForProvider.ClusterID
	if clusterID != nil {
		url := fmt.Sprintf("/v1/kubernetes/external-clusters/%s/assume-role-user", *clusterID)
		resp, err := e.client.Get(ctx, url)
		if err != nil {
			return err
		}

		defer close(resp.Body)

		bodyBytes, _ := io.ReadAll(resp.Body)

		var result *v1alpha1.AWSUserARNObservation
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return err
		}

		if result.ARN == nil {
			err := errors.New(fmt.Sprintf("error in api response, %s", bodyBytes))
			return err
		}

		re := regexp.MustCompile(`arn:aws:iam::(\d+):user/(.*)-([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})`)

		matches := re.FindAllStringSubmatch(*result.ARN, -1)
		if len(matches) == 1 && len(matches[0]) == 4 {
			result.ManagementAccountID = &matches[0][1]
			result.UserPrefix = &matches[0][2]
			result.ClusterID = &matches[0][3]
		}
		meta.SetExternalName(cr, *result.ARN)
		cr.Status.AtProvider = *result
		cr.Status.SetConditions(xpv1.Available())

		if err := e.kube.Status().Update(ctx, cr); err != nil {
			return errors.Wrap(err, errStatusUpdate)
		}
		if err := e.kube.Update(ctx, cr); err != nil {
			return errors.Wrap(err, errUpdate)
		}
		return nil
	}
	err := errors.New("no clusterID given")
	return err

}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	err := e.createOrUpdate(ctx, mg)
	return managed.ExternalUpdate{}, err
}

func (e *external) Delete(_ context.Context, _ resource.Managed) error {
	return nil
}

func close(body io.Closer) {
	err := body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
