// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class DdosPoliciesListWatermarkKeyResult
    {
        public readonly string Content;
        public readonly string? CreateTime;
        public readonly string Id;
        public readonly bool OpenSwitch;

        [OutputConstructor]
        private DdosPoliciesListWatermarkKeyResult(
            string content,

            string? createTime,

            string id,

            bool openSwitch)
        {
            Content = content;
            CreateTime = createTime;
            Id = id;
            OpenSwitch = openSwitch;
        }
    }
}