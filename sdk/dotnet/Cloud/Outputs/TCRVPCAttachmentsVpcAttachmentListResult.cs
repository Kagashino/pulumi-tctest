// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cloud.Outputs
{

    [OutputType]
    public sealed class TCRVPCAttachmentsVpcAttachmentListResult
    {
        public readonly string AccessIp;
        public readonly bool EnablePublicDomainDns;
        public readonly bool EnableVpcDomainDns;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string VpcId;

        [OutputConstructor]
        private TCRVPCAttachmentsVpcAttachmentListResult(
            string accessIp,

            bool enablePublicDomainDns,

            bool enableVpcDomainDns,

            string status,

            string subnetId,

            string vpcId)
        {
            AccessIp = accessIp;
            EnablePublicDomainDns = enablePublicDomainDns;
            EnableVpcDomainDns = enableVpcDomainDns;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
