// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInstanceAgentConfigDetails Configuration options for the Oracle Cloud Agent software running on the instance.
type UpdateInstanceAgentConfigDetails struct {

	// Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the
	// monitoring plugins.
	// These are the monitoring plugins: Compute Instance Monitoring
	// and Custom Logs Monitoring.
	// The monitoring plugins are controlled by this parameter and by the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You
	// can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig`
	// object.
	IsMonitoringDisabled *bool `mandatory:"false" json:"isMonitoringDisabled"`

	// Whether Oracle Cloud Agent can run all the available management plugins.
	// These are the management plugins: OS Management Service Agent and Compute Instance
	// Run Command.
	// The management plugins are controlled by this parameter and by the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isManagementDisabled` is false, all of the management plugins are enabled. You
	// can optionally disable individual management plugins by providing a value in the `pluginsConfig`
	// object.
	IsManagementDisabled *bool `mandatory:"false" json:"isManagementDisabled"`

	// Whether Oracle Cloud Agent can run all the available plugins.
	// This includes the management and monitoring plugins.
	// To get a list of available plugins, use the
	// ListInstanceagentAvailablePlugins
	// operation in the Oracle Cloud Agent API. For more information about the available plugins, see
	// Managing Plugins with Oracle Cloud Agent (https://docs.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
	AreAllPluginsDisabled *bool `mandatory:"false" json:"areAllPluginsDisabled"`

	// The configuration of plugins associated with this instance.
	PluginsConfig []InstanceAgentPluginConfigDetails `mandatory:"false" json:"pluginsConfig"`
}

func (m UpdateInstanceAgentConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInstanceAgentConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
