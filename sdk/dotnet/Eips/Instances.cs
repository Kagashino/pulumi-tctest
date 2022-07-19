// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Eips
{
    public static class Instances
    {
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("tctest:Eips/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesResult>("tctest:Eips/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : Pulumi.InvokeArgs
    {
        [Input("eipId")]
        public string? EipId { get; set; }

        [Input("eipName")]
        public string? EipName { get; set; }

        [Input("publicIp")]
        public string? PublicIp { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public InstancesArgs()
        {
        }
    }

    public sealed class InstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        [Input("eipName")]
        public Input<string>? EipName { get; set; }

        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public InstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesResult
    {
        public readonly string? EipId;
        public readonly ImmutableArray<Outputs.InstancesEipListResult> EipLists;
        public readonly string? EipName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? PublicIp;
        public readonly string? ResultOutputFile;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private InstancesResult(
            string? eipId,

            ImmutableArray<Outputs.InstancesEipListResult> eipLists,

            string? eipName,

            string id,

            string? publicIp,

            string? resultOutputFile,

            ImmutableDictionary<string, object>? tags)
        {
            EipId = eipId;
            EipLists = eipLists;
            EipName = eipName;
            Id = id;
            PublicIp = publicIp;
            ResultOutputFile = resultOutputFile;
            Tags = tags;
        }
    }
}
