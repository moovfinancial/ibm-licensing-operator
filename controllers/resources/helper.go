//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"time"

	rhmp "github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/v1beta1"

	networkingv1 "k8s.io/api/networking/v1"

	odlm "github.com/IBM/operand-deployment-lifecycle-manager/api/v1alpha1"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/go-logr/logr"
	servicecav1 "github.com/openshift/api/operator/v1"
	routev1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	c "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/IBM/ibm-licensing-operator/api/v1alpha1"
)

// cannot set to const due to k8s struct needing pointers to primitive types

var TrueVar = true
var FalseVar = false

var DefaultSecretMode int32 = 420
var Seconds60 int64 = 60

var IsRouteAPI = true
var IsServiceCAAPI = true
var RHMPEnabled = false
var IsUIEnabled = false
var IsODLM = true
var UIPlatformSecretName = "platform-oidc-credentials"

var PathType = networkingv1.PathTypeImplementationSpecific

// Important product values needed for annotations
const LicensingProductName = "IBM Cloud Platform Common Services"
const LicensingProductID = "068a62892a1e4db39641342e592daa25"
const LicensingProductMetric = "FREE"

const randStringCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ocpCertSecretNameTag = "service.beta.openshift.io/serving-cert-secret-name" // #nosec
const OcpCheckString = "ocp-check-secret"
const OcpPrometheusCheckString = "ocp-prometheus-check-secret"

var randStringCharsetLength = big.NewInt(int64(len(randStringCharset)))

var annotationsForServicesToCheck = [...]string{ocpCertSecretNameTag}

type ResourceObject interface {
	metav1.Object
	runtime.Object
}

func RandString(length int) (string, error) {
	reader := rand.Reader
	outputStringByte := make([]byte, length)
	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(reader, randStringCharsetLength)
		if err != nil {
			return "", err
		}
		outputStringByte[i] = randStringCharset[charIndex.Int64()]
	}
	return string(outputStringByte), nil
}

