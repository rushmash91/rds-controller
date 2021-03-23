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

type AccountQuota struct {
	AccountQuotaName *string `json:"accountQuotaName,omitempty"`
}

type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

type AvailableProcessorFeature struct {
	AllowedValues *string `json:"allowedValues,omitempty"`
	DefaultValue  *string `json:"defaultValue,omitempty"`
	Name          *string `json:"name,omitempty"`
}

type Certificate struct {
	CertificateARN        *string `json:"certificateARN,omitempty"`
	CertificateIdentifier *string `json:"certificateIdentifier,omitempty"`
	CertificateType       *string `json:"certificateType,omitempty"`
	Thumbprint            *string `json:"thumbprint,omitempty"`
}

type CharacterSet struct {
	CharacterSetDescription *string `json:"characterSetDescription,omitempty"`
	CharacterSetName        *string `json:"characterSetName,omitempty"`
}

type ClusterPendingModifiedValues struct {
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`
	EngineVersion       *string `json:"engineVersion,omitempty"`
	MasterUserPassword  *string `json:"masterUserPassword,omitempty"`
}

type ConnectionPoolConfiguration struct {
	ConnectionBorrowTimeout   *int64  `json:"connectionBorrowTimeout,omitempty"`
	InitQuery                 *string `json:"initQuery,omitempty"`
	MaxConnectionsPercent     *int64  `json:"maxConnectionsPercent,omitempty"`
	MaxIdleConnectionsPercent *int64  `json:"maxIdleConnectionsPercent,omitempty"`
}

type ConnectionPoolConfigurationInfo struct {
	InitQuery *string `json:"initQuery,omitempty"`
}

type CustomAvailabilityZone struct {
	CustomAvailabilityZoneID     *string `json:"customAvailabilityZoneID,omitempty"`
	CustomAvailabilityZoneName   *string `json:"customAvailabilityZoneName,omitempty"`
	CustomAvailabilityZoneStatus *string `json:"customAvailabilityZoneStatus,omitempty"`
}

type DBCluster struct {
	ActivityStreamKinesisStreamName *string `json:"activityStreamKinesisStreamName,omitempty"`
	ActivityStreamKMSKeyID          *string `json:"activityStreamKMSKeyID,omitempty"`
	AllocatedStorage                *int64  `json:"allocatedStorage,omitempty"`
	BackupRetentionPeriod           *int64  `json:"backupRetentionPeriod,omitempty"`
	Capacity                        *int64  `json:"capacity,omitempty"`
	CharacterSetName                *string `json:"characterSetName,omitempty"`
	CloneGroupID                    *string `json:"cloneGroupID,omitempty"`
	DBClusterARN                    *string `json:"dbClusterARN,omitempty"`
	DBClusterIdentifier             *string `json:"dbClusterIdentifier,omitempty"`
	DBClusterParameterGroup         *string `json:"dbClusterParameterGroup,omitempty"`
	DBSubnetGroup                   *string `json:"dbSubnetGroup,omitempty"`
	DatabaseName                    *string `json:"databaseName,omitempty"`
	DBClusterResourceID             *string `json:"dbClusterResourceID,omitempty"`
	Endpoint                        *string `json:"endpoint,omitempty"`
	Engine                          *string `json:"engine,omitempty"`
	EngineMode                      *string `json:"engineMode,omitempty"`
	EngineVersion                   *string `json:"engineVersion,omitempty"`
	HostedZoneID                    *string `json:"hostedZoneID,omitempty"`
	KMSKeyID                        *string `json:"kmsKeyID,omitempty"`
	MasterUsername                  *string `json:"masterUsername,omitempty"`
	PercentProgress                 *string `json:"percentProgress,omitempty"`
	Port                            *int64  `json:"port,omitempty"`
	PreferredBackupWindow           *string `json:"preferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow      *string `json:"preferredMaintenanceWindow,omitempty"`
	ReaderEndpoint                  *string `json:"readerEndpoint,omitempty"`
	ReplicationSourceIdentifier     *string `json:"replicationSourceIdentifier,omitempty"`
	Status                          *string `json:"status,omitempty"`
	TagList                         []*Tag  `json:"tagList,omitempty"`
}

