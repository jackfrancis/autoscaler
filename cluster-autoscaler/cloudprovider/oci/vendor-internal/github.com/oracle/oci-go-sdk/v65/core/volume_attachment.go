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
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VolumeAttachment A base object for all types of attachments between a storage volume and an instance.
// For specific details about iSCSI attachments, see
// IScsiVolumeAttachment.
// For general information about volume attachments, see
// Overview of Block Volume Storage (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type VolumeAttachment interface {

	// The availability domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	GetAvailabilityDomain() *string

	// The OCID of the compartment.
	GetCompartmentId() *string

	// The OCID of the volume attachment.
	GetId() *string

	// The OCID of the instance the volume is attached to.
	GetInstanceId() *string

	// The current state of the volume attachment.
	GetLifecycleState() VolumeAttachmentLifecycleStateEnum

	// The date and time the volume was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The OCID of the volume.
	GetVolumeId() *string

	// The device name.
	GetDevice() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// Whether the attachment was created in read-only mode.
	GetIsReadOnly() *bool

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	GetIsShareable() *bool

	// Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not.
	GetIsPvEncryptionInTransitEnabled() *bool

	// Whether the Iscsi or Paravirtualized attachment is multipath or not, it is not applicable to NVMe attachment.
	GetIsMultipath() *bool

	// The iscsi login state of the volume attachment. For a Iscsi volume attachment,
	// all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
	GetIscsiLoginState() VolumeAttachmentIscsiLoginStateEnum

	// Flag indicating if this volume was created for the customer as part of a simplified launch.
	// Used to determine whether the volume requires deletion on instance termination.
	GetIsVolumeCreatedDuringLaunch() *bool
}

type volumeattachment struct {
	JsonData                       []byte
	Device                         *string                             `mandatory:"false" json:"device"`
	DisplayName                    *string                             `mandatory:"false" json:"displayName"`
	IsReadOnly                     *bool                               `mandatory:"false" json:"isReadOnly"`
	IsShareable                    *bool                               `mandatory:"false" json:"isShareable"`
	IsPvEncryptionInTransitEnabled *bool                               `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`
	IsMultipath                    *bool                               `mandatory:"false" json:"isMultipath"`
	IscsiLoginState                VolumeAttachmentIscsiLoginStateEnum `mandatory:"false" json:"iscsiLoginState,omitempty"`
	IsVolumeCreatedDuringLaunch    *bool                               `mandatory:"false" json:"isVolumeCreatedDuringLaunch"`
	AvailabilityDomain             *string                             `mandatory:"true" json:"availabilityDomain"`
	CompartmentId                  *string                             `mandatory:"true" json:"compartmentId"`
	Id                             *string                             `mandatory:"true" json:"id"`
	InstanceId                     *string                             `mandatory:"true" json:"instanceId"`
	LifecycleState                 VolumeAttachmentLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	TimeCreated                    *common.SDKTime                     `mandatory:"true" json:"timeCreated"`
	VolumeId                       *string                             `mandatory:"true" json:"volumeId"`
	AttachmentType                 string                              `json:"attachmentType"`
}

