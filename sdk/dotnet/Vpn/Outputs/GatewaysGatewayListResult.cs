// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpn.Outputs
{

    [OutputType]
    public sealed class GatewaysGatewayListResult
    {
        public readonly int Bandwidth;
        public readonly string ChargeType;
        public readonly string CreateTime;
        public readonly string ExpiredTime;
        public readonly string Id;
        public readonly bool IsAddressBlocked;
        public readonly string Name;
        public readonly string NewPurchasePlan;
        public readonly string PrepaidRenewFlag;
        public readonly string PublicIpAddress;
        public readonly string RestrictState;
        public readonly string State;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string Type;
        public readonly string VpcId;
        public readonly string Zone;

        [OutputConstructor]
        private GatewaysGatewayListResult(
            int bandwidth,

            string chargeType,

            string createTime,

            string expiredTime,

            string id,

            bool isAddressBlocked,

            string name,

            string newPurchasePlan,

            string prepaidRenewFlag,

            string publicIpAddress,

            string restrictState,

            string state,

            ImmutableDictionary<string, object> tags,

            string type,

            string vpcId,

            string zone)
        {
            Bandwidth = bandwidth;
            ChargeType = chargeType;
            CreateTime = createTime;
            ExpiredTime = expiredTime;
            Id = id;
            IsAddressBlocked = isAddressBlocked;
            Name = name;
            NewPurchasePlan = newPurchasePlan;
            PrepaidRenewFlag = prepaidRenewFlag;
            PublicIpAddress = publicIpAddress;
            RestrictState = restrictState;
            State = state;
            Tags = tags;
            Type = type;
            VpcId = vpcId;
            Zone = zone;
        }
    }
}
