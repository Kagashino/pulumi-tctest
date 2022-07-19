// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cloud.Outputs
{

    [OutputType]
    public sealed class APIGatewayAPIKeysListResult
    {
        public readonly string AccessKeySecret;
        public readonly string ApiKeyId;
        public readonly string CreateTime;
        public readonly string ModifyTime;
        public readonly string Status;

        [OutputConstructor]
        private APIGatewayAPIKeysListResult(
            string accessKeySecret,

            string apiKeyId,

            string createTime,

            string modifyTime,

            string status)
        {
            AccessKeySecret = accessKeySecret;
            ApiKeyId = apiKeyId;
            CreateTime = createTime;
            ModifyTime = modifyTime;
            Status = status;
        }
    }
}
