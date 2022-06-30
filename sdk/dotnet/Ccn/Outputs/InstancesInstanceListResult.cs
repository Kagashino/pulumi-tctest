// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceListResult
    {
        public readonly ImmutableArray<Outputs.InstancesInstanceListAttachmentListResult> AttachmentLists;
        public readonly string BandwidthLimitType;
        public readonly string CcnId;
        public readonly string ChargeType;
        public readonly string CreateTime;
        public readonly string Description;
        public readonly string Name;
        public readonly string Qos;
        public readonly string State;

        [OutputConstructor]
        private InstancesInstanceListResult(
            ImmutableArray<Outputs.InstancesInstanceListAttachmentListResult> attachmentLists,

            string bandwidthLimitType,

            string ccnId,

            string chargeType,

            string createTime,

            string description,

            string name,

            string qos,

            string state)
        {
            AttachmentLists = attachmentLists;
            BandwidthLimitType = bandwidthLimitType;
            CcnId = ccnId;
            ChargeType = chargeType;
            CreateTime = createTime;
            Description = description;
            Name = name;
            Qos = qos;
            State = state;
        }
    }
}
