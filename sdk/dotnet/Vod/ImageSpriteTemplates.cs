// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod
{
    public static class ImageSpriteTemplates
    {
        public static Task<ImageSpriteTemplatesResult> InvokeAsync(ImageSpriteTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ImageSpriteTemplatesResult>("tctest:Vod/imageSpriteTemplates:ImageSpriteTemplates", args ?? new ImageSpriteTemplatesArgs(), options.WithDefaults());

        public static Output<ImageSpriteTemplatesResult> Invoke(ImageSpriteTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ImageSpriteTemplatesResult>("tctest:Vod/imageSpriteTemplates:ImageSpriteTemplates", args ?? new ImageSpriteTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ImageSpriteTemplatesArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public string? Definition { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public int? SubAppId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public ImageSpriteTemplatesArgs()
        {
        }
    }

    public sealed class ImageSpriteTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ImageSpriteTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ImageSpriteTemplatesResult
    {
        public readonly string? Definition;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        public readonly ImmutableArray<Outputs.ImageSpriteTemplatesTemplateListResult> TemplateLists;
        public readonly string? Type;

        [OutputConstructor]
        private ImageSpriteTemplatesResult(
            string? definition,

            string id,

            string? resultOutputFile,

            int? subAppId,

            ImmutableArray<Outputs.ImageSpriteTemplatesTemplateListResult> templateLists,

            string? type)
        {
            Definition = definition;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            TemplateLists = templateLists;
            Type = type;
        }
    }
}
