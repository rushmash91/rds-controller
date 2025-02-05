// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type ActivityStreamMode string

const (
	ActivityStreamMode_async ActivityStreamMode = "async"
	ActivityStreamMode_sync  ActivityStreamMode = "sync"
)

type ActivityStreamPolicyStatus string

const (
	ActivityStreamPolicyStatus_locked           ActivityStreamPolicyStatus = "locked"
	ActivityStreamPolicyStatus_locking_policy   ActivityStreamPolicyStatus = "locking-policy"
	ActivityStreamPolicyStatus_unlocked         ActivityStreamPolicyStatus = "unlocked"
	ActivityStreamPolicyStatus_unlocking_policy ActivityStreamPolicyStatus = "unlocking-policy"
)

type ActivityStreamStatus string

const (
	ActivityStreamStatus_started  ActivityStreamStatus = "started"
	ActivityStreamStatus_starting ActivityStreamStatus = "starting"
	ActivityStreamStatus_stopped  ActivityStreamStatus = "stopped"
	ActivityStreamStatus_stopping ActivityStreamStatus = "stopping"
)

type ApplyMethod string

const (
	ApplyMethod_immediate      ApplyMethod = "immediate"
	ApplyMethod_pending_reboot ApplyMethod = "pending-reboot"
)

type AuditPolicyState string

const (
	AuditPolicyState_locked   AuditPolicyState = "locked"
	AuditPolicyState_unlocked AuditPolicyState = "unlocked"
)

type AuthScheme string

const (
	AuthScheme_SECRETS AuthScheme = "SECRETS"
)

type AutomationMode string

const (
	AutomationMode_all_paused AutomationMode = "all-paused"
	AutomationMode_full       AutomationMode = "full"
)

type ClientPasswordAuthType string

const (
	ClientPasswordAuthType_MYSQL_NATIVE_PASSWORD     ClientPasswordAuthType = "MYSQL_NATIVE_PASSWORD"
	ClientPasswordAuthType_POSTGRES_MD5              ClientPasswordAuthType = "POSTGRES_MD5"
	ClientPasswordAuthType_POSTGRES_SCRAM_SHA_256    ClientPasswordAuthType = "POSTGRES_SCRAM_SHA_256"
	ClientPasswordAuthType_SQL_SERVER_AUTHENTICATION ClientPasswordAuthType = "SQL_SERVER_AUTHENTICATION"
)

type ClusterScalabilityType string

const (
	ClusterScalabilityType_limitless ClusterScalabilityType = "limitless"
	ClusterScalabilityType_standard  ClusterScalabilityType = "standard"
)

type CustomEngineVersionStatus string

const (
	CustomEngineVersionStatus_available               CustomEngineVersionStatus = "available"
	CustomEngineVersionStatus_inactive                CustomEngineVersionStatus = "inactive"
	CustomEngineVersionStatus_inactive_except_restore CustomEngineVersionStatus = "inactive-except-restore"
)

type DBProxyEndpointStatus string

const (
	DBProxyEndpointStatus_available                    DBProxyEndpointStatus = "available"
	DBProxyEndpointStatus_creating                     DBProxyEndpointStatus = "creating"
	DBProxyEndpointStatus_deleting                     DBProxyEndpointStatus = "deleting"
	DBProxyEndpointStatus_incompatible_network         DBProxyEndpointStatus = "incompatible-network"
	DBProxyEndpointStatus_insufficient_resource_limits DBProxyEndpointStatus = "insufficient-resource-limits"
	DBProxyEndpointStatus_modifying                    DBProxyEndpointStatus = "modifying"
)

type DBProxyEndpointTargetRole string

const (
	DBProxyEndpointTargetRole_READ_ONLY  DBProxyEndpointTargetRole = "READ_ONLY"
	DBProxyEndpointTargetRole_READ_WRITE DBProxyEndpointTargetRole = "READ_WRITE"
)

type DBProxyStatus_SDK string

const (
	DBProxyStatus_SDK_available                    DBProxyStatus_SDK = "available"
	DBProxyStatus_SDK_creating                     DBProxyStatus_SDK = "creating"
	DBProxyStatus_SDK_deleting                     DBProxyStatus_SDK = "deleting"
	DBProxyStatus_SDK_incompatible_network         DBProxyStatus_SDK = "incompatible-network"
	DBProxyStatus_SDK_insufficient_resource_limits DBProxyStatus_SDK = "insufficient-resource-limits"
	DBProxyStatus_SDK_modifying                    DBProxyStatus_SDK = "modifying"
	DBProxyStatus_SDK_reactivating                 DBProxyStatus_SDK = "reactivating"
	DBProxyStatus_SDK_suspended                    DBProxyStatus_SDK = "suspended"
	DBProxyStatus_SDK_suspending                   DBProxyStatus_SDK = "suspending"
)

type DatabaseInsightsMode string

const (
	DatabaseInsightsMode_advanced DatabaseInsightsMode = "advanced"
	DatabaseInsightsMode_standard DatabaseInsightsMode = "standard"
)

type EngineFamily string

const (
	EngineFamily_MYSQL      EngineFamily = "MYSQL"
	EngineFamily_POSTGRESQL EngineFamily = "POSTGRESQL"
	EngineFamily_SQLSERVER  EngineFamily = "SQLSERVER"
)

type ExportSourceType string

const (
	ExportSourceType_CLUSTER  ExportSourceType = "CLUSTER"
	ExportSourceType_SNAPSHOT ExportSourceType = "SNAPSHOT"
)

type FailoverStatus string

