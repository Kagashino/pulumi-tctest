// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Emr.Outputs
{

    [OutputType]
    public sealed class InstancesClusterResult
    {
        public readonly string AddTime;
        public readonly int ChargeType;
        public readonly string ClusterId;
        public readonly string ClusterName;
        public readonly string Ftitle;
        public readonly int Id;
        public readonly string MasterIp;
        public readonly int ProjectId;
        public readonly int RegionId;
        public readonly int Status;
        public readonly string Zone;
        public readonly int ZoneId;

        [OutputConstructor]
        private InstancesClusterResult(
            string addTime,

            int chargeType,

            string clusterId,

            string clusterName,

            string ftitle,

            int id,

            string masterIp,

            int projectId,

            int regionId,

            int status,

            string zone,

            int zoneId)
        {
            AddTime = addTime;
            ChargeType = chargeType;
            ClusterId = clusterId;
            ClusterName = clusterName;
            Ftitle = ftitle;
            Id = id;
            MasterIp = masterIp;
            ProjectId = projectId;
            RegionId = regionId;
            Status = status;
            Zone = zone;
            ZoneId = zoneId;
        }
    }
}
