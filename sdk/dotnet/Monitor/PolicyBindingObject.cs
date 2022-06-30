// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    [TencentcloudResourceType("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject")]
    public partial class PolicyBindingObject : Pulumi.CustomResource
    {
        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.PolicyBindingObjectDimension>> Dimensions { get; private set; } = null!;

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyBindingObject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyBindingObject(string name, PolicyBindingObjectArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, args ?? new PolicyBindingObjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyBindingObject(string name, Input<string> id, PolicyBindingObjectState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyBindingObject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyBindingObject Get(string name, Input<string> id, PolicyBindingObjectState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyBindingObject(name, id, state, options);
        }
    }

    public sealed class PolicyBindingObjectArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions", required: true)]
        private InputList<Inputs.PolicyBindingObjectDimensionArgs>? _dimensions;

        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.PolicyBindingObjectDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.PolicyBindingObjectDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        public PolicyBindingObjectArgs()
        {
        }
    }

    public sealed class PolicyBindingObjectState : Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.PolicyBindingObjectDimensionGetArgs>? _dimensions;

        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.PolicyBindingObjectDimensionGetArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.PolicyBindingObjectDimensionGetArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        public PolicyBindingObjectState()
        {
        }
    }
}
