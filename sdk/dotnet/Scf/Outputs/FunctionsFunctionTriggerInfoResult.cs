// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class FunctionsFunctionTriggerInfoResult
    {
        public readonly string CreateTime;
        public readonly string CustomArgument;
        public readonly bool Enable;
        public readonly string ModifyTime;
        public readonly string Name;
        public readonly string TriggerDesc;
        public readonly string Type;

        [OutputConstructor]
        private FunctionsFunctionTriggerInfoResult(
            string createTime,

            string customArgument,

            bool enable,

            string modifyTime,

            string name,

            string triggerDesc,

            string type)
        {
            CreateTime = createTime;
            CustomArgument = customArgument;
            Enable = enable;
            ModifyTime = modifyTime;
            Name = name;
            TriggerDesc = triggerDesc;
            Type = type;
        }
    }
}
