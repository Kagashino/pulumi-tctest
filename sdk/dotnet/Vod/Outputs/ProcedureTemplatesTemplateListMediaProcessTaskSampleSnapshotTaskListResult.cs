// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class ProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListResult
    {
        public readonly string Definition;
        public readonly ImmutableArray<Outputs.ProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListWatermarkListResult> WatermarkLists;

        [OutputConstructor]
        private ProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListResult(
            string definition,

            ImmutableArray<Outputs.ProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListWatermarkListResult> watermarkLists)
        {
            Definition = definition;
            WatermarkLists = watermarkLists;
        }
    }
}
