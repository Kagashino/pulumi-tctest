// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cam
{
    [TctestResourceType("tctest:Cam/rolePolicyAttachment:RolePolicyAttachment")]
    public partial class RolePolicyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and
        /// the others indicate syntax strategy ways.
        /// </summary>
        [Output("createMode")]
        public Output<int> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The create time of the CAM role policy attachment.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// ID of the attached CAM role.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a RolePolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RolePolicyAttachment(string name, RolePolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tctest:Cam/rolePolicyAttachment:RolePolicyAttachment", name, args ?? new RolePolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RolePolicyAttachment(string name, Input<string> id, RolePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Cam/rolePolicyAttachment:RolePolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RolePolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RolePolicyAttachment Get(string name, Input<string> id, RolePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new RolePolicyAttachment(name, id, state, options);
        }
    }

    public sealed class RolePolicyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// ID of the attached CAM role.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public RolePolicyAttachmentArgs()
        {
        }
    }

    public sealed class RolePolicyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and
        /// the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public Input<int>? CreateMode { get; set; }

        /// <summary>
        /// The create time of the CAM role policy attachment.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// ID of the attached CAM role.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public RolePolicyAttachmentState()
        {
        }
    }
}
