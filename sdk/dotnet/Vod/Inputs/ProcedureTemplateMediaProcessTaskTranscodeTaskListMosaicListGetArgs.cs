// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicListGetArgs : Pulumi.ResourceArgs
    {
        [Input("coordinateOrigin")]
        public Input<string>? CoordinateOrigin { get; set; }

        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        [Input("height")]
        public Input<string>? Height { get; set; }

        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        [Input("width")]
        public Input<string>? Width { get; set; }

        [Input("xPos")]
        public Input<string>? XPos { get; set; }

        [Input("yPos")]
        public Input<string>? YPos { get; set; }

        public ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicListGetArgs()
        {
        }
    }
}