func Contains(s []corev1.LocalObjectReference, e corev1.LocalObjectReference) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// we could use reflection to have this method for all types but for now strings would be enough
func ListsEqualsLikeSets(list1 []string, list2 []string) bool {
	if list1 == nil {
		return list2 == nil
	}
	if len(list1) != len(list2) {
		return false
	}
	for _, item1 := range list1 {
		found := false
		for _, item2 := range list2 {
			if item2 == item1 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func AnnotationsForPod() map[string]string {
	return map[string]string{"productName": LicensingProductName,
		"productID": LicensingProductID, "productMetric": LicensingProductMetric}
}

func GetSecretToken(name string, namespace string, secretKey string, metaLabels map[string]string) (*corev1.Secret, error) {
	randString, err := RandString(24)
	if err != nil {
		return nil, err
	}
	expectedSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    metaLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{secretKey: randString},
	}
	return expectedSecret, nil
}

func AnnotateForService(httpCertSource v1alpha1.HTTPSCertsSource, isHTTPS bool, certName string) map[string]string {
	if IsServiceCAAPI && isHTTPS && httpCertSource == v1alpha1.OcpCertsSource {
		return map[string]string{ocpCertSecretNameTag: certName}
	}
	return map[string]string{}
}

func UpdateResource(reqLogger *logr.Logger, client c.Client,
	expectedResource ResourceObject, foundResource ResourceObject) (reconcile.Result, error) {
	resTypeString := reflect.TypeOf(expectedResource).String()
	(*reqLogger).Info("Updating " + resTypeString)
	expectedResource.SetResourceVersion(foundResource.GetResourceVersion())
	err := client.Update(context.TODO(), expectedResource)
	if err != nil {
		// only need to delete resource as new will be recreated on next reconciliation
		(*reqLogger).Info("Could not update "+resTypeString+", due to having not compatible changes between expected and updated resource, "+
			"will try to delete it and create new one...", "Namespace", foundResource.GetNamespace(), "Name", foundResource.GetName())
		return DeleteResource(reqLogger, client, foundResource)
	}
	(*reqLogger).Info("Updated "+resTypeString+" successfully", "Namespace", expectedResource.GetNamespace(), "Name", expectedResource.GetName())
	// Resource updated - return and do not requeue as it might not consider extra values
	return reconcile.Result{}, nil
}

func UpdateServiceIfNeeded(reqLogger *logr.Logger, client c.Client, expectedService *corev1.Service, foundService *corev1.Service) (reconcile.Result, error) {
	for _, annotation := range annotationsForServicesToCheck {
		if foundService.Annotations[annotation] != expectedService.Annotations[annotation] {
			expectedService.Spec.ClusterIP = foundService.Spec.ClusterIP
			return UpdateResource(reqLogger, client, expectedService, foundService)
		}
	}
	return reconcile.Result{}, nil
}

func UpdateServiceMonitor(reqLogger *logr.Logger, client c.Client, expected, found *monitoringv1.ServiceMonitor) (reconcile.Result, error) {
	if expected == nil || found == nil {
		err := errors.New("cannot update to empty service monitor")
		(*reqLogger).Error(err, "Expected and found service monitor cannot be nil.")
		return reconcile.Result{}, err
	}
	updateResource := func() (reconcile.Result, error) {
		return UpdateResource(reqLogger, client, expected, found)
	}
	for _, annotation := range annotationsForServicesToCheck {
		if found.Annotations[annotation] != expected.Annotations[annotation] {
			return updateResource()
		}
	}
	expectedSpec := expected.Spec
	foundSpec := found.Spec
	// we assume only one endpoint, if changed in expected service monitor then modify this method as well
	if len(expectedSpec.Endpoints) != 1 {
		err := errors.New("expected service monitor endpoints error")
		(*reqLogger).Error(
			err, "Expected service monitor should have 1 endpoint, change this function otherwise.",
			"Namespace", found.GetNamespace(), "Name", found.GetName())
		return reconcile.Result{}, err
	}
	if len(foundSpec.Endpoints) != len(expectedSpec.Endpoints) {
		// deleting will also cause updating, this needs to be done when in-place update cannot work
		return updateResource()
	}
	expectedEndpoint := expectedSpec.Endpoints[0]
	foundEndpoint := foundSpec.Endpoints[0]
	if expectedEndpoint.Scheme != foundEndpoint.Scheme {
		return updateResource()
	}
	if expectedEndpoint.TargetPort != nil {
		if foundEndpoint.TargetPort == nil ||
			expectedEndpoint.TargetPort.StrVal != foundEndpoint.TargetPort.StrVal ||
			expectedEndpoint.TargetPort.IntVal != foundEndpoint.TargetPort.IntVal {
			return updateResource()
		}
	} else {
		if foundEndpoint.TargetPort != nil {
			return updateResource()
		}
	}
	if expectedEndpoint.Interval != foundEndpoint.Interval || expectedEndpoint.Path != foundEndpoint.Path {
		return updateResource()
	}
	if expectedEndpoint.RelabelConfigs != nil {
		if foundEndpoint.RelabelConfigs == nil ||
			len(expectedEndpoint.RelabelConfigs) != len(foundEndpoint.RelabelConfigs) {
			return updateResource()
		}
		// we assume only one relabeling, if changed in expected service monitor then modify this method as well
		if len(expectedEndpoint.RelabelConfigs) != 1 {
			err := errors.New("expected service monitor relabeling error")
			(*reqLogger).Error(
				err, "Expected service monitor should have 1 relabeling, change this function otherwise.",
				"Namespace", found.GetNamespace(), "Name", found.GetName())
			return reconcile.Result{}, err
		}
		expectedRelabeling := expectedEndpoint.RelabelConfigs[0]
		foundRelabeling := foundEndpoint.RelabelConfigs[0]
		if expectedRelabeling.Replacement != foundRelabeling.Replacement ||
			expectedRelabeling.TargetLabel != foundRelabeling.TargetLabel {
			return updateResource()
		}
	} else {
		if foundEndpoint.RelabelConfigs != nil {
			return updateResource()
		}
	}
	if expectedEndpoint.TLSConfig != nil {
		if foundEndpoint.TLSConfig == nil ||
			!reflect.DeepEqual(expectedEndpoint.TLSConfig, foundEndpoint.TLSConfig) {
			return updateResource()
		}
	} else {
		if foundEndpoint.TLSConfig != nil {
			return updateResource()
		}
	}
	if !reflect.DeepEqual(expectedSpec.Selector, foundSpec.Selector) {
		return updateResource()
	}
	return reconcile.Result{}, nil
}

func DeleteResource(reqLogger *logr.Logger, client c.Client, foundResource ResourceObject) (reconcile.Result, error) {
	resTypeString := reflect.TypeOf(foundResource).String()
	err := client.Delete(context.TODO(), foundResource)
	if err != nil {
		if apierrors.IsNotFound(err) {
			(*reqLogger).Info("Could not delete "+resTypeString+", as it was already deleted", "Namespace", foundResource.GetNamespace(), "Name", foundResource.GetName())
		} else {
			(*reqLogger).Error(err, "Failed to delete "+resTypeString+" during recreation", "Namespace", foundResource.GetNamespace(), "Name", foundResource.GetName())
			return reconcile.Result{}, err
		}
	} else {
		// Resource deleted successfully - return and requeue to create new one
		(*reqLogger).Info("Deleted "+resTypeString+" successfully", "Namespace", foundResource.GetNamespace(), "Name", foundResource.GetName())
	}
	return reconcile.Result{Requeue: true, RequeueAfter: time.Second * 30}, nil
}

func UpdateOwner(reqLogger *logr.Logger, client c.Client, owner ResourceObject) (reconcile.Result, error) {
	resTypeString := reflect.TypeOf(owner).String()
	err := client.Get(context.TODO(), types.NamespacedName{Name: owner.GetName(), Namespace: owner.GetNamespace()}, owner)
	if err != nil {
		(*reqLogger).Error(err, "Failed to update owner data "+resTypeString+"", "Namespace", owner.GetNamespace(), "Name", owner.GetName())
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}

func GetOCPSecretCheckScript() string {
	script := `while true; do
  echo "$(date): Checking for ocp secret"
  ls /opt/licensing/certs/* && break
  echo "$(date): Required ocp secret not found ... try again in 30s"
  sleep 30
done
echo "$(date): All required secrets exist"
`
	return script
}

func GetOCPPrometheusSecretCheckScript() string {
	script := `while true; do
  echo "$(date): Checking for ocp prometheus secret"
  ls /opt/prometheus/certs/* && break
  echo "$(date): Required ocp prometheus secret not found ... try again in 30s"
  sleep 30
done
echo "$(date): All required secrets exist"
`
	return script
}

func UpdateCacheClusterExtensions(client c.Reader) error {
	var watchNamespaceEnvVar = "WATCH_NAMESPACE"

	namespace, found := os.LookupEnv(watchNamespaceEnvVar)
	if !found {
		return errors.New("WATCH_NAMESPACE not found")
	}

	listOpts := []c.ListOption{
		c.InNamespace(namespace),
	}

	MeterDefinitionCRD := &rhmp.MeterDefinition{}
	if err := client.List(context.TODO(), MeterDefinitionCRD, listOpts...); err == nil {
		RHMPEnabled = true
	} else {
		RHMPEnabled = false
	}

	routeTestInstance := &routev1.Route{}
	if err := client.List(context.TODO(), routeTestInstance, listOpts...); err == nil {
		IsRouteAPI = true
	} else {
		IsRouteAPI = false
	}

	serviceCAInstance := &servicecav1.ServiceCA{}
	if err := client.List(context.TODO(), serviceCAInstance, listOpts...); err == nil {
		IsServiceCAAPI = true
	} else {
		IsServiceCAAPI = false
	}

	odlmTestInstance := &odlm.OperandBindInfo{}
	if err := client.List(context.TODO(), odlmTestInstance, listOpts...); err == nil {
		IsODLM = true
	} else {
		IsODLM = false
	}

	return nil
}

// Returns true if configmaps are equal
func CompareConfigMap(cm1, cm2 *corev1.ConfigMap) bool {
	return reflect.DeepEqual(cm1.Data, cm2.Data) && reflect.DeepEqual(cm1.Labels, cm2.Labels)
}

// Returns true if routes are equal
func CompareRoutes(reqLogger logr.Logger, expectedRoute, foundRoute *routev1.Route) bool {
	areEqual := false
	if foundRoute.ObjectMeta.Name != expectedRoute.ObjectMeta.Name {
		reqLogger.Info("Names not equal", "old", foundRoute.ObjectMeta.Name, "new", expectedRoute.ObjectMeta.Name)
	} else if foundRoute.Spec.To.Name != expectedRoute.Spec.To.Name {
		reqLogger.Info("Specs To Name not equal",
			"old", fmt.Sprintf("%v", foundRoute.Spec),
			"new", fmt.Sprintf("%v", expectedRoute.Spec))
	} else if foundRoute.Spec.TLS == nil && expectedRoute.Spec.TLS != nil {
		reqLogger.Info("Found Route has empty TLS options, but Expected Route has not empty TLS options",
			"old", fmt.Sprintf("%v", foundRoute.Spec.TLS),
			"new", fmt.Sprintf("%v", GetTLSDataAsString(expectedRoute)))
	} else if foundRoute.Spec.TLS != nil && expectedRoute.Spec.TLS == nil {
		reqLogger.Info("Expected Route has empty TLS options, but Found Route has not empty TLS options",
			"old", fmt.Sprintf("%v", GetTLSDataAsString(foundRoute)),
			"new", fmt.Sprintf("%v", expectedRoute.Spec.TLS))
	} else if foundRoute.Spec.TLS != nil && expectedRoute.Spec.TLS != nil {
		if foundRoute.Spec.TLS.Termination != expectedRoute.Spec.TLS.Termination {
			reqLogger.Info("Expected Route has different TLS Termination option than Found Route",
				"old", fmt.Sprintf("%v", foundRoute.Spec.TLS.Termination),
				"new", fmt.Sprintf("%v", expectedRoute.Spec.TLS.Termination))
		}
		if foundRoute.Spec.TLS.InsecureEdgeTerminationPolicy != expectedRoute.Spec.TLS.InsecureEdgeTerminationPolicy {
			reqLogger.Info("Expected Route has different TLS InsecureEdgeTerminationPolicy option than Found Route",
				"old", fmt.Sprintf("%v", foundRoute.Spec.TLS.InsecureEdgeTerminationPolicy),
				"new", fmt.Sprintf("%v", expectedRoute.Spec.TLS.InsecureEdgeTerminationPolicy))
		}
	} else {
		areEqual = true
	}
	return areEqual
}

func GetTLSDataAsString(route *routev1.Route) string {
	return fmt.Sprintf("{Termination: %v, InsecureEdgeTerminationPolicy: %v, Certificate: %s, CACertificate: %s, DestinationCACertificate: %s}",
		route.Spec.TLS.Termination, route.Spec.TLS.InsecureEdgeTerminationPolicy,
		route.Spec.TLS.Certificate, route.Spec.TLS.CACertificate, route.Spec.TLS.DestinationCACertificate)
}
