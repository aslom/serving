/*
Copyright 2019 The Knative Authors

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
package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knative/pkg/apis/duck"
	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func TestCertificateDuckTypes(t *testing.T) {
	tests := []struct {
		name string
		t    duck.Implementable
	}{{
		name: "conditions",
		t:    &duckv1alpha1.Conditions{},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := duck.VerifyType(&Certificate{}, test.t)
			if err != nil {
				t.Errorf("VerifyType(Certificate, %T) = %v", test.t, err)
			}
		})
	}
}

func TestCertificateGetGroupVersionKind(t *testing.T) {
	c := Certificate{}
	expected := SchemeGroupVersion.WithKind("Certificate")
	if diff := cmp.Diff(expected, c.GetGroupVersionKind()); diff != "" {
		t.Errorf("Unexpected diff (-want, +got) = %s", diff)
	}
}

func TestMarkReady(t *testing.T) {
	c := Certificate{}
	c.Status.InitializeConditions()
	checkCondition(c.Status, CertificateCondidtionReady, corev1.ConditionUnknown, t)

	c.Status.MarkReady()
	checkCondition(c.Status, CertificateCondidtionReady, corev1.ConditionTrue, t)
}

func checkCondition(cs CertificateStatus, ct duckv1alpha1.ConditionType, status corev1.ConditionStatus, t *testing.T) *duckv1alpha1.Condition {
	cond := cs.GetCondition(ct)
	if cond == nil {
		t.Fatalf("Get(%v) = nil, wanted %v=%v", ct, ct, status)
	}
	if cond.Status != status {
		t.Fatalf("Get(%v) = %v, wanted %v", ct, cond.Status, status)
	}
	return cond
}
