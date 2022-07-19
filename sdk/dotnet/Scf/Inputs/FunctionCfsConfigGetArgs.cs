// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Scf.Inputs
{

    public sealed class FunctionCfsConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("cfsId", required: true)]
        public Input<string> CfsId { get; set; } = null!;

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("localMountDir", required: true)]
        public Input<string> LocalMountDir { get; set; } = null!;

        [Input("mountInsId", required: true)]
        public Input<string> MountInsId { get; set; } = null!;

        [Input("mountSubnetId")]
        public Input<string>? MountSubnetId { get; set; }

        [Input("mountVpcId")]
        public Input<string>? MountVpcId { get; set; }

        [Input("remoteMountDir", required: true)]
        public Input<string> RemoteMountDir { get; set; } = null!;

        [Input("userGroupId", required: true)]
        public Input<string> UserGroupId { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public FunctionCfsConfigGetArgs()
        {
        }
    }
}
