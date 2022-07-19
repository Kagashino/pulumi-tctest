// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Ccn.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceListAttachmentListResult
    {
        public readonly string AttachedTime;
        public readonly ImmutableArray<string> CidrBlocks;
        public readonly string InstanceId;
        public readonly string InstanceRegion;
        public readonly string InstanceType;
        public readonly string State;

        [OutputConstructor]
        private InstancesInstanceListAttachmentListResult(
            string attachedTime,

            ImmutableArray<string> cidrBlocks,

            string instanceId,

            string instanceRegion,

            string instanceType,

            string state)
        {
            AttachedTime = attachedTime;
            CidrBlocks = cidrBlocks;
            InstanceId = instanceId;
            InstanceRegion = instanceRegion;
            InstanceType = instanceType;
            State = state;
        }
    }
}
