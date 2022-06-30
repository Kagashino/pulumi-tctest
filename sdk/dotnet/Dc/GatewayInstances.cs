// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dc
{
    public static class GatewayInstances
    {
        public static Task<GatewayInstancesResult> InvokeAsync(GatewayInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GatewayInstancesResult>("tencentcloud:Dc/gatewayInstances:GatewayInstances", args ?? new GatewayInstancesArgs(), options.WithDefaults());

        public static Output<GatewayInstancesResult> Invoke(GatewayInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GatewayInstancesResult>("tencentcloud:Dc/gatewayInstances:GatewayInstances", args ?? new GatewayInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GatewayInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("dcgId")]
        public string? DcgId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GatewayInstancesArgs()
        {
        }
    }

    public sealed class GatewayInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dcgId")]
        public Input<string>? DcgId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GatewayInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GatewayInstancesResult
    {
        public readonly string? DcgId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GatewayInstancesInstanceListResult> InstanceLists;
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GatewayInstancesResult(
            string? dcgId,

            string id,

            ImmutableArray<Outputs.GatewayInstancesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile)
        {
            DcgId = dcgId;
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
