// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListGetArgs : Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("extTimeOffsetLists")]
        private InputList<string>? _extTimeOffsetLists;
        public InputList<string> ExtTimeOffsetLists
        {
            get => _extTimeOffsetLists ?? (_extTimeOffsetLists = new InputList<string>());
            set => _extTimeOffsetLists = value;
        }

        [Input("watermarkLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListWatermarkListGetArgs>? _watermarkLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListWatermarkListGetArgs> WatermarkLists
        {
            get => _watermarkLists ?? (_watermarkLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListWatermarkListGetArgs>());
            set => _watermarkLists = value;
        }

        public ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListGetArgs()
        {
        }
    }
}
