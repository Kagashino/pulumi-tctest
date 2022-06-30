// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cloud
{
    public static class TCRInstances
    {
        public static Task<TCRInstancesResult> InvokeAsync(TCRInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<TCRInstancesResult>("tencentcloud:Cloud/tCRInstances:TCRInstances", args ?? new TCRInstancesArgs(), options.WithDefaults());

        public static Output<TCRInstancesResult> Invoke(TCRInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<TCRInstancesResult>("tencentcloud:Cloud/tCRInstances:TCRInstances", args ?? new TCRInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class TCRInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public TCRInstancesArgs()
        {
        }
    }

    public sealed class TCRInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public TCRInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class TCRInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly ImmutableArray<Outputs.TCRInstancesInstanceListResult> InstanceLists;
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private TCRInstancesResult(
            string id,

            string? instanceId,

            ImmutableArray<Outputs.TCRInstancesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
