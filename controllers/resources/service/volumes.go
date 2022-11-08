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

package service

import (
	corev1 "k8s.io/api/core/v1"

	operatorv1alpha1 "github.com/IBM/ibm-licensing-operator/api/v1alpha1"
	"github.com/IBM/ibm-licensing-operator/controllers/resources"
)

const APISecretTokenVolumeName = "api-token"
const APIUploadTokenVolumeName = "token-upload"
const MeteringAPICertsVolumeName = "metering-api-certs"
const LicensingHTTPSCertsVolumeName = "licensing-https-certs"
const PrometheusHTTPSCertsVolumeName = "prometheus-https-certs"

func getLicensingVolumeMounts(spec operatorv1alpha1.IBMLicensingSpec) []corev1.VolumeMount {
	var volumeMounts = []corev1.VolumeMount{
		{
			Name:      APISecretTokenVolumeName,
			MountPath: "/opt/ibm/licensing/" + APISecretTokenKeyName,
			SubPath:   APISecretTokenKeyName,
			ReadOnly:  true,
		},
		{
			Name:      APIUploadTokenVolumeName,
			MountPath: "/opt/ibm/licensing/" + APIUploadTokenKeyName,
			SubPath:   APIUploadTokenKeyName,
			ReadOnly:  true,
		},
	}
	if spec.HTTPSEnable {
		if spec.HTTPSCertsSource == operatorv1alpha1.CustomCertsSource || (resources.IsServiceCAAPI && spec.HTTPSCertsSource == operatorv1alpha1.OcpCertsSource) {
			volumeMounts = append(volumeMounts, []corev1.VolumeMount{
				{
					Name:      LicensingHTTPSCertsVolumeName,
					MountPath: "/opt/licensing/certs/",
					ReadOnly:  true,
				},
			}...)
			if spec.IsPrometheusServiceNeeded() {
				volumeMounts = append(volumeMounts, []corev1.VolumeMount{
					{
						Name:      PrometheusHTTPSCertsVolumeName,
						MountPath: "/opt/prometheus/certs",
						ReadOnly:  true,
					},
				}...)
			}
		}

	}
	if spec.IsMetering() {
		volumeMounts = append(volumeMounts, []corev1.VolumeMount{
			{
				Name:      MeteringAPICertsVolumeName,
				MountPath: "/opt/metering/certs/",
				ReadOnly:  true,
			},
		}...)
	}
	return volumeMounts
}

func getLicensingVolumes(spec operatorv1alpha1.IBMLicensingSpec) []corev1.Volume {
	var volumes []corev1.Volume

	apiSecretTokenVolume := corev1.Volume{
		Name: APISecretTokenVolumeName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName:  spec.APISecretToken,
				DefaultMode: &resources.DefaultSecretMode,
			},
		},
	}

	volumes = append(volumes, apiSecretTokenVolume)

	apiUploadTokenVolume := corev1.Volume{
		Name: APIUploadTokenVolumeName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName:  APIUploadTokenName,
				DefaultMode: &resources.DefaultSecretMode,
			},
		},
	}

	volumes = append(volumes, apiUploadTokenVolume)

	if spec.IsMetering() {
		meteringAPICertVolume := corev1.Volume{
			Name: MeteringAPICertsVolumeName,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName:  "icp-metering-api-secret",
					DefaultMode: &resources.DefaultSecretMode,
					Optional:    &resources.TrueVar,
				},
			},
		}

		volumes = append(volumes, meteringAPICertVolume)
	}

	if spec.HTTPSEnable {
		if spec.HTTPSCertsSource == operatorv1alpha1.CustomCertsSource {
			volumes = append(volumes, resources.GetVolume(LicensingHTTPSCertsVolumeName, "ibm-licensing-certs"))
		} else if resources.IsServiceCAAPI && spec.HTTPSCertsSource == operatorv1alpha1.OcpCertsSource {
			volumes = append(volumes, resources.GetVolume(LicensingHTTPSCertsVolumeName, LicenseServiceOCPCertName))
		}
		if spec.IsPrometheusServiceNeeded() {
			volumes = append(volumes, resources.GetVolume(PrometheusHTTPSCertsVolumeName, PrometheusServiceOCPCertName))
		}
	}

	return volumes
}
