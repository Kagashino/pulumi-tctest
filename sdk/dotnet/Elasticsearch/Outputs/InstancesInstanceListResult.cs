// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceListResult
    {
        public readonly string AvailabilityZone;
        public readonly int BasicSecurityType;
        public readonly string ChargeType;
        public readonly string CreateTime;
        public readonly int DeployMode;
        public readonly string ElasticsearchDomain;
        public readonly int ElasticsearchPort;
        public readonly string ElasticsearchVip;
        public readonly string InstanceId;
        public readonly string InstanceName;
        public readonly string KibanaUrl;
        public readonly string LicenseType;
        public readonly ImmutableArray<Outputs.InstancesInstanceListMultiZoneInfoResult> MultiZoneInfos;
        public readonly ImmutableArray<Outputs.InstancesInstanceListNodeInfoListResult> NodeInfoLists;
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string Version;
        public readonly string VpcId;

        [OutputConstructor]
        private InstancesInstanceListResult(
            string availabilityZone,

            int basicSecurityType,

            string chargeType,

            string createTime,

            int deployMode,

            string elasticsearchDomain,

            int elasticsearchPort,

            string elasticsearchVip,

            string instanceId,

            string instanceName,

            string kibanaUrl,

            string licenseType,

            ImmutableArray<Outputs.InstancesInstanceListMultiZoneInfoResult> multiZoneInfos,

            ImmutableArray<Outputs.InstancesInstanceListNodeInfoListResult> nodeInfoLists,

            string subnetId,

            ImmutableDictionary<string, object> tags,

            string version,

            string vpcId)
        {
            AvailabilityZone = availabilityZone;
            BasicSecurityType = basicSecurityType;
            ChargeType = chargeType;
            CreateTime = createTime;
            DeployMode = deployMode;
            ElasticsearchDomain = elasticsearchDomain;
            ElasticsearchPort = elasticsearchPort;
            ElasticsearchVip = elasticsearchVip;
            InstanceId = instanceId;
            InstanceName = instanceName;
            KibanaUrl = kibanaUrl;
            LicenseType = licenseType;
            MultiZoneInfos = multiZoneInfos;
            NodeInfoLists = nodeInfoLists;
            SubnetId = subnetId;
            Tags = tags;
            Version = version;
            VpcId = vpcId;
        }
    }
}
