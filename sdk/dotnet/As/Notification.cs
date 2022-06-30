// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As
{
    [TencentcloudResourceType("tencentcloud:As/notification:Notification")]
    public partial class Notification : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`,
        /// `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and
        /// `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
        /// </summary>
        [Output("notificationTypes")]
        public Output<ImmutableArray<string>> NotificationTypes { get; private set; } = null!;

        /// <summary>
        /// A group of user IDs to be notified.
        /// </summary>
        [Output("notificationUserGroupIds")]
        public Output<ImmutableArray<string>> NotificationUserGroupIds { get; private set; } = null!;

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:As/notification:Notification", name, args ?? new NotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:As/notification:Notification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : Pulumi.ResourceArgs
    {
        [Input("notificationTypes", required: true)]
        private InputList<string>? _notificationTypes;

        /// <summary>
        /// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`,
        /// `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and
        /// `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
        /// </summary>
        public InputList<string> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<string>());
            set => _notificationTypes = value;
        }

        [Input("notificationUserGroupIds", required: true)]
        private InputList<string>? _notificationUserGroupIds;

        /// <summary>
        /// A group of user IDs to be notified.
        /// </summary>
        public InputList<string> NotificationUserGroupIds
        {
            get => _notificationUserGroupIds ?? (_notificationUserGroupIds = new InputList<string>());
            set => _notificationUserGroupIds = value;
        }

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        public NotificationArgs()
        {
        }
    }

    public sealed class NotificationState : Pulumi.ResourceArgs
    {
        [Input("notificationTypes")]
        private InputList<string>? _notificationTypes;

        /// <summary>
        /// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`,
        /// `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and
        /// `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
        /// </summary>
        public InputList<string> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<string>());
            set => _notificationTypes = value;
        }

        [Input("notificationUserGroupIds")]
        private InputList<string>? _notificationUserGroupIds;

        /// <summary>
        /// A group of user IDs to be notified.
        /// </summary>
        public InputList<string> NotificationUserGroupIds
        {
            get => _notificationUserGroupIds ?? (_notificationUserGroupIds = new InputList<string>());
            set => _notificationUserGroupIds = value;
        }

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public NotificationState()
        {
        }
    }
}
