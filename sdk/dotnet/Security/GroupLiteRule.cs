// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Security
{
    [TencentcloudResourceType("tencentcloud:Security/groupLiteRule:GroupLiteRule")]
    public partial class GroupLiteRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        [Output("egresses")]
        public Output<ImmutableArray<string>> Egresses { get; private set; } = null!;

        /// <summary>
        /// Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        [Output("ingresses")]
        public Output<ImmutableArray<string>> Ingresses { get; private set; } = null!;

        /// <summary>
        /// ID of the security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupLiteRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupLiteRule(string name, GroupLiteRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Security/groupLiteRule:GroupLiteRule", name, args ?? new GroupLiteRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupLiteRule(string name, Input<string> id, GroupLiteRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Security/groupLiteRule:GroupLiteRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupLiteRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupLiteRule Get(string name, Input<string> id, GroupLiteRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupLiteRule(name, id, state, options);
        }
    }

    public sealed class GroupLiteRuleArgs : Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<string>? _egresses;

        /// <summary>
        /// Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        public InputList<string> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<string>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<string>? _ingresses;

        /// <summary>
        /// Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        public InputList<string> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<string>());
            set => _ingresses = value;
        }

        /// <summary>
        /// ID of the security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public GroupLiteRuleArgs()
        {
        }
    }

    public sealed class GroupLiteRuleState : Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<string>? _egresses;

        /// <summary>
        /// Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        public InputList<string> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<string>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<string>? _ingresses;

        /// <summary>
        /// Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of
        /// 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address
        /// Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`,
        /// `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        /// </summary>
        public InputList<string> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<string>());
            set => _ingresses = value;
        }

        /// <summary>
        /// ID of the security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public GroupLiteRuleState()
        {
        }
    }
}