type DBClusterEndpoint struct {
	CustomEndpointType                  *string `json:"customEndpointType,omitempty"`
	DBClusterEndpointARN                *string `json:"dbClusterEndpointARN,omitempty"`
	DBClusterEndpointIdentifier         *string `json:"dbClusterEndpointIdentifier,omitempty"`
	DBClusterEndpointResourceIdentifier *string `json:"dbClusterEndpointResourceIdentifier,omitempty"`
	DBClusterIdentifier                 *string `json:"dbClusterIdentifier,omitempty"`
	Endpoint                            *string `json:"endpoint,omitempty"`
	EndpointType                        *string `json:"endpointType,omitempty"`
	Status                              *string `json:"status,omitempty"`
}

type DBClusterMember struct {
	DBClusterParameterGroupStatus *string `json:"dbClusterParameterGroupStatus,omitempty"`
	DBInstanceIdentifier          *string `json:"dbInstanceIdentifier,omitempty"`
	PromotionTier                 *int64  `json:"promotionTier,omitempty"`
}

type DBClusterOptionGroupStatus struct {
	DBClusterOptionGroupName *string `json:"dbClusterOptionGroupName,omitempty"`
	Status                   *string `json:"status,omitempty"`
}

type DBClusterParameterGroup struct {
	DBClusterParameterGroupARN  *string `json:"dbClusterParameterGroupARN,omitempty"`
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`
	DBParameterGroupFamily      *string `json:"dbParameterGroupFamily,omitempty"`
	Description                 *string `json:"description,omitempty"`
}

type DBClusterRole struct {
	FeatureName *string `json:"featureName,omitempty"`
	RoleARN     *string `json:"roleARN,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type DBClusterSnapshot struct {
	DBClusterIdentifier         *string `json:"dbClusterIdentifier,omitempty"`
	DBClusterSnapshotARN        *string `json:"dbClusterSnapshotARN,omitempty"`
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty"`
	Engine                      *string `json:"engine,omitempty"`
	EngineVersion               *string `json:"engineVersion,omitempty"`
	KMSKeyID                    *string `json:"kmsKeyID,omitempty"`
	LicenseModel                *string `json:"licenseModel,omitempty"`
	MasterUsername              *string `json:"masterUsername,omitempty"`
	SnapshotType                *string `json:"snapshotType,omitempty"`
	SourceDBClusterSnapshotARN  *string `json:"sourceDBClusterSnapshotARN,omitempty"`
	Status                      *string `json:"status,omitempty"`
	TagList                     []*Tag  `json:"tagList,omitempty"`
	VPCID                       *string `json:"vpcID,omitempty"`
}

type DBClusterSnapshotAttribute struct {
	AttributeName *string `json:"attributeName,omitempty"`
}

type DBClusterSnapshotAttributesResult struct {
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty"`
}

type DBEngineVersion struct {
	DBEngineDescription        *string `json:"dbEngineDescription,omitempty"`
	DBEngineVersionDescription *string `json:"dbEngineVersionDescription,omitempty"`
	DBParameterGroupFamily     *string `json:"dbParameterGroupFamily,omitempty"`
	Engine                     *string `json:"engine,omitempty"`
	EngineVersion              *string `json:"engineVersion,omitempty"`
	Status                     *string `json:"status,omitempty"`
}

