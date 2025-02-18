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

package db_cluster

import (
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if !ackcompare.SliceStringPEqual(a.ko.Spec.AvailabilityZones, b.ko.Spec.AvailabilityZones) {
		delta.Add("Spec.AvailabilityZones", a.ko.Spec.AvailabilityZones, b.ko.Spec.AvailabilityZones)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BacktrackWindow, b.ko.Spec.BacktrackWindow) {
		delta.Add("Spec.BacktrackWindow", a.ko.Spec.BacktrackWindow, b.ko.Spec.BacktrackWindow)
	} else if a.ko.Spec.BacktrackWindow != nil && b.ko.Spec.BacktrackWindow != nil {
		if *a.ko.Spec.BacktrackWindow != *b.ko.Spec.BacktrackWindow {
			delta.Add("Spec.BacktrackWindow", a.ko.Spec.BacktrackWindow, b.ko.Spec.BacktrackWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod) {
		delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
	} else if a.ko.Spec.BackupRetentionPeriod != nil && b.ko.Spec.BackupRetentionPeriod != nil {
		if *a.ko.Spec.BackupRetentionPeriod != *b.ko.Spec.BackupRetentionPeriod {
			delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName) {
		delta.Add("Spec.CharacterSetName", a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName)
	} else if a.ko.Spec.CharacterSetName != nil && b.ko.Spec.CharacterSetName != nil {
		if *a.ko.Spec.CharacterSetName != *b.ko.Spec.CharacterSetName {
			delta.Add("Spec.CharacterSetName", a.ko.Spec.CharacterSetName, b.ko.Spec.CharacterSetName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot) {
		delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
	} else if a.ko.Spec.CopyTagsToSnapshot != nil && b.ko.Spec.CopyTagsToSnapshot != nil {
		if *a.ko.Spec.CopyTagsToSnapshot != *b.ko.Spec.CopyTagsToSnapshot {
			delta.Add("Spec.CopyTagsToSnapshot", a.ko.Spec.CopyTagsToSnapshot, b.ko.Spec.CopyTagsToSnapshot)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier) {
		delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
	} else if a.ko.Spec.DBClusterIdentifier != nil && b.ko.Spec.DBClusterIdentifier != nil {
		if *a.ko.Spec.DBClusterIdentifier != *b.ko.Spec.DBClusterIdentifier {
			delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName) {
		delta.Add("Spec.DBClusterParameterGroupName", a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName)
	} else if a.ko.Spec.DBClusterParameterGroupName != nil && b.ko.Spec.DBClusterParameterGroupName != nil {
		if *a.ko.Spec.DBClusterParameterGroupName != *b.ko.Spec.DBClusterParameterGroupName {
			delta.Add("Spec.DBClusterParameterGroupName", a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName) {
		delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
	} else if a.ko.Spec.DBSubnetGroupName != nil && b.ko.Spec.DBSubnetGroupName != nil {
		if *a.ko.Spec.DBSubnetGroupName != *b.ko.Spec.DBSubnetGroupName {
			delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DatabaseName, b.ko.Spec.DatabaseName) {
		delta.Add("Spec.DatabaseName", a.ko.Spec.DatabaseName, b.ko.Spec.DatabaseName)
	} else if a.ko.Spec.DatabaseName != nil && b.ko.Spec.DatabaseName != nil {
		if *a.ko.Spec.DatabaseName != *b.ko.Spec.DatabaseName {
			delta.Add("Spec.DatabaseName", a.ko.Spec.DatabaseName, b.ko.Spec.DatabaseName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection) {
		delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
	} else if a.ko.Spec.DeletionProtection != nil && b.ko.Spec.DeletionProtection != nil {
		if *a.ko.Spec.DeletionProtection != *b.ko.Spec.DeletionProtection {
			delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion) {
		delta.Add("Spec.DestinationRegion", a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion)
	} else if a.ko.Spec.DestinationRegion != nil && b.ko.Spec.DestinationRegion != nil {
		if *a.ko.Spec.DestinationRegion != *b.ko.Spec.DestinationRegion {
			delta.Add("Spec.DestinationRegion", a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Domain, b.ko.Spec.Domain) {
		delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
	} else if a.ko.Spec.Domain != nil && b.ko.Spec.Domain != nil {
		if *a.ko.Spec.Domain != *b.ko.Spec.Domain {
			delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName) {
		delta.Add("Spec.DomainIAMRoleName", a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName)
	} else if a.ko.Spec.DomainIAMRoleName != nil && b.ko.Spec.DomainIAMRoleName != nil {
		if *a.ko.Spec.DomainIAMRoleName != *b.ko.Spec.DomainIAMRoleName {
			delta.Add("Spec.DomainIAMRoleName", a.ko.Spec.DomainIAMRoleName, b.ko.Spec.DomainIAMRoleName)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports) {
		delta.Add("Spec.EnableCloudwatchLogsExports", a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableGlobalWriteForwarding, b.ko.Spec.EnableGlobalWriteForwarding) {
		delta.Add("Spec.EnableGlobalWriteForwarding", a.ko.Spec.EnableGlobalWriteForwarding, b.ko.Spec.EnableGlobalWriteForwarding)
	} else if a.ko.Spec.EnableGlobalWriteForwarding != nil && b.ko.Spec.EnableGlobalWriteForwarding != nil {
		if *a.ko.Spec.EnableGlobalWriteForwarding != *b.ko.Spec.EnableGlobalWriteForwarding {
			delta.Add("Spec.EnableGlobalWriteForwarding", a.ko.Spec.EnableGlobalWriteForwarding, b.ko.Spec.EnableGlobalWriteForwarding)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableHTTPEndpoint, b.ko.Spec.EnableHTTPEndpoint) {
		delta.Add("Spec.EnableHTTPEndpoint", a.ko.Spec.EnableHTTPEndpoint, b.ko.Spec.EnableHTTPEndpoint)
	} else if a.ko.Spec.EnableHTTPEndpoint != nil && b.ko.Spec.EnableHTTPEndpoint != nil {
		if *a.ko.Spec.EnableHTTPEndpoint != *b.ko.Spec.EnableHTTPEndpoint {
			delta.Add("Spec.EnableHTTPEndpoint", a.ko.Spec.EnableHTTPEndpoint, b.ko.Spec.EnableHTTPEndpoint)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication) {
		delta.Add("Spec.EnableIAMDatabaseAuthentication", a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication)
	} else if a.ko.Spec.EnableIAMDatabaseAuthentication != nil && b.ko.Spec.EnableIAMDatabaseAuthentication != nil {
		if *a.ko.Spec.EnableIAMDatabaseAuthentication != *b.ko.Spec.EnableIAMDatabaseAuthentication {
			delta.Add("Spec.EnableIAMDatabaseAuthentication", a.ko.Spec.EnableIAMDatabaseAuthentication, b.ko.Spec.EnableIAMDatabaseAuthentication)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineMode, b.ko.Spec.EngineMode) {
		delta.Add("Spec.EngineMode", a.ko.Spec.EngineMode, b.ko.Spec.EngineMode)
	} else if a.ko.Spec.EngineMode != nil && b.ko.Spec.EngineMode != nil {
		if *a.ko.Spec.EngineMode != *b.ko.Spec.EngineMode {
			delta.Add("Spec.EngineMode", a.ko.Spec.EngineMode, b.ko.Spec.EngineMode)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion) {
		delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
	} else if a.ko.Spec.EngineVersion != nil && b.ko.Spec.EngineVersion != nil {
		if *a.ko.Spec.EngineVersion != *b.ko.Spec.EngineVersion {
			delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier) {
		delta.Add("Spec.GlobalClusterIdentifier", a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier)
	} else if a.ko.Spec.GlobalClusterIdentifier != nil && b.ko.Spec.GlobalClusterIdentifier != nil {
		if *a.ko.Spec.GlobalClusterIdentifier != *b.ko.Spec.GlobalClusterIdentifier {
			delta.Add("Spec.GlobalClusterIdentifier", a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword) {
		delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
	} else if a.ko.Spec.MasterUserPassword != nil && b.ko.Spec.MasterUserPassword != nil {
		if *a.ko.Spec.MasterUserPassword != *b.ko.Spec.MasterUserPassword {
			delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername) {
		delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
	} else if a.ko.Spec.MasterUsername != nil && b.ko.Spec.MasterUsername != nil {
		if *a.ko.Spec.MasterUsername != *b.ko.Spec.MasterUsername {
			delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName) {
		delta.Add("Spec.OptionGroupName", a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName)
	} else if a.ko.Spec.OptionGroupName != nil && b.ko.Spec.OptionGroupName != nil {
		if *a.ko.Spec.OptionGroupName != *b.ko.Spec.OptionGroupName {
			delta.Add("Spec.OptionGroupName", a.ko.Spec.OptionGroupName, b.ko.Spec.OptionGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Port, b.ko.Spec.Port) {
		delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
	} else if a.ko.Spec.Port != nil && b.ko.Spec.Port != nil {
		if *a.ko.Spec.Port != *b.ko.Spec.Port {
			delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL) {
		delta.Add("Spec.PreSignedURL", a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL)
	} else if a.ko.Spec.PreSignedURL != nil && b.ko.Spec.PreSignedURL != nil {
		if *a.ko.Spec.PreSignedURL != *b.ko.Spec.PreSignedURL {
			delta.Add("Spec.PreSignedURL", a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow) {
		delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
	} else if a.ko.Spec.PreferredBackupWindow != nil && b.ko.Spec.PreferredBackupWindow != nil {
		if *a.ko.Spec.PreferredBackupWindow != *b.ko.Spec.PreferredBackupWindow {
			delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow) {
		delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
	} else if a.ko.Spec.PreferredMaintenanceWindow != nil && b.ko.Spec.PreferredMaintenanceWindow != nil {
		if *a.ko.Spec.PreferredMaintenanceWindow != *b.ko.Spec.PreferredMaintenanceWindow {
			delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ReplicationSourceIdentifier, b.ko.Spec.ReplicationSourceIdentifier) {
		delta.Add("Spec.ReplicationSourceIdentifier", a.ko.Spec.ReplicationSourceIdentifier, b.ko.Spec.ReplicationSourceIdentifier)
	} else if a.ko.Spec.ReplicationSourceIdentifier != nil && b.ko.Spec.ReplicationSourceIdentifier != nil {
		if *a.ko.Spec.ReplicationSourceIdentifier != *b.ko.Spec.ReplicationSourceIdentifier {
			delta.Add("Spec.ReplicationSourceIdentifier", a.ko.Spec.ReplicationSourceIdentifier, b.ko.Spec.ReplicationSourceIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration, b.ko.Spec.ScalingConfiguration) {
		delta.Add("Spec.ScalingConfiguration", a.ko.Spec.ScalingConfiguration, b.ko.Spec.ScalingConfiguration)
	} else if a.ko.Spec.ScalingConfiguration != nil && b.ko.Spec.ScalingConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration.AutoPause, b.ko.Spec.ScalingConfiguration.AutoPause) {
			delta.Add("Spec.ScalingConfiguration.AutoPause", a.ko.Spec.ScalingConfiguration.AutoPause, b.ko.Spec.ScalingConfiguration.AutoPause)
		} else if a.ko.Spec.ScalingConfiguration.AutoPause != nil && b.ko.Spec.ScalingConfiguration.AutoPause != nil {
			if *a.ko.Spec.ScalingConfiguration.AutoPause != *b.ko.Spec.ScalingConfiguration.AutoPause {
				delta.Add("Spec.ScalingConfiguration.AutoPause", a.ko.Spec.ScalingConfiguration.AutoPause, b.ko.Spec.ScalingConfiguration.AutoPause)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration.MaxCapacity, b.ko.Spec.ScalingConfiguration.MaxCapacity) {
			delta.Add("Spec.ScalingConfiguration.MaxCapacity", a.ko.Spec.ScalingConfiguration.MaxCapacity, b.ko.Spec.ScalingConfiguration.MaxCapacity)
		} else if a.ko.Spec.ScalingConfiguration.MaxCapacity != nil && b.ko.Spec.ScalingConfiguration.MaxCapacity != nil {
			if *a.ko.Spec.ScalingConfiguration.MaxCapacity != *b.ko.Spec.ScalingConfiguration.MaxCapacity {
				delta.Add("Spec.ScalingConfiguration.MaxCapacity", a.ko.Spec.ScalingConfiguration.MaxCapacity, b.ko.Spec.ScalingConfiguration.MaxCapacity)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration.MinCapacity, b.ko.Spec.ScalingConfiguration.MinCapacity) {
			delta.Add("Spec.ScalingConfiguration.MinCapacity", a.ko.Spec.ScalingConfiguration.MinCapacity, b.ko.Spec.ScalingConfiguration.MinCapacity)
		} else if a.ko.Spec.ScalingConfiguration.MinCapacity != nil && b.ko.Spec.ScalingConfiguration.MinCapacity != nil {
			if *a.ko.Spec.ScalingConfiguration.MinCapacity != *b.ko.Spec.ScalingConfiguration.MinCapacity {
				delta.Add("Spec.ScalingConfiguration.MinCapacity", a.ko.Spec.ScalingConfiguration.MinCapacity, b.ko.Spec.ScalingConfiguration.MinCapacity)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause, b.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause) {
			delta.Add("Spec.ScalingConfiguration.SecondsUntilAutoPause", a.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause, b.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause)
		} else if a.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause != nil && b.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause != nil {
			if *a.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause != *b.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause {
				delta.Add("Spec.ScalingConfiguration.SecondsUntilAutoPause", a.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause, b.ko.Spec.ScalingConfiguration.SecondsUntilAutoPause)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfiguration.TimeoutAction, b.ko.Spec.ScalingConfiguration.TimeoutAction) {
			delta.Add("Spec.ScalingConfiguration.TimeoutAction", a.ko.Spec.ScalingConfiguration.TimeoutAction, b.ko.Spec.ScalingConfiguration.TimeoutAction)
		} else if a.ko.Spec.ScalingConfiguration.TimeoutAction != nil && b.ko.Spec.ScalingConfiguration.TimeoutAction != nil {
			if *a.ko.Spec.ScalingConfiguration.TimeoutAction != *b.ko.Spec.ScalingConfiguration.TimeoutAction {
				delta.Add("Spec.ScalingConfiguration.TimeoutAction", a.ko.Spec.ScalingConfiguration.TimeoutAction, b.ko.Spec.ScalingConfiguration.TimeoutAction)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion) {
		delta.Add("Spec.SourceRegion", a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion)
	} else if a.ko.Spec.SourceRegion != nil && b.ko.Spec.SourceRegion != nil {
		if *a.ko.Spec.SourceRegion != *b.ko.Spec.SourceRegion {
			delta.Add("Spec.SourceRegion", a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted) {
		delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
	} else if a.ko.Spec.StorageEncrypted != nil && b.ko.Spec.StorageEncrypted != nil {
		if *a.ko.Spec.StorageEncrypted != *b.ko.Spec.StorageEncrypted {
			delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs) {
		delta.Add("Spec.VPCSecurityGroupIDs", a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs)
	}

	return delta
}
