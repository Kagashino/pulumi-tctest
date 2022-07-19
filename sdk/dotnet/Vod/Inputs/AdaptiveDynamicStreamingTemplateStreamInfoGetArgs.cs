// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Inputs
{

    public sealed class AdaptiveDynamicStreamingTemplateStreamInfoGetArgs : Pulumi.ResourceArgs
    {
        [Input("audio", required: true)]
        public Input<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoAudioGetArgs> Audio { get; set; } = null!;

        [Input("removeAudio")]
        public Input<bool>? RemoveAudio { get; set; }

        [Input("video", required: true)]
        public Input<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoVideoGetArgs> Video { get; set; } = null!;

        public AdaptiveDynamicStreamingTemplateStreamInfoGetArgs()
        {
        }
    }
}