type DBInstance struct {
	AvailabilityZone                      *string            `json:"availabilityZone,omitempty"`
	CACertificateIdentifier               *string            `json:"caCertificateIdentifier,omitempty"`
	CharacterSetName                      *string            `json:"characterSetName,omitempty"`
	DBClusterIdentifier                   *string            `json:"dbClusterIdentifier,omitempty"`
	DBInstanceARN                         *string            `json:"dbInstanceARN,omitempty"`
	DBInstanceClass                       *string            `json:"dbInstanceClass,omitempty"`
	DBInstanceIdentifier                  *string            `json:"dbInstanceIdentifier,omitempty"`
	DBInstanceStatus                      *string            `json:"dbInstanceStatus,omitempty"`
	DBName                                *string            `json:"dbName,omitempty"`
	DBSubnetGroup                         *DBSubnetGroup_SDK `json:"dbSubnetGroup,omitempty"`
	DBIResourceID                         *string            `json:"dbiResourceID,omitempty"`
	Engine                                *string            `json:"engine,omitempty"`
	EngineVersion                         *string            `json:"engineVersion,omitempty"`
	EnhancedMonitoringResourceARN         *string            `json:"enhancedMonitoringResourceARN,omitempty"`
	IOPS                                  *int64             `json:"iops,omitempty"`
	KMSKeyID                              *string            `json:"kmsKeyID,omitempty"`
	LicenseModel                          *string            `json:"licenseModel,omitempty"`
	MasterUsername                        *string            `json:"masterUsername,omitempty"`
	MaxAllocatedStorage                   *int64             `json:"maxAllocatedStorage,omitempty"`
	MonitoringInterval                    *int64             `json:"monitoringInterval,omitempty"`
	MonitoringRoleARN                     *string            `json:"monitoringRoleARN,omitempty"`
	NcharCharacterSetName                 *string            `json:"ncharCharacterSetName,omitempty"`
	PerformanceInsightsKMSKeyID           *string            `json:"performanceInsightsKMSKeyID,omitempty"`
	PerformanceInsightsRetentionPeriod    *int64             `json:"performanceInsightsRetentionPeriod,omitempty"`
	PreferredBackupWindow                 *string            `json:"preferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow            *string            `json:"preferredMaintenanceWindow,omitempty"`
	PromotionTier                         *int64             `json:"promotionTier,omitempty"`
	ReadReplicaSourceDBInstanceIdentifier *string            `json:"readReplicaSourceDBInstanceIdentifier,omitempty"`
	SecondaryAvailabilityZone             *string            `json:"secondaryAvailabilityZone,omitempty"`
	StorageType                           *string            `json:"storageType,omitempty"`
	TagList                               []*Tag             `json:"tagList,omitempty"`
	TDECredentialARN                      *string            `json:"tdeCredentialARN,omitempty"`
	Timezone                              *string            `json:"timezone,omitempty"`
}

type DBInstanceAutomatedBackup struct {
	AvailabilityZone              *string `json:"availabilityZone,omitempty"`
	BackupRetentionPeriod         *int64  `json:"backupRetentionPeriod,omitempty"`
	DBInstanceARN                 *string `json:"dbInstanceARN,omitempty"`
	DBInstanceAutomatedBackupsARN *string `json:"dbInstanceAutomatedBackupsARN,omitempty"`
	DBInstanceIdentifier          *string `json:"dbInstanceIdentifier,omitempty"`
	DBIResourceID                 *string `json:"dbiResourceID,omitempty"`
	Engine                        *string `json:"engine,omitempty"`
	EngineVersion                 *string `json:"engineVersion,omitempty"`
	IOPS                          *int64  `json:"iops,omitempty"`
	KMSKeyID                      *string `json:"kmsKeyID,omitempty"`
	LicenseModel                  *string `json:"licenseModel,omitempty"`
	MasterUsername                *string `json:"masterUsername,omitempty"`
	OptionGroupName               *string `json:"optionGroupName,omitempty"`
	Region                        *string `json:"region,omitempty"`
	Status                        *string `json:"status,omitempty"`
	StorageType                   *string `json:"storageType,omitempty"`
	TDECredentialARN              *string `json:"tdeCredentialARN,omitempty"`
	Timezone                      *string `json:"timezone,omitempty"`
	VPCID                         *string `json:"vpcID,omitempty"`
}

type DBInstanceAutomatedBackupsReplication struct {
	DBInstanceAutomatedBackupsARN *string `json:"dbInstanceAutomatedBackupsARN,omitempty"`
}

type DBInstanceRole struct {
	FeatureName *string `json:"featureName,omitempty"`
	RoleARN     *string `json:"roleARN,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type DBInstanceStatusInfo struct {
	Message    *string `json:"message,omitempty"`
	Status     *string `json:"status,omitempty"`
	StatusType *string `json:"statusType,omitempty"`
}

type DBParameterGroup struct {
	DBParameterGroupARN    *string `json:"dbParameterGroupARN,omitempty"`
	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`
	DBParameterGroupName   *string `json:"dbParameterGroupName,omitempty"`
	Description            *string `json:"description,omitempty"`
}

type DBParameterGroupStatus struct {
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty"`
	ParameterApplyStatus *string `json:"parameterApplyStatus,omitempty"`
}

