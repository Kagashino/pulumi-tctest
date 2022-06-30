// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class RedirectionsRedirectionListResult
    {
        public readonly string ClbId;
        public readonly string SourceListenerId;
        public readonly string SourceRuleId;
        public readonly string TargetListenerId;
        public readonly string TargetRuleId;

        [OutputConstructor]
        private RedirectionsRedirectionListResult(
            string clbId,

            string sourceListenerId,

            string sourceRuleId,

            string targetListenerId,

            string targetRuleId)
        {
            ClbId = clbId;
            SourceListenerId = sourceListenerId;
            SourceRuleId = sourceRuleId;
            TargetListenerId = targetListenerId;
            TargetRuleId = targetRuleId;
        }
    }
}