const (
	FailoverStatus_cancelling   FailoverStatus = "cancelling"
	FailoverStatus_failing_over FailoverStatus = "failing-over"
	FailoverStatus_pending      FailoverStatus = "pending"
)

type GlobalClusterMemberSynchronizationStatus string

const (
	GlobalClusterMemberSynchronizationStatus_connected      GlobalClusterMemberSynchronizationStatus = "connected"
	GlobalClusterMemberSynchronizationStatus_pending_resync GlobalClusterMemberSynchronizationStatus = "pending-resync"
)

type IAMAuthMode string

const (
	IAMAuthMode_DISABLED IAMAuthMode = "DISABLED"
	IAMAuthMode_ENABLED  IAMAuthMode = "ENABLED"
	IAMAuthMode_REQUIRED IAMAuthMode = "REQUIRED"
)

type IntegrationStatus string

const (
	IntegrationStatus_active          IntegrationStatus = "active"
	IntegrationStatus_creating        IntegrationStatus = "creating"
	IntegrationStatus_deleting        IntegrationStatus = "deleting"
	IntegrationStatus_failed          IntegrationStatus = "failed"
	IntegrationStatus_modifying       IntegrationStatus = "modifying"
	IntegrationStatus_needs_attention IntegrationStatus = "needs_attention"
	IntegrationStatus_syncing         IntegrationStatus = "syncing"
)

type LimitlessDatabaseStatus string

const (
	LimitlessDatabaseStatus_active                 LimitlessDatabaseStatus = "active"
	LimitlessDatabaseStatus_disabled               LimitlessDatabaseStatus = "disabled"
	LimitlessDatabaseStatus_disabling              LimitlessDatabaseStatus = "disabling"
	LimitlessDatabaseStatus_enabled                LimitlessDatabaseStatus = "enabled"
	LimitlessDatabaseStatus_enabling               LimitlessDatabaseStatus = "enabling"
	LimitlessDatabaseStatus_error                  LimitlessDatabaseStatus = "error"
	LimitlessDatabaseStatus_modifying_max_capacity LimitlessDatabaseStatus = "modifying-max-capacity"
	LimitlessDatabaseStatus_not_in_use             LimitlessDatabaseStatus = "not-in-use"
)

type LocalWriteForwardingStatus string

const (
	LocalWriteForwardingStatus_disabled  LocalWriteForwardingStatus = "disabled"
	LocalWriteForwardingStatus_disabling LocalWriteForwardingStatus = "disabling"
	LocalWriteForwardingStatus_enabled   LocalWriteForwardingStatus = "enabled"
	LocalWriteForwardingStatus_enabling  LocalWriteForwardingStatus = "enabling"
	LocalWriteForwardingStatus_requested LocalWriteForwardingStatus = "requested"
)

type ReplicaMode string

const (
	ReplicaMode_mounted        ReplicaMode = "mounted"
	ReplicaMode_open_read_only ReplicaMode = "open-read-only"
)

type SourceType string

const (
	SourceType_blue_green_deployment SourceType = "blue-green-deployment"
	SourceType_custom_engine_version SourceType = "custom-engine-version"
	SourceType_db_cluster            SourceType = "db-cluster"
	SourceType_db_cluster_snapshot   SourceType = "db-cluster-snapshot"
	SourceType_db_instance           SourceType = "db-instance"
	SourceType_db_parameter_group    SourceType = "db-parameter-group"
	SourceType_db_proxy              SourceType = "db-proxy"
	SourceType_db_security_group     SourceType = "db-security-group"
	SourceType_db_snapshot           SourceType = "db-snapshot"
)

type TargetHealthReason string

const (
	TargetHealthReason_AUTH_FAILURE              TargetHealthReason = "AUTH_FAILURE"
	TargetHealthReason_CONNECTION_FAILED         TargetHealthReason = "CONNECTION_FAILED"
	TargetHealthReason_INVALID_REPLICATION_STATE TargetHealthReason = "INVALID_REPLICATION_STATE"
	TargetHealthReason_PENDING_PROXY_CAPACITY    TargetHealthReason = "PENDING_PROXY_CAPACITY"
	TargetHealthReason_UNREACHABLE               TargetHealthReason = "UNREACHABLE"
)

type TargetRole string

const (
	TargetRole_READ_ONLY  TargetRole = "READ_ONLY"
	TargetRole_READ_WRITE TargetRole = "READ_WRITE"
	TargetRole_UNKNOWN    TargetRole = "UNKNOWN"
)

type TargetState string

const (
	TargetState_AVAILABLE   TargetState = "AVAILABLE"
	TargetState_REGISTERING TargetState = "REGISTERING"
	TargetState_UNAVAILABLE TargetState = "UNAVAILABLE"
)

type TargetType string

const (
	TargetType_RDS_INSTANCE            TargetType = "RDS_INSTANCE"
	TargetType_RDS_SERVERLESS_ENDPOINT TargetType = "RDS_SERVERLESS_ENDPOINT"
	TargetType_TRACKED_CLUSTER         TargetType = "TRACKED_CLUSTER"
)

type WriteForwardingStatus string

const (
	WriteForwardingStatus_disabled  WriteForwardingStatus = "disabled"
	WriteForwardingStatus_disabling WriteForwardingStatus = "disabling"
	WriteForwardingStatus_enabled   WriteForwardingStatus = "enabled"
	WriteForwardingStatus_enabling  WriteForwardingStatus = "enabling"
	WriteForwardingStatus_unknown   WriteForwardingStatus = "unknown"
)
