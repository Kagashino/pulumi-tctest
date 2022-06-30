// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cbs.Outputs
{

    [OutputType]
    public sealed class SnapshotsSnapshotListResult
    {
        public readonly string AvailabilityZone;
        public readonly string CreateTime;
        public readonly bool Encrypt;
        public readonly int Percent;
        public readonly int ProjectId;
        public readonly string SnapshotId;
        public readonly string SnapshotName;
        public readonly string StorageId;
        public readonly int StorageSize;
        public readonly string StorageUsage;

        [OutputConstructor]
        private SnapshotsSnapshotListResult(
            string availabilityZone,

            string createTime,

            bool encrypt,

            int percent,

            int projectId,

            string snapshotId,

            string snapshotName,

            string storageId,

            int storageSize,

            string storageUsage)
        {
            AvailabilityZone = availabilityZone;
            CreateTime = createTime;
            Encrypt = encrypt;
            Percent = percent;
            ProjectId = projectId;
            SnapshotId = snapshotId;
            SnapshotName = snapshotName;
            StorageId = storageId;
            StorageSize = storageSize;
            StorageUsage = storageUsage;
        }
    }
}