type DBProxy struct {
	DBProxyARN   *string `json:"dbProxyARN,omitempty"`
	DBProxyName  *string `json:"dbProxyName,omitempty"`
	Endpoint     *string `json:"endpoint,omitempty"`
	EngineFamily *string `json:"engineFamily,omitempty"`
	RoleARN      *string `json:"roleARN,omitempty"`
}

type DBProxyTarget struct {
	Endpoint         *string `json:"endpoint,omitempty"`
	RdsResourceID    *string `json:"rdsResourceID,omitempty"`
	TargetARN        *string `json:"targetARN,omitempty"`
	TrackedClusterID *string `json:"trackedClusterID,omitempty"`
}

type DBProxyTargetGroup struct {
	DBProxyName     *string `json:"dbProxyName,omitempty"`
	Status          *string `json:"status,omitempty"`
	TargetGroupARN  *string `json:"targetGroupARN,omitempty"`
	TargetGroupName *string `json:"targetGroupName,omitempty"`
}

type DBSecurityGroup struct {
	DBSecurityGroupARN         *string `json:"dbSecurityGroupARN,omitempty"`
	DBSecurityGroupDescription *string `json:"dbSecurityGroupDescription,omitempty"`
	DBSecurityGroupName        *string `json:"dbSecurityGroupName,omitempty"`
	OwnerID                    *string `json:"ownerID,omitempty"`
	VPCID                      *string `json:"vpcID,omitempty"`
}

type DBSecurityGroupMembership struct {
	DBSecurityGroupName *string `json:"dbSecurityGroupName,omitempty"`
	Status              *string `json:"status,omitempty"`
}

type DBSnapshot struct {
	AvailabilityZone           *string `json:"availabilityZone,omitempty"`
	DBInstanceIdentifier       *string `json:"dbInstanceIdentifier,omitempty"`
	DBSnapshotARN              *string `json:"dbSnapshotARN,omitempty"`
	DBSnapshotIdentifier       *string `json:"dbSnapshotIdentifier,omitempty"`
	DBIResourceID              *string `json:"dbiResourceID,omitempty"`
	Engine                     *string `json:"engine,omitempty"`
	EngineVersion              *string `json:"engineVersion,omitempty"`
	IOPS                       *int64  `json:"iops,omitempty"`
	KMSKeyID                   *string `json:"kmsKeyID,omitempty"`
	LicenseModel               *string `json:"licenseModel,omitempty"`
	MasterUsername             *string `json:"masterUsername,omitempty"`
	OptionGroupName            *string `json:"optionGroupName,omitempty"`
	SnapshotType               *string `json:"snapshotType,omitempty"`
	SourceDBSnapshotIdentifier *string `json:"sourceDBSnapshotIdentifier,omitempty"`
	SourceRegion               *string `json:"sourceRegion,omitempty"`
	Status                     *string `json:"status,omitempty"`
	StorageType                *string `json:"storageType,omitempty"`
	TagList                    []*Tag  `json:"tagList,omitempty"`
	TDECredentialARN           *string `json:"tdeCredentialARN,omitempty"`
	Timezone                   *string `json:"timezone,omitempty"`
	VPCID                      *string `json:"vpcID,omitempty"`
}

type DBSnapshotAttribute struct {
	AttributeName *string `json:"attributeName,omitempty"`
}

type DBSnapshotAttributesResult struct {
	DBSnapshotIdentifier *string `json:"dbSnapshotIdentifier,omitempty"`
}

type DBSubnetGroup_SDK struct {
	DBSubnetGroupARN         *string   `json:"dbSubnetGroupARN,omitempty"`
	DBSubnetGroupDescription *string   `json:"dbSubnetGroupDescription,omitempty"`
	DBSubnetGroupName        *string   `json:"dbSubnetGroupName,omitempty"`
	SubnetGroupStatus        *string   `json:"subnetGroupStatus,omitempty"`
	Subnets                  []*Subnet `json:"subnets,omitempty"`
	VPCID                    *string   `json:"vpcID,omitempty"`
}

type DescribeDBLogFilesDetails struct {
	LogFileName *string `json:"logFileName,omitempty"`
}

