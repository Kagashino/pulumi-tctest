// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eks.Outputs
{

    [OutputType]
    public sealed class ContainerInstanceCbsVolume
    {
        public readonly string DiskId;
        public readonly string Name;

        [OutputConstructor]
        private ContainerInstanceCbsVolume(
            string diskId,

            string name)
        {
            DiskId = diskId;
            Name = name;
        }
    }
}
