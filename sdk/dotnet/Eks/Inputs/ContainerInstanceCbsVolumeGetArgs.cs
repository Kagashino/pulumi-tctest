// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Eks.Inputs
{

    public sealed class ContainerInstanceCbsVolumeGetArgs : Pulumi.ResourceArgs
    {
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ContainerInstanceCbsVolumeGetArgs()
        {
        }
    }
}
