// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Clb
{
    [TctestResourceType("tctest:Clb/targetGroupAttachment:TargetGroupAttachment")]
    public partial class TargetGroupAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the CLB.
        /// </summary>
        [Output("clbId")]
        public Output<string> ClbId { get; private set; } = null!;

        /// <summary>
        /// ID of the CLB listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// ID of the CLB listener rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string?> RuleId { get; private set; } = null!;

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Output("targetGroupId")]
        public Output<string?> TargetGroupId { get; private set; } = null!;

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Output("targrtGroupId")]
        public Output<string?> TargrtGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a TargetGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetGroupAttachment(string name, TargetGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tctest:Clb/targetGroupAttachment:TargetGroupAttachment", name, args ?? new TargetGroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetGroupAttachment(string name, Input<string> id, TargetGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Clb/targetGroupAttachment:TargetGroupAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetGroupAttachment Get(string name, Input<string> id, TargetGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetGroupAttachment(name, id, state, options);
        }
    }

    public sealed class TargetGroupAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the CLB.
        /// </summary>
        [Input("clbId", required: true)]
        public Input<string> ClbId { get; set; } = null!;

        /// <summary>
        /// ID of the CLB listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// ID of the CLB listener rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Input("targetGroupId")]
        public Input<string>? TargetGroupId { get; set; }

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Input("targrtGroupId")]
        public Input<string>? TargrtGroupId { get; set; }

        public TargetGroupAttachmentArgs()
        {
        }
    }

    public sealed class TargetGroupAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the CLB.
        /// </summary>
        [Input("clbId")]
        public Input<string>? ClbId { get; set; }

        /// <summary>
        /// ID of the CLB listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// ID of the CLB listener rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Input("targetGroupId")]
        public Input<string>? TargetGroupId { get; set; }

        /// <summary>
        /// ID of the CLB target group.
        /// </summary>
        [Input("targrtGroupId")]
        public Input<string>? TargrtGroupId { get; set; }

        public TargetGroupAttachmentState()
        {
        }
    }
}
