// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class FunctionTriggerInfoArgs : Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("customArgument")]
        public Input<string>? CustomArgument { get; set; }

        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("triggerDesc")]
        public Input<string>? TriggerDesc { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public FunctionTriggerInfoArgs()
        {
        }
    }
}
