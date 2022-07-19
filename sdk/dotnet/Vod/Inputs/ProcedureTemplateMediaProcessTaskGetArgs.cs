// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskGetArgs : Pulumi.ResourceArgs
    {
        [Input("adaptiveDynamicStreamingTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListGetArgs>? _adaptiveDynamicStreamingTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListGetArgs> AdaptiveDynamicStreamingTaskLists
        {
            get => _adaptiveDynamicStreamingTaskLists ?? (_adaptiveDynamicStreamingTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListGetArgs>());
            set => _adaptiveDynamicStreamingTaskLists = value;
        }

        [Input("animatedGraphicTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskAnimatedGraphicTaskListGetArgs>? _animatedGraphicTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskAnimatedGraphicTaskListGetArgs> AnimatedGraphicTaskLists
        {
            get => _animatedGraphicTaskLists ?? (_animatedGraphicTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskAnimatedGraphicTaskListGetArgs>());
            set => _animatedGraphicTaskLists = value;
        }

        [Input("coverBySnapshotTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListGetArgs>? _coverBySnapshotTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListGetArgs> CoverBySnapshotTaskLists
        {
            get => _coverBySnapshotTaskLists ?? (_coverBySnapshotTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListGetArgs>());
            set => _coverBySnapshotTaskLists = value;
        }

        [Input("imageSpriteTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskImageSpriteTaskListGetArgs>? _imageSpriteTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskImageSpriteTaskListGetArgs> ImageSpriteTaskLists
        {
            get => _imageSpriteTaskLists ?? (_imageSpriteTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskImageSpriteTaskListGetArgs>());
            set => _imageSpriteTaskLists = value;
        }

        [Input("sampleSnapshotTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListGetArgs>? _sampleSnapshotTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListGetArgs> SampleSnapshotTaskLists
        {
            get => _sampleSnapshotTaskLists ?? (_sampleSnapshotTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListGetArgs>());
            set => _sampleSnapshotTaskLists = value;
        }

        [Input("snapshotByTimeOffsetTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListGetArgs>? _snapshotByTimeOffsetTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListGetArgs> SnapshotByTimeOffsetTaskLists
        {
            get => _snapshotByTimeOffsetTaskLists ?? (_snapshotByTimeOffsetTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListGetArgs>());
            set => _snapshotByTimeOffsetTaskLists = value;
        }

        [Input("transcodeTaskLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListGetArgs>? _transcodeTaskLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListGetArgs> TranscodeTaskLists
        {
            get => _transcodeTaskLists ?? (_transcodeTaskLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskTranscodeTaskListGetArgs>());
            set => _transcodeTaskLists = value;
        }

        public ProcedureTemplateMediaProcessTaskGetArgs()
        {
        }
    }
}
