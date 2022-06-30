// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb
{
    [TencentcloudResourceType("tencentcloud:Clb/targetGroup:TargetGroup")]
    public partial class TargetGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The default port of target group, add server after can use it.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The backend server of target group bind.
        /// </summary>
        [Output("targetGroupInstances")]
        public Output<ImmutableArray<Outputs.TargetGroupTargetGroupInstance>> TargetGroupInstances { get; private set; } = null!;

        /// <summary>
        /// Target group name.
        /// </summary>
        [Output("targetGroupName")]
        public Output<string?> TargetGroupName { get; private set; } = null!;

        /// <summary>
        /// VPC ID, default is based on the network.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a TargetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetGroup(string name, TargetGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/targetGroup:TargetGroup", name, args ?? new TargetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetGroup(string name, Input<string> id, TargetGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/targetGroup:TargetGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetGroup Get(string name, Input<string> id, TargetGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetGroup(name, id, state, options);
        }
    }

    public sealed class TargetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default port of target group, add server after can use it.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("targetGroupInstances")]
        private InputList<Inputs.TargetGroupTargetGroupInstanceArgs>? _targetGroupInstances;

        /// <summary>
        /// The backend server of target group bind.
        /// </summary>
        public InputList<Inputs.TargetGroupTargetGroupInstanceArgs> TargetGroupInstances
        {
            get => _targetGroupInstances ?? (_targetGroupInstances = new InputList<Inputs.TargetGroupTargetGroupInstanceArgs>());
            set => _targetGroupInstances = value;
        }

        /// <summary>
        /// Target group name.
        /// </summary>
        [Input("targetGroupName")]
        public Input<string>? TargetGroupName { get; set; }

        /// <summary>
        /// VPC ID, default is based on the network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public TargetGroupArgs()
        {
        }
    }

    public sealed class TargetGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default port of target group, add server after can use it.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("targetGroupInstances")]
        private InputList<Inputs.TargetGroupTargetGroupInstanceGetArgs>? _targetGroupInstances;

        /// <summary>
        /// The backend server of target group bind.
        /// </summary>
        public InputList<Inputs.TargetGroupTargetGroupInstanceGetArgs> TargetGroupInstances
        {
            get => _targetGroupInstances ?? (_targetGroupInstances = new InputList<Inputs.TargetGroupTargetGroupInstanceGetArgs>());
            set => _targetGroupInstances = value;
        }

        /// <summary>
        /// Target group name.
        /// </summary>
        [Input("targetGroupName")]
        public Input<string>? TargetGroupName { get; set; }

        /// <summary>
        /// VPC ID, default is based on the network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public TargetGroupState()
        {
        }
    }
}
