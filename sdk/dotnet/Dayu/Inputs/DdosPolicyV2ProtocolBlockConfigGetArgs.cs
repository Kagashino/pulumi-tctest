// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Inputs
{

    public sealed class DdosPolicyV2ProtocolBlockConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("dropIcmp", required: true)]
        public Input<int> DropIcmp { get; set; } = null!;

        [Input("dropOther", required: true)]
        public Input<int> DropOther { get; set; } = null!;

        [Input("dropTcp", required: true)]
        public Input<int> DropTcp { get; set; } = null!;

        [Input("dropUdp", required: true)]
        public Input<int> DropUdp { get; set; } = null!;

        public DdosPolicyV2ProtocolBlockConfigGetArgs()
        {
        }
    }
}
