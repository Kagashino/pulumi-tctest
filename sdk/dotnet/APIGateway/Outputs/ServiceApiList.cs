// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.APIGateway.Outputs
{

    [OutputType]
    public sealed class ServiceApiList
    {
        public readonly string? ApiDesc;
        public readonly string? ApiId;
        public readonly string? ApiName;
        public readonly string? Method;
        public readonly string? Path;

        [OutputConstructor]
        private ServiceApiList(
            string? apiDesc,

            string? apiId,

            string? apiName,

            string? method,

            string? path)
        {
            ApiDesc = apiDesc;
            ApiId = apiId;
            ApiName = apiName;
            Method = method;
            Path = path;
        }
    }
}