type DomainMembership struct {
	Domain      *string `json:"domain,omitempty"`
	FQDN        *string `json:"fQDN,omitempty"`
	IAMRoleName *string `json:"iamRoleName,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type EC2SecurityGroup struct {
	EC2SecurityGroupID      *string `json:"ec2SecurityGroupID,omitempty"`
	EC2SecurityGroupName    *string `json:"ec2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerID *string `json:"ec2SecurityGroupOwnerID,omitempty"`
	Status                  *string `json:"status,omitempty"`
}

type Endpoint struct {
	Address      *string `json:"address,omitempty"`
	HostedZoneID *string `json:"hostedZoneID,omitempty"`
}

type EngineDefaults struct {
	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`
	Marker                 *string `json:"marker,omitempty"`
}

type Event struct {
	Message          *string `json:"message,omitempty"`
	SourceARN        *string `json:"sourceARN,omitempty"`
	SourceIdentifier *string `json:"sourceIdentifier,omitempty"`
}

type EventCategoriesMap struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type EventSubscription struct {
	CustSubscriptionID       *string `json:"custSubscriptionID,omitempty"`
	CustomerAWSID            *string `json:"customerAWSID,omitempty"`
	EventSubscriptionARN     *string `json:"eventSubscriptionARN,omitempty"`
	SnsTopicARN              *string `json:"snsTopicARN,omitempty"`
	SourceType               *string `json:"sourceType,omitempty"`
	Status                   *string `json:"status,omitempty"`
	SubscriptionCreationTime *string `json:"subscriptionCreationTime,omitempty"`
}

type ExportTask struct {
	ExportTaskIdentifier *string `json:"exportTaskIdentifier,omitempty"`
	FailureCause         *string `json:"failureCause,omitempty"`
	IAMRoleARN           *string `json:"iamRoleARN,omitempty"`
	KMSKeyID             *string `json:"kmsKeyID,omitempty"`
	S3Bucket             *string `json:"s3Bucket,omitempty"`
	S3Prefix             *string `json:"s3Prefix,omitempty"`
	SourceARN            *string `json:"sourceARN,omitempty"`
	Status               *string `json:"status,omitempty"`
	WarningMessage       *string `json:"warningMessage,omitempty"`
}

type Filter struct {
	Name   *string   `json:"name,omitempty"`
	Values []*string `json:"values,omitempty"`
}

type GlobalCluster struct {
	DatabaseName            *string `json:"databaseName,omitempty"`
	Engine                  *string `json:"engine,omitempty"`
	EngineVersion           *string `json:"engineVersion,omitempty"`
	GlobalClusterARN        *string `json:"globalClusterARN,omitempty"`
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty"`
	GlobalClusterResourceID *string `json:"globalClusterResourceID,omitempty"`
	Status                  *string `json:"status,omitempty"`
}

type GlobalClusterMember struct {
	DBClusterARN *string `json:"dbClusterARN,omitempty"`
}

type IPRange struct {
	CIDRIP *string `json:"cidrIP,omitempty"`
	Status *string `json:"status,omitempty"`
}

type InstallationMedia struct {
	CustomAvailabilityZoneID    *string `json:"customAvailabilityZoneID,omitempty"`
	Engine                      *string `json:"engine,omitempty"`
	EngineInstallationMediaPath *string `json:"engineInstallationMediaPath,omitempty"`
	EngineVersion               *string `json:"engineVersion,omitempty"`
	InstallationMediaID         *string `json:"installationMediaID,omitempty"`
	OSInstallationMediaPath     *string `json:"oSInstallationMediaPath,omitempty"`
	Status                      *string `json:"status,omitempty"`
}

type InstallationMediaFailureCause struct {
	Message *string `json:"message,omitempty"`
}

type MinimumEngineVersionPerAllowedValue struct {
	AllowedValue         *string `json:"allowedValue,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
}

type Option struct {
	OptionDescription *string `json:"optionDescription,omitempty"`
	OptionName        *string `json:"optionName,omitempty"`
	OptionVersion     *string `json:"optionVersion,omitempty"`
	Port              *int64  `json:"port,omitempty"`
}

type OptionConfiguration struct {
	OptionName    *string `json:"optionName,omitempty"`
	OptionVersion *string `json:"optionVersion,omitempty"`
	Port          *int64  `json:"port,omitempty"`
}

type OptionGroup struct {
	EngineName             *string `json:"engineName,omitempty"`
	MajorEngineVersion     *string `json:"majorEngineVersion,omitempty"`
	OptionGroupARN         *string `json:"optionGroupARN,omitempty"`
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty"`
	OptionGroupName        *string `json:"optionGroupName,omitempty"`
	VPCID                  *string `json:"vpcID,omitempty"`
}

