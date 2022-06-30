// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod
{
    public static class SnapshotByTimeOffsetTemplates
    {
        public static Task<SnapshotByTimeOffsetTemplatesResult> InvokeAsync(SnapshotByTimeOffsetTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SnapshotByTimeOffsetTemplatesResult>("tencentcloud:Vod/snapshotByTimeOffsetTemplates:SnapshotByTimeOffsetTemplates", args ?? new SnapshotByTimeOffsetTemplatesArgs(), options.WithDefaults());

        public static Output<SnapshotByTimeOffsetTemplatesResult> Invoke(SnapshotByTimeOffsetTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SnapshotByTimeOffsetTemplatesResult>("tencentcloud:Vod/snapshotByTimeOffsetTemplates:SnapshotByTimeOffsetTemplates", args ?? new SnapshotByTimeOffsetTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class SnapshotByTimeOffsetTemplatesArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public string? Definition { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public int? SubAppId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public SnapshotByTimeOffsetTemplatesArgs()
        {
        }
    }

    public sealed class SnapshotByTimeOffsetTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public SnapshotByTimeOffsetTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SnapshotByTimeOffsetTemplatesResult
    {
        public readonly string? Definition;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        public readonly ImmutableArray<Outputs.SnapshotByTimeOffsetTemplatesTemplateListResult> TemplateLists;
        public readonly string? Type;

        [OutputConstructor]
        private SnapshotByTimeOffsetTemplatesResult(
            string? definition,

            string id,

            string? resultOutputFile,

            int? subAppId,

            ImmutableArray<Outputs.SnapshotByTimeOffsetTemplatesTemplateListResult> templateLists,

            string? type)
        {
            Definition = definition;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            TemplateLists = templateLists;
            Type = type;
        }
    }
}