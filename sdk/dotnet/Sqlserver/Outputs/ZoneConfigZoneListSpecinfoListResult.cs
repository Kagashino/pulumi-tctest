// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver.Outputs
{

    [OutputType]
    public sealed class ZoneConfigZoneListSpecinfoListResult
    {
        public readonly string ChargeType;
        public readonly int Cpu;
        public readonly string DbVersion;
        public readonly string DbVersionName;
        public readonly string MachineType;
        public readonly int MaxStorageSize;
        public readonly int Memory;
        public readonly int MinStorageSize;
        public readonly int Qps;
        public readonly int SpecId;

        [OutputConstructor]
        private ZoneConfigZoneListSpecinfoListResult(
            string chargeType,

            int cpu,

            string dbVersion,

            string dbVersionName,

            string machineType,

            int maxStorageSize,

            int memory,

            int minStorageSize,

            int qps,

            int specId)
        {
            ChargeType = chargeType;
            Cpu = cpu;
            DbVersion = dbVersion;
            DbVersionName = dbVersionName;
            MachineType = machineType;
            MaxStorageSize = maxStorageSize;
            Memory = memory;
            MinStorageSize = minStorageSize;
            Qps = qps;
            SpecId = specId;
        }
    }
}
