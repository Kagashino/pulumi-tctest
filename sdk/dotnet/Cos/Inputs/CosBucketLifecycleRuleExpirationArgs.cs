// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class CosBucketLifecycleRuleExpirationArgs : Pulumi.ResourceArgs
    {
        [Input("date")]
        public Input<string>? Date { get; set; }

        [Input("days")]
        public Input<int>? Days { get; set; }

        [Input("deleteMarker")]
        public Input<bool>? DeleteMarker { get; set; }

        public CosBucketLifecycleRuleExpirationArgs()
        {
        }
    }
}
