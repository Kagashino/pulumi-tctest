// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Inputs
{

    public sealed class CCHttpsPolicyRuleListGetArgs : Pulumi.ResourceArgs
    {
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        [Input("skey", required: true)]
        public Input<string> Skey { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CCHttpsPolicyRuleListGetArgs()
        {
        }
    }
}
