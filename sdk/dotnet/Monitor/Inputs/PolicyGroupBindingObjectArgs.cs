// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Inputs
{

    public sealed class PolicyGroupBindingObjectArgs : Pulumi.ResourceArgs
    {
        [Input("dimensionsJson")]
        public Input<string>? DimensionsJson { get; set; }

        [Input("isShielded")]
        public Input<int>? IsShielded { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        public PolicyGroupBindingObjectArgs()
        {
        }
    }
}
