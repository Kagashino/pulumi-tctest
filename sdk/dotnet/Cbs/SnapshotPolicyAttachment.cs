// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cbs
{
    [TencentcloudResourceType("tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment")]
    public partial class SnapshotPolicyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of CBS snapshot policy.
        /// </summary>
        [Output("snapshotPolicyId")]
        public Output<string> SnapshotPolicyId { get; private set; } = null!;

        /// <summary>
        /// ID of CBS.
        /// </summary>
        [Output("storageId")]
        public Output<string> StorageId { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotPolicyAttachment(string name, SnapshotPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment", name, args ?? new SnapshotPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotPolicyAttachment(string name, Input<string> id, SnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotPolicyAttachment Get(string name, Input<string> id, SnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class SnapshotPolicyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of CBS snapshot policy.
        /// </summary>
        [Input("snapshotPolicyId", required: true)]
        public Input<string> SnapshotPolicyId { get; set; } = null!;

        /// <summary>
        /// ID of CBS.
        /// </summary>
        [Input("storageId", required: true)]
        public Input<string> StorageId { get; set; } = null!;

        public SnapshotPolicyAttachmentArgs()
        {
        }
    }

    public sealed class SnapshotPolicyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of CBS snapshot policy.
        /// </summary>
        [Input("snapshotPolicyId")]
        public Input<string>? SnapshotPolicyId { get; set; }

        /// <summary>
        /// ID of CBS.
        /// </summary>
        [Input("storageId")]
        public Input<string>? StorageId { get; set; }

        public SnapshotPolicyAttachmentState()
        {
        }
    }
}
