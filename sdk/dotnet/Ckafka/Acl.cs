// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka
{
    [TencentcloudResourceType("tencentcloud:Ckafka/acl:Acl")]
    public partial class Acl : Pulumi.CustomResource
    {
        /// <summary>
        /// IP address allowed to access. The default value is `*`, which means that any host can access.
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// ID of the ckafka instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`,
        /// `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        /// </summary>
        [Output("operationType")]
        public Output<string> OperationType { get; private set; } = null!;

        /// <summary>
        /// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports
        /// `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        /// </summary>
        [Output("permissionType")]
        public Output<string?> PermissionType { get; private set; } = null!;

        /// <summary>
        /// User list. The default value is `*`, which means that any user can access. The current user can only be one included in
        /// the user list.
        /// </summary>
        [Output("principal")]
        public Output<string?> Principal { get; private set; } = null!;

        /// <summary>
        /// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates
        /// the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        /// </summary>
        [Output("resourceName")]
        public Output<string> ResourceName { get; private set; } = null!;

        /// <summary>
        /// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by
        /// default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source
        /// Kafka.
        /// </summary>
        [Output("resourceType")]
        public Output<string?> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ckafka/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ckafka/acl:Acl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acl Get(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
        {
            return new Acl(name, id, state, options);
        }
    }

    public sealed class AclArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address allowed to access. The default value is `*`, which means that any host can access.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// ID of the ckafka instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`,
        /// `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        /// </summary>
        [Input("operationType", required: true)]
        public Input<string> OperationType { get; set; } = null!;

        /// <summary>
        /// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports
        /// `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        /// </summary>
        [Input("permissionType")]
        public Input<string>? PermissionType { get; set; }

        /// <summary>
        /// User list. The default value is `*`, which means that any user can access. The current user can only be one included in
        /// the user list.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates
        /// the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by
        /// default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source
        /// Kafka.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public AclArgs()
        {
        }
    }

    public sealed class AclState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address allowed to access. The default value is `*`, which means that any host can access.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// ID of the ckafka instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`,
        /// `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        /// </summary>
        [Input("operationType")]
        public Input<string>? OperationType { get; set; }

        /// <summary>
        /// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports
        /// `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        /// </summary>
        [Input("permissionType")]
        public Input<string>? PermissionType { get; set; }

        /// <summary>
        /// User list. The default value is `*`, which means that any user can access. The current user can only be one included in
        /// the user list.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates
        /// the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by
        /// default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source
        /// Kafka.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public AclState()
        {
        }
    }
}
