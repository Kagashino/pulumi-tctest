// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Inputs
{

    public sealed class InstanceConfigArgs : Pulumi.ResourceArgs
    {
        [Input("autoCreateTopicEnable", required: true)]
        public Input<bool> AutoCreateTopicEnable { get; set; } = null!;

        [Input("defaultNumPartitions", required: true)]
        public Input<int> DefaultNumPartitions { get; set; } = null!;

        [Input("defaultReplicationFactor", required: true)]
        public Input<int> DefaultReplicationFactor { get; set; } = null!;

        public InstanceConfigArgs()
        {
        }
    }
}
