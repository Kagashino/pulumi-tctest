// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Outputs
{

    [OutputType]
    public sealed class AdaptiveDynamicStreamingTemplateStreamInfoVideo
    {
        public readonly int Bitrate;
        public readonly string Codec;
        public readonly string? FillType;
        public readonly int Fps;
        public readonly int? Height;
        public readonly bool? ResolutionAdaptive;
        public readonly int? Width;

        [OutputConstructor]
        private AdaptiveDynamicStreamingTemplateStreamInfoVideo(
            int bitrate,

            string codec,

            string? fillType,

            int fps,

            int? height,

            bool? resolutionAdaptive,

            int? width)
        {
            Bitrate = bitrate;
            Codec = codec;
            FillType = fillType;
            Fps = fps;
            Height = height;
            ResolutionAdaptive = resolutionAdaptive;
            Width = width;
        }
    }
}
