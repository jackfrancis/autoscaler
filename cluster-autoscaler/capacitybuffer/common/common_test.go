/*
Copyright 2025 The Kubernetes Authors.

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

package common

import (
	"errors"
	"testing"

	v1 "k8s.io/autoscaler/cluster-autoscaler/apis/capacitybuffer/autoscaling.x-k8s.io/v1beta1"
	"k8s.io/autoscaler/cluster-autoscaler/capacitybuffer"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSetBufferAsReadyForProvisioning(t *testing.T) {
	tests := []struct {
		name               string
		existingConditions []metav1.Condition
		wantConditionTypes []string
	}{
		{
			name:               "adds ReadyForProvisioning when no conditions exist",
			existingConditions: nil,
			wantConditionTypes: []string{capacitybuffer.ReadyForProvisioningCondition},
		},
		{
			name: "preserves existing Provisioning condition",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.ProvisioningCondition, Status: metav1.ConditionTrue},
			},
			wantConditionTypes: []string{capacitybuffer.ProvisioningCondition, capacitybuffer.ReadyForProvisioningCondition},
		},
		{
			name: "preserves existing LimitedByQuotas condition",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.LimitedByQuotasCondition, Status: metav1.ConditionFalse},
			},
			wantConditionTypes: []string{capacitybuffer.LimitedByQuotasCondition, capacitybuffer.ReadyForProvisioningCondition},
		},
		{
			name: "preserves multiple existing conditions",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.ProvisioningCondition, Status: metav1.ConditionTrue},
				{Type: capacitybuffer.LimitedByQuotasCondition, Status: metav1.ConditionFalse},
			},
			wantConditionTypes: []string{capacitybuffer.ProvisioningCondition, capacitybuffer.LimitedByQuotasCondition, capacitybuffer.ReadyForProvisioningCondition},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &v1.CapacityBuffer{
				Status: v1.CapacityBufferStatus{
					Conditions: tt.existingConditions,
				},
			}

			SetBufferAsReadyForProvisioning(buffer, nil, nil, nil, nil)

			if len(buffer.Status.Conditions) != len(tt.wantConditionTypes) {
				t.Fatalf("got %d conditions, want %d", len(buffer.Status.Conditions), len(tt.wantConditionTypes))
			}

			conditionTypesMap := make(map[string]bool)
			for _, cond := range buffer.Status.Conditions {
				conditionTypesMap[cond.Type] = true
			}

			for _, wantType := range tt.wantConditionTypes {
				if !conditionTypesMap[wantType] {
					t.Errorf("missing expected condition type %q", wantType)
				}
			}
		})
	}
}

func TestSetBufferAsNotReadyForProvisioning(t *testing.T) {
	tests := []struct {
		name               string
		existingConditions []metav1.Condition
		wantConditionTypes []string
	}{
		{
			name:               "adds ReadyForProvisioning=False when no conditions exist",
			existingConditions: nil,
			wantConditionTypes: []string{capacitybuffer.ReadyForProvisioningCondition},
		},
		{
			name: "preserves existing Provisioning condition",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.ProvisioningCondition, Status: metav1.ConditionTrue},
			},
			wantConditionTypes: []string{capacitybuffer.ProvisioningCondition, capacitybuffer.ReadyForProvisioningCondition},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &v1.CapacityBuffer{
				Status: v1.CapacityBufferStatus{
					Conditions: tt.existingConditions,
				},
			}

			SetBufferAsNotReadyForProvisioning(buffer, nil, nil, nil, nil, errors.New("test error"))

			if len(buffer.Status.Conditions) != len(tt.wantConditionTypes) {
				t.Fatalf("got %d conditions, want %d", len(buffer.Status.Conditions), len(tt.wantConditionTypes))
			}

			conditionTypesMap := make(map[string]bool)
			for _, cond := range buffer.Status.Conditions {
				conditionTypesMap[cond.Type] = true
			}

			for _, wantType := range tt.wantConditionTypes {
				if !conditionTypesMap[wantType] {
					t.Errorf("missing expected condition type %q", wantType)
				}
			}
		})
	}
}

func TestUpdateBufferStatusToFailedProvisioning(t *testing.T) {
	tests := []struct {
		name               string
		existingConditions []metav1.Condition
		wantConditionTypes []string
	}{
		{
			name:               "adds Provisioning=False when no conditions exist",
			existingConditions: nil,
			wantConditionTypes: []string{capacitybuffer.ProvisioningCondition},
		},
		{
			name: "preserves existing ReadyForProvisioning condition",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.ReadyForProvisioningCondition, Status: metav1.ConditionTrue},
			},
			wantConditionTypes: []string{capacitybuffer.ReadyForProvisioningCondition, capacitybuffer.ProvisioningCondition},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &v1.CapacityBuffer{
				Status: v1.CapacityBufferStatus{
					Conditions: tt.existingConditions,
				},
			}

			UpdateBufferStatusToFailedProvisioning(buffer, "TestReason", "Test error message")

			if len(buffer.Status.Conditions) != len(tt.wantConditionTypes) {
				t.Fatalf("got %d conditions, want %d", len(buffer.Status.Conditions), len(tt.wantConditionTypes))
			}

			conditionTypesMap := make(map[string]bool)
			for _, cond := range buffer.Status.Conditions {
				conditionTypesMap[cond.Type] = true
			}

			for _, wantType := range tt.wantConditionTypes {
				if !conditionTypesMap[wantType] {
					t.Errorf("missing expected condition type %q", wantType)
				}
			}
		})
	}
}

func TestUpdateBufferStatusToSuccessfullyProvisioning(t *testing.T) {
	tests := []struct {
		name               string
		existingConditions []metav1.Condition
		wantConditionTypes []string
	}{
		{
			name:               "adds Provisioning=True when no conditions exist",
			existingConditions: nil,
			wantConditionTypes: []string{capacitybuffer.ProvisioningCondition},
		},
		{
			name: "preserves existing ReadyForProvisioning condition",
			existingConditions: []metav1.Condition{
				{Type: capacitybuffer.ReadyForProvisioningCondition, Status: metav1.ConditionTrue},
			},
			wantConditionTypes: []string{capacitybuffer.ReadyForProvisioningCondition, capacitybuffer.ProvisioningCondition},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &v1.CapacityBuffer{
				Status: v1.CapacityBufferStatus{
					Conditions: tt.existingConditions,
				},
			}

			UpdateBufferStatusToSuccessfullyProvisioning(buffer, "TestReason")

			if len(buffer.Status.Conditions) != len(tt.wantConditionTypes) {
				t.Fatalf("got %d conditions, want %d", len(buffer.Status.Conditions), len(tt.wantConditionTypes))
			}

			conditionTypesMap := make(map[string]bool)
			for _, cond := range buffer.Status.Conditions {
				conditionTypesMap[cond.Type] = true
			}

			for _, wantType := range tt.wantConditionTypes {
				if !conditionTypesMap[wantType] {
					t.Errorf("missing expected condition type %q", wantType)
				}
			}
		})
	}
}