type OptionGroupMembership struct {
	OptionGroupName *string `json:"optionGroupName,omitempty"`
	Status          *string `json:"status,omitempty"`
}

type OptionGroupOption struct {
	DefaultPort                       *int64  `json:"defaultPort,omitempty"`
	Description                       *string `json:"description,omitempty"`
	EngineName                        *string `json:"engineName,omitempty"`
	MajorEngineVersion                *string `json:"majorEngineVersion,omitempty"`
	MinimumRequiredMinorEngineVersion *string `json:"minimumRequiredMinorEngineVersion,omitempty"`
	Name                              *string `json:"name,omitempty"`
}

type OptionGroupOptionSetting struct {
	AllowedValues      *string `json:"allowedValues,omitempty"`
	ApplyType          *string `json:"applyType,omitempty"`
	DefaultValue       *string `json:"defaultValue,omitempty"`
	SettingDescription *string `json:"settingDescription,omitempty"`
	SettingName        *string `json:"settingName,omitempty"`
}

type OptionSetting struct {
	AllowedValues *string `json:"allowedValues,omitempty"`
	ApplyType     *string `json:"applyType,omitempty"`
	DataType      *string `json:"dataType,omitempty"`
	DefaultValue  *string `json:"defaultValue,omitempty"`
	Description   *string `json:"description,omitempty"`
	Name          *string `json:"name,omitempty"`
	Value         *string `json:"value,omitempty"`
}

type OptionVersion struct {
	Version *string `json:"version,omitempty"`
}

type OrderableDBInstanceOption struct {
	AvailabilityZoneGroup *string `json:"availabilityZoneGroup,omitempty"`
	DBInstanceClass       *string `json:"dbInstanceClass,omitempty"`
	Engine                *string `json:"engine,omitempty"`
	EngineVersion         *string `json:"engineVersion,omitempty"`
	LicenseModel          *string `json:"licenseModel,omitempty"`
	MaxIOPSPerDBInstance  *int64  `json:"maxIOPSPerDBInstance,omitempty"`
	MaxStorageSize        *int64  `json:"maxStorageSize,omitempty"`
	MinIOPSPerDBInstance  *int64  `json:"minIOPSPerDBInstance,omitempty"`
	MinStorageSize        *int64  `json:"minStorageSize,omitempty"`
	StorageType           *string `json:"storageType,omitempty"`
}

type Outpost struct {
	ARN *string `json:"arn,omitempty"`
}

type Parameter struct {
	AllowedValues        *string `json:"allowedValues,omitempty"`
	ApplyType            *string `json:"applyType,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	Description          *string `json:"description,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	ParameterName        *string `json:"parameterName,omitempty"`
	ParameterValue       *string `json:"parameterValue,omitempty"`
	Source               *string `json:"source,omitempty"`
}

type PendingMaintenanceAction struct {
	Action      *string `json:"action,omitempty"`
	Description *string `json:"description,omitempty"`
	OptInStatus *string `json:"optInStatus,omitempty"`
}

