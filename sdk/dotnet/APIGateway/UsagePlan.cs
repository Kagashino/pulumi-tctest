// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.APIGateway
{
    [TencentcloudResourceType("tencentcloud:APIGateway/usagePlan:UsagePlan")]
    public partial class UsagePlan : Pulumi.CustomResource
    {
        /// <summary>
        /// Attach API keys list.
        /// </summary>
        [Output("attachApiKeys")]
        public Output<ImmutableArray<string>> AttachApiKeys { get; private set; } = null!;

        /// <summary>
        /// Attach service and API list.
        /// </summary>
        [Output("attachLists")]
        public Output<ImmutableArray<Outputs.UsagePlanAttachList>> AttachLists { get; private set; } = null!;

        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
        /// </summary>
        [Output("maxRequestNum")]
        public Output<int?> MaxRequestNum { get; private set; } = null!;

        /// <summary>
        /// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
        /// </summary>
        [Output("maxRequestNumPreSec")]
        public Output<int?> MaxRequestNumPreSec { get; private set; } = null!;

        /// <summary>
        /// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// Custom usage plan description.
        /// </summary>
        [Output("usagePlanDesc")]
        public Output<string?> UsagePlanDesc { get; private set; } = null!;

        /// <summary>
        /// Custom usage plan name.
        /// </summary>
        [Output("usagePlanName")]
        public Output<string> UsagePlanName { get; private set; } = null!;


        /// <summary>
        /// Create a UsagePlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsagePlan(string name, UsagePlanArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:APIGateway/usagePlan:UsagePlan", name, args ?? new UsagePlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UsagePlan(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:APIGateway/usagePlan:UsagePlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsagePlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsagePlan Get(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
        {
            return new UsagePlan(name, id, state, options);
        }
    }

    public sealed class UsagePlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
        /// </summary>
        [Input("maxRequestNum")]
        public Input<int>? MaxRequestNum { get; set; }

        /// <summary>
        /// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
        /// </summary>
        [Input("maxRequestNumPreSec")]
        public Input<int>? MaxRequestNumPreSec { get; set; }

        /// <summary>
        /// Custom usage plan description.
        /// </summary>
        [Input("usagePlanDesc")]
        public Input<string>? UsagePlanDesc { get; set; }

        /// <summary>
        /// Custom usage plan name.
        /// </summary>
        [Input("usagePlanName", required: true)]
        public Input<string> UsagePlanName { get; set; } = null!;

        public UsagePlanArgs()
        {
        }
    }

    public sealed class UsagePlanState : Pulumi.ResourceArgs
    {
        [Input("attachApiKeys")]
        private InputList<string>? _attachApiKeys;

        /// <summary>
        /// Attach API keys list.
        /// </summary>
        public InputList<string> AttachApiKeys
        {
            get => _attachApiKeys ?? (_attachApiKeys = new InputList<string>());
            set => _attachApiKeys = value;
        }

        [Input("attachLists")]
        private InputList<Inputs.UsagePlanAttachListGetArgs>? _attachLists;

        /// <summary>
        /// Attach service and API list.
        /// </summary>
        public InputList<Inputs.UsagePlanAttachListGetArgs> AttachLists
        {
            get => _attachLists ?? (_attachLists = new InputList<Inputs.UsagePlanAttachListGetArgs>());
            set => _attachLists = value;
        }

        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
        /// </summary>
        [Input("maxRequestNum")]
        public Input<int>? MaxRequestNum { get; set; }

        /// <summary>
        /// Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
        /// </summary>
        [Input("maxRequestNumPreSec")]
        public Input<int>? MaxRequestNumPreSec { get; set; }

        /// <summary>
        /// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// Custom usage plan description.
        /// </summary>
        [Input("usagePlanDesc")]
        public Input<string>? UsagePlanDesc { get; set; }

        /// <summary>
        /// Custom usage plan name.
        /// </summary>
        [Input("usagePlanName")]
        public Input<string>? UsagePlanName { get; set; }

        public UsagePlanState()
        {
        }
    }
}
