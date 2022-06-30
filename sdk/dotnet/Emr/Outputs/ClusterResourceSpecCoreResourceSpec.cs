// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterResourceSpecCoreResourceSpec
    {
        public readonly int? Cpu;
        public readonly int? DiskSize;
        public readonly string? DiskType;
        public readonly int? MemSize;
        public readonly int? RootSize;
        public readonly string? Spec;
        public readonly int? StorageType;

        [OutputConstructor]
        private ClusterResourceSpecCoreResourceSpec(
            int? cpu,

            int? diskSize,

            string? diskType,

            int? memSize,

            int? rootSize,

            string? spec,

            int? storageType)
        {
            Cpu = cpu;
            DiskSize = diskSize;
            DiskType = diskType;
            MemSize = memSize;
            RootSize = rootSize;
            Spec = spec;
            StorageType = storageType;
        }
    }
}
