// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ha
{
    [TencentcloudResourceType("tencentcloud:Ha/vipEipAttachment:VipEipAttachment")]
    public partial class VipEipAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Public address of the EIP.
        /// </summary>
        [Output("addressIp")]
        public Output<string> AddressIp { get; private set; } = null!;

        /// <summary>
        /// ID of the attached HA VIP.
        /// </summary>
        [Output("havipId")]
        public Output<string> HavipId { get; private set; } = null!;


        /// <summary>
        /// Create a VipEipAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VipEipAttachment(string name, VipEipAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ha/vipEipAttachment:VipEipAttachment", name, args ?? new VipEipAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VipEipAttachment(string name, Input<string> id, VipEipAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ha/vipEipAttachment:VipEipAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VipEipAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VipEipAttachment Get(string name, Input<string> id, VipEipAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new VipEipAttachment(name, id, state, options);
        }
    }

    public sealed class VipEipAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public address of the EIP.
        /// </summary>
        [Input("addressIp", required: true)]
        public Input<string> AddressIp { get; set; } = null!;

        /// <summary>
        /// ID of the attached HA VIP.
        /// </summary>
        [Input("havipId", required: true)]
        public Input<string> HavipId { get; set; } = null!;

        public VipEipAttachmentArgs()
        {
        }
    }

    public sealed class VipEipAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public address of the EIP.
        /// </summary>
        [Input("addressIp")]
        public Input<string>? AddressIp { get; set; }

        /// <summary>
        /// ID of the attached HA VIP.
        /// </summary>
        [Input("havipId")]
        public Input<string>? HavipId { get; set; }

        public VipEipAttachmentState()
        {
        }
    }
}
