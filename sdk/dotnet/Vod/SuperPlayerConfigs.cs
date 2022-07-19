// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod
{
    public static class SuperPlayerConfigs
    {
        public static Task<SuperPlayerConfigsResult> InvokeAsync(SuperPlayerConfigsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SuperPlayerConfigsResult>("tctest:Vod/superPlayerConfigs:SuperPlayerConfigs", args ?? new SuperPlayerConfigsArgs(), options.WithDefaults());

        public static Output<SuperPlayerConfigsResult> Invoke(SuperPlayerConfigsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SuperPlayerConfigsResult>("tctest:Vod/superPlayerConfigs:SuperPlayerConfigs", args ?? new SuperPlayerConfigsInvokeArgs(), options.WithDefaults());
    }


    public sealed class SuperPlayerConfigsArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public int? SubAppId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public SuperPlayerConfigsArgs()
        {
        }
    }

    public sealed class SuperPlayerConfigsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public SuperPlayerConfigsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SuperPlayerConfigsResult
    {
        public readonly ImmutableArray<Outputs.SuperPlayerConfigsConfigListResult> ConfigLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        public readonly string? Type;

        [OutputConstructor]
        private SuperPlayerConfigsResult(
            ImmutableArray<Outputs.SuperPlayerConfigsConfigListResult> configLists,

            string id,

            string? name,

            string? resultOutputFile,

            int? subAppId,

            string? type)
        {
            ConfigLists = configLists;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            Type = type;
        }
    }
}
