// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Postgresql
{
    public static class Xlogs
    {
        public static Task<XlogsResult> InvokeAsync(XlogsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<XlogsResult>("tctest:Postgresql/xlogs:Xlogs", args ?? new XlogsArgs(), options.WithDefaults());

        public static Output<XlogsResult> Invoke(XlogsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<XlogsResult>("tctest:Postgresql/xlogs:Xlogs", args ?? new XlogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class XlogsArgs : Pulumi.InvokeArgs
    {
        [Input("endTime")]
        public string? EndTime { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime")]
        public string? StartTime { get; set; }

        public XlogsArgs()
        {
        }
    }

    public sealed class XlogsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public XlogsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class XlogsResult
    {
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.XlogsListResult> Lists;
        public readonly string? ResultOutputFile;
        public readonly string? StartTime;

        [OutputConstructor]
        private XlogsResult(
            string? endTime,

            string id,

            string instanceId,

            ImmutableArray<Outputs.XlogsListResult> lists,

            string? resultOutputFile,

            string? startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}
