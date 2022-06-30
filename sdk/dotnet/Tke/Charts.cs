// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tke
{
    public static class Charts
    {
        public static Task<ChartsResult> InvokeAsync(ChartsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ChartsResult>("tencentcloud:Tke/charts:Charts", args ?? new ChartsArgs(), options.WithDefaults());

        public static Output<ChartsResult> Invoke(ChartsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ChartsResult>("tencentcloud:Tke/charts:Charts", args ?? new ChartsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ChartsArgs : Pulumi.InvokeArgs
    {
        [Input("arch")]
        public string? Arch { get; set; }

        [Input("clusterType")]
        public string? ClusterType { get; set; }

        [Input("kind")]
        public string? Kind { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public ChartsArgs()
        {
        }
    }

    public sealed class ChartsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arch")]
        public Input<string>? Arch { get; set; }

        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public ChartsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ChartsResult
    {
        public readonly string? Arch;
        public readonly ImmutableArray<Outputs.ChartsChartListResult> ChartLists;
        public readonly string? ClusterType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Kind;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private ChartsResult(
            string? arch,

            ImmutableArray<Outputs.ChartsChartListResult> chartLists,

            string? clusterType,

            string id,

            string? kind,

            string? resultOutputFile)
        {
            Arch = arch;
            ChartLists = chartLists;
            ClusterType = clusterType;
            Id = id;
            Kind = kind;
            ResultOutputFile = resultOutputFile;
        }
    }
}