// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskTranscodeTaskListGetArgs : Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("mosaicLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicListGetArgs>? _mosaicLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicListGetArgs> MosaicLists
        {
            get => _mosaicLists ?? (_mosaicLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicListGetArgs>());
            set => _mosaicLists = value;
        }

        [Input("watermarkLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListWatermarkListGetArgs>? _watermarkLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListWatermarkListGetArgs> WatermarkLists
        {
            get => _watermarkLists ?? (_watermarkLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListWatermarkListGetArgs>());
            set => _watermarkLists = value;
        }

        public ProcedureTemplateMediaProcessTaskTranscodeTaskListGetArgs()
        {
        }
    }
}
