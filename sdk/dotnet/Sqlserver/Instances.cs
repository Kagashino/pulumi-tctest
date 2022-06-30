// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    public static class Instances
    {
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("tencentcloud:Sqlserver/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesResult>("tencentcloud:Sqlserver/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public InstancesArgs()
        {
        }
    }

    public sealed class InstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public InstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.InstancesInstanceListResult> InstanceLists;
        public readonly string? Name;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;
        public readonly string? SubnetId;
        public readonly string? VpcId;

        [OutputConstructor]
        private InstancesResult(
            string? id,

            ImmutableArray<Outputs.InstancesInstanceListResult> instanceLists,

            string? name,

            int? projectId,

            string? resultOutputFile,

            string? subnetId,

            string? vpcId)
        {
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}