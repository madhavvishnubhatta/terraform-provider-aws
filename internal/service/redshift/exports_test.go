// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshift

// Exports for use in tests only.
var (
	ResourceAuthenticationProfile        = resourceAuthenticationProfile
	ResourceCluster                      = resourceCluster
	ResourceClusterIAMRoles              = resourceClusterIAMRoles
	ResourceClusterSnapshot              = resourceClusterSnapshot
	ResourceDataShareAuthorization       = newDataShareAuthorizationResource
	ResourceDataShareConsumerAssociation = newDataShareConsumerAssociationResource
	ResourceEndpointAccess               = resourceEndpointAccess
	ResourceEndpointAuthorization        = resourceEndpointAuthorization
	ResourceEventSubscription            = resourceEventSubscription
	ResourceHSMClientCertificate         = resourceHSMClientCertificate
	ResourceHSMConfiguration             = resourceHSMConfiguration
	ResourceIntegration                  = newIntegrationResource
	ResourceLogging                      = newLoggingResource
	ResourceParameterGroup               = resourceParameterGroup
	ResourcePartner                      = resourcePartner
	ResourceResourcePolicy               = resourceResourcePolicy
	ResourceScheduledAction              = resourceScheduledAction
	ResourceSnapshotCopy                 = newSnapshotCopyResource
	ResourceSnapshotCopyGrant            = resourceSnapshotCopyGrant
	ResourceSnapshotSchedule             = resourceSnapshotSchedule
	ResourceSnapshotScheduleAssociation  = resourceSnapshotScheduleAssociation
	ResourceSubnetGroup                  = resourceSubnetGroup
	ResourceUsageLimit                   = resourceUsageLimit

	FindAuthenticationProfileByID               = findAuthenticationProfileByID
	FindClusterByID                             = findClusterByID
	FindClusterSnapshotByID                     = findClusterSnapshotByID
	FindDataShareAuthorizationByID              = findDataShareAuthorizationByID
	FindDataShareConsumerAssociationByID        = findDataShareConsumerAssociationByID
	FindEndpointAccessByName                    = findEndpointAccessByName
	FindEndpointAuthorizationByID               = findEndpointAuthorizationByID
	FindEventSubscriptionByName                 = findEventSubscriptionByName
	FindHSMClientCertificateByID                = findHSMClientCertificateByID
	FindHSMConfigurationByID                    = findHSMConfigurationByID
	FindIntegrationByARN                        = findIntegrationByARN
	FindLoggingByID                             = findLoggingByID
	FindParameterGroupByName                    = findParameterGroupByName
	FindPartnerByID                             = findPartnerByID
	FindResourcePolicyByARN                     = findResourcePolicyByARN
	FindScheduledActionByName                   = findScheduledActionByName
	FindSnapshotCopyByID                        = findSnapshotCopyByID
	FindSnapshotCopyGrantByName                 = findSnapshotCopyGrantByName
	FindSnapshotScheduleAssociationByTwoPartKey = findSnapshotScheduleAssociationByTwoPartKey
	FindSnapshotScheduleByID                    = findSnapshotScheduleByID
	FindSubnetGroupByName                       = findSubnetGroupByName
	FindUsageLimitByID                          = findUsageLimitByID
	WaitSnapshotScheduleAssociationCreated      = waitSnapshotScheduleAssociationCreated
)
