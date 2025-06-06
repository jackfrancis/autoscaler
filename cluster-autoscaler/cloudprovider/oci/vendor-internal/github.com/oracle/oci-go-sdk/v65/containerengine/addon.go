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
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Addon The properties that define an addon.
type Addon struct {

	// The name of the addon.
	Name *string `mandatory:"true" json:"name"`

	// The state of the addon.
	LifecycleState AddonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// selected addon version, or null indicates autoUpdate
	Version *string `mandatory:"false" json:"version"`

	// current installed version of the addon
	CurrentInstalledVersion *string `mandatory:"false" json:"currentInstalledVersion"`

	// The time the cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Addon configuration details.
	Configurations []AddonConfiguration `mandatory:"false" json:"configurations"`

	// The error info of the addon.
	AddonError *AddonError `mandatory:"false" json:"addonError"`
}

func (m Addon) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Addon) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAddonLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