// UnmarshalJSON unmarshals json
func (m *volumeattachment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervolumeattachment volumeattachment
	s := struct {
		Model Unmarshalervolumeattachment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.CompartmentId = s.Model.CompartmentId
	m.Id = s.Model.Id
	m.InstanceId = s.Model.InstanceId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.VolumeId = s.Model.VolumeId
	m.Device = s.Model.Device
	m.DisplayName = s.Model.DisplayName
	m.IsReadOnly = s.Model.IsReadOnly
	m.IsShareable = s.Model.IsShareable
	m.IsPvEncryptionInTransitEnabled = s.Model.IsPvEncryptionInTransitEnabled
	m.IsMultipath = s.Model.IsMultipath
	m.IscsiLoginState = s.Model.IscsiLoginState
	m.IsVolumeCreatedDuringLaunch = s.Model.IsVolumeCreatedDuringLaunch
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *volumeattachment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttachmentType {
	case "iscsi":
		mm := IScsiVolumeAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "emulated":
		mm := EmulatedVolumeAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "paravirtualized":
		mm := ParavirtualizedVolumeAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for VolumeAttachment: %s.", m.AttachmentType)
		return *m, nil
	}
}

// GetDevice returns Device
func (m volumeattachment) GetDevice() *string {
	return m.Device
}

// GetDisplayName returns DisplayName
func (m volumeattachment) GetDisplayName() *string {
	return m.DisplayName
}

// GetIsReadOnly returns IsReadOnly
func (m volumeattachment) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

// GetIsShareable returns IsShareable
func (m volumeattachment) GetIsShareable() *bool {
	return m.IsShareable
}

// GetIsPvEncryptionInTransitEnabled returns IsPvEncryptionInTransitEnabled
func (m volumeattachment) GetIsPvEncryptionInTransitEnabled() *bool {
	return m.IsPvEncryptionInTransitEnabled
}

// GetIsMultipath returns IsMultipath
func (m volumeattachment) GetIsMultipath() *bool {
	return m.IsMultipath
}

// GetIscsiLoginState returns IscsiLoginState
func (m volumeattachment) GetIscsiLoginState() VolumeAttachmentIscsiLoginStateEnum {
	return m.IscsiLoginState
}

// GetIsVolumeCreatedDuringLaunch returns IsVolumeCreatedDuringLaunch
func (m volumeattachment) GetIsVolumeCreatedDuringLaunch() *bool {
	return m.IsVolumeCreatedDuringLaunch
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m volumeattachment) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetCompartmentId returns CompartmentId
func (m volumeattachment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetId returns Id
func (m volumeattachment) GetId() *string {
	return m.Id
}

// GetInstanceId returns InstanceId
func (m volumeattachment) GetInstanceId() *string {
	return m.InstanceId
}

// GetLifecycleState returns LifecycleState
func (m volumeattachment) GetLifecycleState() VolumeAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m volumeattachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetVolumeId returns VolumeId
func (m volumeattachment) GetVolumeId() *string {
	return m.VolumeId
}

func (m volumeattachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m volumeattachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVolumeAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVolumeAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVolumeAttachmentIscsiLoginStateEnum(string(m.IscsiLoginState)); !ok && m.IscsiLoginState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IscsiLoginState: %s. Supported values are: %s.", m.IscsiLoginState, strings.Join(GetVolumeAttachmentIscsiLoginStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VolumeAttachmentLifecycleStateEnum Enum with underlying type: string
type VolumeAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for VolumeAttachmentLifecycleStateEnum
const (
	VolumeAttachmentLifecycleStateAttaching VolumeAttachmentLifecycleStateEnum = "ATTACHING"
	VolumeAttachmentLifecycleStateAttached  VolumeAttachmentLifecycleStateEnum = "ATTACHED"
	VolumeAttachmentLifecycleStateDetaching VolumeAttachmentLifecycleStateEnum = "DETACHING"
	VolumeAttachmentLifecycleStateDetached  VolumeAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingVolumeAttachmentLifecycleStateEnum = map[string]VolumeAttachmentLifecycleStateEnum{
	"ATTACHING": VolumeAttachmentLifecycleStateAttaching,
	"ATTACHED":  VolumeAttachmentLifecycleStateAttached,
	"DETACHING": VolumeAttachmentLifecycleStateDetaching,
	"DETACHED":  VolumeAttachmentLifecycleStateDetached,
}

var mappingVolumeAttachmentLifecycleStateEnumLowerCase = map[string]VolumeAttachmentLifecycleStateEnum{
	"attaching": VolumeAttachmentLifecycleStateAttaching,
	"attached":  VolumeAttachmentLifecycleStateAttached,
	"detaching": VolumeAttachmentLifecycleStateDetaching,
	"detached":  VolumeAttachmentLifecycleStateDetached,
}

// GetVolumeAttachmentLifecycleStateEnumValues Enumerates the set of values for VolumeAttachmentLifecycleStateEnum
func GetVolumeAttachmentLifecycleStateEnumValues() []VolumeAttachmentLifecycleStateEnum {
	values := make([]VolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingVolumeAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVolumeAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for VolumeAttachmentLifecycleStateEnum
func GetVolumeAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ATTACHING",
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

// GetMappingVolumeAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVolumeAttachmentLifecycleStateEnum(val string) (VolumeAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingVolumeAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VolumeAttachmentIscsiLoginStateEnum Enum with underlying type: string
type VolumeAttachmentIscsiLoginStateEnum string

// Set of constants representing the allowable values for VolumeAttachmentIscsiLoginStateEnum
const (
	VolumeAttachmentIscsiLoginStateUnknown         VolumeAttachmentIscsiLoginStateEnum = "UNKNOWN"
	VolumeAttachmentIscsiLoginStateLoggingIn       VolumeAttachmentIscsiLoginStateEnum = "LOGGING_IN"
	VolumeAttachmentIscsiLoginStateLoginSucceeded  VolumeAttachmentIscsiLoginStateEnum = "LOGIN_SUCCEEDED"
	VolumeAttachmentIscsiLoginStateLoginFailed     VolumeAttachmentIscsiLoginStateEnum = "LOGIN_FAILED"
	VolumeAttachmentIscsiLoginStateLoggingOut      VolumeAttachmentIscsiLoginStateEnum = "LOGGING_OUT"
	VolumeAttachmentIscsiLoginStateLogoutSucceeded VolumeAttachmentIscsiLoginStateEnum = "LOGOUT_SUCCEEDED"
	VolumeAttachmentIscsiLoginStateLogoutFailed    VolumeAttachmentIscsiLoginStateEnum = "LOGOUT_FAILED"
)

var mappingVolumeAttachmentIscsiLoginStateEnum = map[string]VolumeAttachmentIscsiLoginStateEnum{
	"UNKNOWN":          VolumeAttachmentIscsiLoginStateUnknown,
	"LOGGING_IN":       VolumeAttachmentIscsiLoginStateLoggingIn,
	"LOGIN_SUCCEEDED":  VolumeAttachmentIscsiLoginStateLoginSucceeded,
	"LOGIN_FAILED":     VolumeAttachmentIscsiLoginStateLoginFailed,
	"LOGGING_OUT":      VolumeAttachmentIscsiLoginStateLoggingOut,
	"LOGOUT_SUCCEEDED": VolumeAttachmentIscsiLoginStateLogoutSucceeded,
	"LOGOUT_FAILED":    VolumeAttachmentIscsiLoginStateLogoutFailed,
}

var mappingVolumeAttachmentIscsiLoginStateEnumLowerCase = map[string]VolumeAttachmentIscsiLoginStateEnum{
	"unknown":          VolumeAttachmentIscsiLoginStateUnknown,
	"logging_in":       VolumeAttachmentIscsiLoginStateLoggingIn,
	"login_succeeded":  VolumeAttachmentIscsiLoginStateLoginSucceeded,
	"login_failed":     VolumeAttachmentIscsiLoginStateLoginFailed,
	"logging_out":      VolumeAttachmentIscsiLoginStateLoggingOut,
	"logout_succeeded": VolumeAttachmentIscsiLoginStateLogoutSucceeded,
	"logout_failed":    VolumeAttachmentIscsiLoginStateLogoutFailed,
}

// GetVolumeAttachmentIscsiLoginStateEnumValues Enumerates the set of values for VolumeAttachmentIscsiLoginStateEnum
func GetVolumeAttachmentIscsiLoginStateEnumValues() []VolumeAttachmentIscsiLoginStateEnum {
	values := make([]VolumeAttachmentIscsiLoginStateEnum, 0)
	for _, v := range mappingVolumeAttachmentIscsiLoginStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVolumeAttachmentIscsiLoginStateEnumStringValues Enumerates the set of values in String for VolumeAttachmentIscsiLoginStateEnum
func GetVolumeAttachmentIscsiLoginStateEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"LOGGING_IN",
		"LOGIN_SUCCEEDED",
		"LOGIN_FAILED",
		"LOGGING_OUT",
		"LOGOUT_SUCCEEDED",
		"LOGOUT_FAILED",
	}
}

// GetMappingVolumeAttachmentIscsiLoginStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVolumeAttachmentIscsiLoginStateEnum(val string) (VolumeAttachmentIscsiLoginStateEnum, bool) {
	enum, ok := mappingVolumeAttachmentIscsiLoginStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
