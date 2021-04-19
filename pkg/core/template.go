/*
Copyright 2017 The Kubernetes Authors.

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

package core

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	compositionv1alpha1 "composition-controller/pkg/apis/compositioncontroller/v1alpha1"
)

func deploymentName(compositionName string) string {
	return compositionName + "-deployment"
}

func serviceName(compositionName string) string {
	return compositionName + "-service"
}

func ingressName(compositionName string) string {
	return compositionName + "-ingress"
}

func hpaName(compositionName string) string {
	return compositionName + "-hpa"
}

func pdbName(compositionName string) string {
	return compositionName + "-pdb"
}

func createLabels(compositionName string) map[string]string {
	return map[string]string{
		"app":        compositionName + "-app",
		"controller": compositionName,
	}
}

// newDeployment creates a new Deployment for a Composition resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Composition resource that 'owns' it.
func newDeployment(composition *compositionv1alpha1.Composition) *appsv1.Deployment {
	labels := createLabels(composition.Name)
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName(composition.Name),
			Namespace: composition.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(composition, compositionv1alpha1.SchemeGroupVersion.WithKind("Composition")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			// TODO: since we have hpa, we can remove this
			Replicas: composition.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: map[string]string{"sidecar.istio.io/rewriteAppHTTPProbers": "true"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
}

func newService(composition *compositionv1alpha1.Composition) *corev1.Service {
	return nil
}

func newIngress(composition *compositionv1alpha1.Composition) *networkingv1beta1.Ingress {
	return nil
}

func newHPA(composition *compositionv1alpha1.Composition) *autoscalingv2beta2.HorizontalPodAutoscaler {
	return nil
}

func newPDB(composition *compositionv1alpha1.Composition) *policyv1beta1.PodDisruptionBudget {
	return nil
}
