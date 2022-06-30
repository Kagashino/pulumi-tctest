// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dcx.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceListResult
    {
        public readonly int Bandwidth;
        public readonly int BgpAsn;
        public readonly string BgpAuthKey;
        public readonly string CreateTime;
        public readonly string CustomerAddress;
        public readonly string DcId;
        public readonly string DcgId;
        public readonly string DcxId;
        public readonly string Name;
        public readonly string NetworkRegion;
        public readonly string NetworkType;
        public readonly ImmutableArray<string> RouteFilterPrefixes;
        public readonly string RouteType;
        public readonly string State;
        public readonly string TencentAddress;
        public readonly int Vlan;
        public readonly string VpcId;

        [OutputConstructor]
        private InstancesInstanceListResult(
            int bandwidth,

            int bgpAsn,

            string bgpAuthKey,

            string createTime,

            string customerAddress,

            string dcId,

            string dcgId,

            string dcxId,

            string name,

            string networkRegion,

            string networkType,

            ImmutableArray<string> routeFilterPrefixes,

            string routeType,

            string state,

            string tencentAddress,

            int vlan,

            string vpcId)
        {
            Bandwidth = bandwidth;
            BgpAsn = bgpAsn;
            BgpAuthKey = bgpAuthKey;
            CreateTime = createTime;
            CustomerAddress = customerAddress;
            DcId = dcId;
            DcgId = dcgId;
            DcxId = dcxId;
            Name = name;
            NetworkRegion = networkRegion;
            NetworkType = networkType;
            RouteFilterPrefixes = routeFilterPrefixes;
            RouteType = routeType;
            State = state;
            TencentAddress = tencentAddress;
            Vlan = vlan;
            VpcId = vpcId;
        }
    }
}
