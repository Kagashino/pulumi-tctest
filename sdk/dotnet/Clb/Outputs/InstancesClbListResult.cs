// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Clb.Outputs
{

    [OutputType]
    public sealed class InstancesClbListResult
    {
        public readonly string AddressIpVersion;
        public readonly string ClbId;
        public readonly string ClbName;
        public readonly ImmutableArray<string> ClbVips;
        public readonly string CreateTime;
        public readonly int InternetBandwidthMaxOut;
        public readonly string InternetChargeType;
        public readonly bool LocalZone;
        public readonly string NetworkType;
        public readonly int ProjectId;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly int Status;
        public readonly string StatusTime;
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string TargetRegionInfoRegion;
        public readonly string TargetRegionInfoVpcId;
        public readonly string VipIsp;
        public readonly string VpcId;
        public readonly string Zone;
        public readonly int ZoneId;
        public readonly string ZoneName;
        public readonly string ZoneRegion;

        [OutputConstructor]
        private InstancesClbListResult(
            string addressIpVersion,

            string clbId,

            string clbName,

            ImmutableArray<string> clbVips,

            string createTime,

            int internetBandwidthMaxOut,

            string internetChargeType,

            bool localZone,

            string networkType,

            int projectId,

            ImmutableArray<string> securityGroups,

            int status,

            string statusTime,

            string subnetId,

            ImmutableDictionary<string, object> tags,

            string targetRegionInfoRegion,

            string targetRegionInfoVpcId,

            string vipIsp,

            string vpcId,

            string zone,

            int zoneId,

            string zoneName,

            string zoneRegion)
        {
            AddressIpVersion = addressIpVersion;
            ClbId = clbId;
            ClbName = clbName;
            ClbVips = clbVips;
            CreateTime = createTime;
            InternetBandwidthMaxOut = internetBandwidthMaxOut;
            InternetChargeType = internetChargeType;
            LocalZone = localZone;
            NetworkType = networkType;
            ProjectId = projectId;
            SecurityGroups = securityGroups;
            Status = status;
            StatusTime = statusTime;
            SubnetId = subnetId;
            Tags = tags;
            TargetRegionInfoRegion = targetRegionInfoRegion;
            TargetRegionInfoVpcId = targetRegionInfoVpcId;
            VipIsp = vipIsp;
            VpcId = vpcId;
            Zone = zone;
            ZoneId = zoneId;
            ZoneName = zoneName;
            ZoneRegion = zoneRegion;
        }
    }
}
