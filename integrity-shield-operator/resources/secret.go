//
// Copyright 2020 IBM Corporation
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
	_ "embed"

	apiv1alpha1 "github.com/IBM/integrity-enforcer/integrity-shield-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed fulcio.pem
var SigStoreDefaultRootPem string

//regkey.yaml
func BuildRegKeySecretForIShield(cr *apiv1alpha1.IntegrityShield) *corev1.Secret {
	sec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.GetRegKeySecretName(),
			Namespace: cr.Namespace,
		},
		Data: map[string][]byte{
			corev1.DockerConfigJsonKey: cr.Spec.RegKeySecret.Value,
		},
		Type: corev1.SecretTypeDockerConfigJson,
	}
	return sec
}

// ishield-server-tls
func BuildTlsSecretForIShield(cr *apiv1alpha1.IntegrityShield) *corev1.Secret {
	var empty []byte
	sec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.GetWebhookServerTlsSecretName(),
			Namespace: cr.Namespace,
		},
		Data: map[string][]byte{
			corev1.TLSCertKey:       empty, // "tls.crt"
			corev1.TLSPrivateKeyKey: empty,
			"ca.crt":                empty,
		},
		Type: corev1.SecretTypeTLS,
	}
	return sec
}

// ishield-sigstore-root-cert
func BuildSigStoreDefaultRootSecretForIShield(cr *apiv1alpha1.IntegrityShield) *corev1.Secret {
	sec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.GetSigStoreDefaultRootSecretName(),
			Namespace: cr.Namespace,
		},
		Data: map[string][]byte{
			apiv1alpha1.DefaultCertFilename: []byte(SigStoreDefaultRootPem),
		},
	}
	return sec
}
