// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dnspod
{
    [TencentcloudResourceType("tencentcloud:Dnspod/domainInstance:DomainInstance")]
    public partial class DomainInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Create time of the domain.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The Group Id of Domain.
        /// </summary>
        [Output("groupId")]
        public Output<int?> GroupId { get; private set; } = null!;

        /// <summary>
        /// Whether to Mark the Domain.
        /// </summary>
        [Output("isMark")]
        public Output<string> IsMark { get; private set; } = null!;

        /// <summary>
        /// The remark of Domain.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// The status of Domain.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DomainInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainInstance(string name, DomainInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/domainInstance:DomainInstance", name, args ?? new DomainInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainInstance(string name, Input<string> id, DomainInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/domainInstance:DomainInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainInstance Get(string name, Input<string> id, DomainInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainInstance(name, id, state, options);
        }
    }

    public sealed class DomainInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The Group Id of Domain.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Whether to Mark the Domain.
        /// </summary>
        [Input("isMark")]
        public Input<string>? IsMark { get; set; }

        /// <summary>
        /// The remark of Domain.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The status of Domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DomainInstanceArgs()
        {
        }
    }

    public sealed class DomainInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time of the domain.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The Group Id of Domain.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Whether to Mark the Domain.
        /// </summary>
        [Input("isMark")]
        public Input<string>? IsMark { get; set; }

        /// <summary>
        /// The remark of Domain.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The status of Domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DomainInstanceState()
        {
        }
    }
}
