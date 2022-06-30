// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf
{
    [TencentcloudResourceType("tencentcloud:Scf/layer:Layer")]
    public partial class Layer : Pulumi.CustomResource
    {
        /// <summary>
        /// The code type of layer.
        /// </summary>
        [Output("codeSha256")]
        public Output<string> CodeSha256 { get; private set; } = null!;

        /// <summary>
        /// The compatible runtimes of layer.
        /// </summary>
        [Output("compatibleRuntimes")]
        public Output<ImmutableArray<string>> CompatibleRuntimes { get; private set; } = null!;

        /// <summary>
        /// The source code of layer.
        /// </summary>
        [Output("content")]
        public Output<Outputs.LayerContent> Content { get; private set; } = null!;

        /// <summary>
        /// The create time of layer.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of layer.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of layer.
        /// </summary>
        [Output("layerName")]
        public Output<string> LayerName { get; private set; } = null!;

        /// <summary>
        /// The version of layer.
        /// </summary>
        [Output("layerVersion")]
        public Output<int> LayerVersion { get; private set; } = null!;

        /// <summary>
        /// The license info of layer.
        /// </summary>
        [Output("licenseInfo")]
        public Output<string?> LicenseInfo { get; private set; } = null!;

        /// <summary>
        /// The download location url of layer.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The current status of layer.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Layer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Layer(string name, LayerArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/layer:Layer", name, args ?? new LayerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Layer(string name, Input<string> id, LayerState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/layer:Layer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Layer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Layer Get(string name, Input<string> id, LayerState? state = null, CustomResourceOptions? options = null)
        {
            return new Layer(name, id, state, options);
        }
    }

    public sealed class LayerArgs : Pulumi.ResourceArgs
    {
        [Input("compatibleRuntimes", required: true)]
        private InputList<string>? _compatibleRuntimes;

        /// <summary>
        /// The compatible runtimes of layer.
        /// </summary>
        public InputList<string> CompatibleRuntimes
        {
            get => _compatibleRuntimes ?? (_compatibleRuntimes = new InputList<string>());
            set => _compatibleRuntimes = value;
        }

        /// <summary>
        /// The source code of layer.
        /// </summary>
        [Input("content", required: true)]
        public Input<Inputs.LayerContentArgs> Content { get; set; } = null!;

        /// <summary>
        /// The description of layer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of layer.
        /// </summary>
        [Input("layerName", required: true)]
        public Input<string> LayerName { get; set; } = null!;

        /// <summary>
        /// The license info of layer.
        /// </summary>
        [Input("licenseInfo")]
        public Input<string>? LicenseInfo { get; set; }

        public LayerArgs()
        {
        }
    }

    public sealed class LayerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The code type of layer.
        /// </summary>
        [Input("codeSha256")]
        public Input<string>? CodeSha256 { get; set; }

        [Input("compatibleRuntimes")]
        private InputList<string>? _compatibleRuntimes;

        /// <summary>
        /// The compatible runtimes of layer.
        /// </summary>
        public InputList<string> CompatibleRuntimes
        {
            get => _compatibleRuntimes ?? (_compatibleRuntimes = new InputList<string>());
            set => _compatibleRuntimes = value;
        }

        /// <summary>
        /// The source code of layer.
        /// </summary>
        [Input("content")]
        public Input<Inputs.LayerContentGetArgs>? Content { get; set; }

        /// <summary>
        /// The create time of layer.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of layer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of layer.
        /// </summary>
        [Input("layerName")]
        public Input<string>? LayerName { get; set; }

        /// <summary>
        /// The version of layer.
        /// </summary>
        [Input("layerVersion")]
        public Input<int>? LayerVersion { get; set; }

        /// <summary>
        /// The license info of layer.
        /// </summary>
        [Input("licenseInfo")]
        public Input<string>? LicenseInfo { get; set; }

        /// <summary>
        /// The download location url of layer.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The current status of layer.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public LayerState()
        {
        }
    }
}