type PendingModifiedValues struct {
	AllocatedStorage        *int64  `json:"allocatedStorage,omitempty"`
	BackupRetentionPeriod   *int64  `json:"backupRetentionPeriod,omitempty"`
	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`
	DBInstanceClass         *string `json:"dbInstanceClass,omitempty"`
	DBInstanceIdentifier    *string `json:"dbInstanceIdentifier,omitempty"`
	DBSubnetGroupName       *string `json:"dbSubnetGroupName,omitempty"`
	EngineVersion           *string `json:"engineVersion,omitempty"`
	IOPS                    *int64  `json:"iops,omitempty"`
	LicenseModel            *string `json:"licenseModel,omitempty"`
	MasterUserPassword      *string `json:"masterUserPassword,omitempty"`
	Port                    *int64  `json:"port,omitempty"`
	StorageType             *string `json:"storageType,omitempty"`
}

type ProcessorFeature struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Range struct {
	Step *int64 `json:"step,omitempty"`
}

type RecurringCharge struct {
	RecurringChargeFrequency *string `json:"recurringChargeFrequency,omitempty"`
}

type ReservedDBInstance struct {
	CurrencyCode                  *string `json:"currencyCode,omitempty"`
	DBInstanceClass               *string `json:"dbInstanceClass,omitempty"`
	LeaseID                       *string `json:"leaseID,omitempty"`
	OfferingType                  *string `json:"offeringType,omitempty"`
	ProductDescription            *string `json:"productDescription,omitempty"`
	ReservedDBInstanceARN         *string `json:"reservedDBInstanceARN,omitempty"`
	ReservedDBInstanceID          *string `json:"reservedDBInstanceID,omitempty"`
	ReservedDBInstancesOfferingID *string `json:"reservedDBInstancesOfferingID,omitempty"`
	State                         *string `json:"state,omitempty"`
}

type ReservedDBInstancesOffering struct {
	CurrencyCode                  *string `json:"currencyCode,omitempty"`
	DBInstanceClass               *string `json:"dbInstanceClass,omitempty"`
	OfferingType                  *string `json:"offeringType,omitempty"`
	ProductDescription            *string `json:"productDescription,omitempty"`
	ReservedDBInstancesOfferingID *string `json:"reservedDBInstancesOfferingID,omitempty"`
}

type ResourcePendingMaintenanceActions struct {
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty"`
}

type ScalingConfiguration struct {
	MaxCapacity           *int64  `json:"maxCapacity,omitempty"`
	MinCapacity           *int64  `json:"minCapacity,omitempty"`
	SecondsUntilAutoPause *int64  `json:"secondsUntilAutoPause,omitempty"`
	TimeoutAction         *string `json:"timeoutAction,omitempty"`
}

type ScalingConfigurationInfo struct {
	MaxCapacity           *int64  `json:"maxCapacity,omitempty"`
	MinCapacity           *int64  `json:"minCapacity,omitempty"`
	SecondsUntilAutoPause *int64  `json:"secondsUntilAutoPause,omitempty"`
	TimeoutAction         *string `json:"timeoutAction,omitempty"`
}

type SourceRegion struct {
	Endpoint   *string `json:"endpoint,omitempty"`
	RegionName *string `json:"regionName,omitempty"`
	Status     *string `json:"status,omitempty"`
}

type Subnet struct {
	SubnetAvailabilityZone *AvailabilityZone `json:"subnetAvailabilityZone,omitempty"`
	SubnetIdentifier       *string           `json:"subnetIdentifier,omitempty"`
	SubnetOutpost          *Outpost          `json:"subnetOutpost,omitempty"`
	SubnetStatus           *string           `json:"subnetStatus,omitempty"`
}

type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TargetHealth struct {
	Description *string `json:"description,omitempty"`
}

type Timezone struct {
	TimezoneName *string `json:"timezoneName,omitempty"`
}

type UpgradeTarget struct {
	Description   *string `json:"description,omitempty"`
	Engine        *string `json:"engine,omitempty"`
	EngineVersion *string `json:"engineVersion,omitempty"`
}

type UserAuthConfig struct {
	Description *string `json:"description,omitempty"`
	SecretARN   *string `json:"secretARN,omitempty"`
	UserName    *string `json:"userName,omitempty"`
}

type UserAuthConfigInfo struct {
	Description *string `json:"description,omitempty"`
	SecretARN   *string `json:"secretARN,omitempty"`
	UserName    *string `json:"userName,omitempty"`
}

type VPCSecurityGroupMembership struct {
	Status             *string `json:"status,omitempty"`
	VPCSecurityGroupID *string `json:"vpcSecurityGroupID,omitempty"`
}

type VPNDetails struct {
	VPNGatewayIP          *string `json:"vpnGatewayIP,omitempty"`
	VPNID                 *string `json:"vpnID,omitempty"`
	VPNName               *string `json:"vpnName,omitempty"`
	VPNState              *string `json:"vpnState,omitempty"`
	VPNTunnelOriginatorIP *string `json:"vpnTunnelOriginatorIP,omitempty"`
}

type ValidStorageOptions struct {
	StorageType *string `json:"storageType,omitempty"`
}
