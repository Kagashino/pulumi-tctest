// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfs
{
    [TencentcloudResourceType("tencentcloud:Cfs/accessGroup:AccessGroup")]
    public partial class AccessGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Create time of the access group.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the access group, and max length is 255.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the access group, and max length is 64.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AccessGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessGroup(string name, AccessGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/accessGroup:AccessGroup", name, args ?? new AccessGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessGroup(string name, Input<string> id, AccessGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/accessGroup:AccessGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessGroup Get(string name, Input<string> id, AccessGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessGroup(name, id, state, options);
        }
    }

    public sealed class AccessGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the access group, and max length is 255.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the access group, and max length is 64.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AccessGroupArgs()
        {
        }
    }

    public sealed class AccessGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time of the access group.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the access group, and max length is 255.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the access group, and max length is 64.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AccessGroupState()
        {
        }
    }
}