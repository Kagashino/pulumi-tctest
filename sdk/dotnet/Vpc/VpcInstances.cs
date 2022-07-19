// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vpc
{
    public static class VpcInstances
    {
        public static Task<VpcInstancesResult> InvokeAsync(VpcInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<VpcInstancesResult>("tctest:Vpc/vpcInstances:VpcInstances", args ?? new VpcInstancesArgs(), options.WithDefaults());

        public static Output<VpcInstancesResult> Invoke(VpcInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<VpcInstancesResult>("tctest:Vpc/vpcInstances:VpcInstances", args ?? new VpcInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("tagKey")]
        public string? TagKey { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public VpcInstancesArgs()
        {
        }
    }

    public sealed class VpcInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class VpcInstancesResult
    {
        public readonly string? CidrBlock;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.VpcInstancesInstanceListResult> InstanceLists;
        public readonly bool? IsDefault;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly string? TagKey;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? VpcId;

        [OutputConstructor]
        private VpcInstancesResult(
            string? cidrBlock,

            string id,

            ImmutableArray<Outputs.VpcInstancesInstanceListResult> instanceLists,

            bool? isDefault,

            string? name,

            string? resultOutputFile,

            string? tagKey,

            ImmutableDictionary<string, object>? tags,

            string? vpcId)
        {
            CidrBlock = cidrBlock;
            Id = id;
            InstanceLists = instanceLists;
            IsDefault = isDefault;
            Name = name;
            ResultOutputFile = resultOutputFile;
            TagKey = tagKey;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
