// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.APIGateway
{
    [TctestResourceType("tctest:APIGateway/usagePlanAttachment:UsagePlanAttachment")]
    public partial class UsagePlanAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the API. This parameter will be required when `bind_type` is `API`.
        /// </summary>
        [Output("apiId")]
        public Output<string?> ApiId { get; private set; } = null!;

        /// <summary>
        /// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        /// </summary>
        [Output("bindType")]
        public Output<string?> BindType { get; private set; } = null!;

        /// <summary>
        /// The environment to be bound. Valid values: `test`, `prepub`, `release`.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// ID of the service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// ID of the usage plan.
        /// </summary>
        [Output("usagePlanId")]
        public Output<string> UsagePlanId { get; private set; } = null!;


        /// <summary>
        /// Create a UsagePlanAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsagePlanAttachment(string name, UsagePlanAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tctest:APIGateway/usagePlanAttachment:UsagePlanAttachment", name, args ?? new UsagePlanAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UsagePlanAttachment(string name, Input<string> id, UsagePlanAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tctest:APIGateway/usagePlanAttachment:UsagePlanAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsagePlanAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsagePlanAttachment Get(string name, Input<string> id, UsagePlanAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new UsagePlanAttachment(name, id, state, options);
        }
    }

    public sealed class UsagePlanAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the API. This parameter will be required when `bind_type` is `API`.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        /// </summary>
        [Input("bindType")]
        public Input<string>? BindType { get; set; }

        /// <summary>
        /// The environment to be bound. Valid values: `test`, `prepub`, `release`.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// ID of the service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// ID of the usage plan.
        /// </summary>
        [Input("usagePlanId", required: true)]
        public Input<string> UsagePlanId { get; set; } = null!;

        public UsagePlanAttachmentArgs()
        {
        }
    }

    public sealed class UsagePlanAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the API. This parameter will be required when `bind_type` is `API`.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        /// </summary>
        [Input("bindType")]
        public Input<string>? BindType { get; set; }

        /// <summary>
        /// The environment to be bound. Valid values: `test`, `prepub`, `release`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// ID of the service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// ID of the usage plan.
        /// </summary>
        [Input("usagePlanId")]
        public Input<string>? UsagePlanId { get; set; }

        public UsagePlanAttachmentState()
        {
        }
    }
}
