// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListWatermarkListArgs : Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        [Input("svgContent")]
        public Input<string>? SvgContent { get; set; }

        [Input("textContent")]
        public Input<string>? TextContent { get; set; }

        public ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListWatermarkListArgs()
        {
        }
    }
}
