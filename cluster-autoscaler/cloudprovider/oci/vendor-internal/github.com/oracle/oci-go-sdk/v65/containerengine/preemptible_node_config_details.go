// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PreemptibleNodeConfigDetails Configuration options for preemptible nodes.
type PreemptibleNodeConfigDetails struct {
	PreemptionAction PreemptionAction `mandatory:"true" json:"preemptionAction"`
}

func (m PreemptibleNodeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PreemptibleNodeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PreemptibleNodeConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PreemptionAction preemptionaction `json:"preemptionAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PreemptionAction.UnmarshalPolymorphicJSON(model.PreemptionAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PreemptionAction = nn.(PreemptionAction)
	} else {
		m.PreemptionAction = nil
	}

	return
}
