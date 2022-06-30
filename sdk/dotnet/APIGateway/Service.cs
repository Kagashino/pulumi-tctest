// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.APIGateway
{
    [TencentcloudResourceType("tencentcloud:APIGateway/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of APIs.
        /// </summary>
        [Output("apiLists")]
        public Output<ImmutableArray<Outputs.ServiceApiList>> ApiLists { get; private set; } = null!;

        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
        /// </summary>
        [Output("exclusiveSetName")]
        public Output<string?> ExclusiveSetName { get; private set; } = null!;

        /// <summary>
        /// Port number for http access over private network.
        /// </summary>
        [Output("innerHttpPort")]
        public Output<int> InnerHttpPort { get; private set; } = null!;

        /// <summary>
        /// Port number for https access over private network.
        /// </summary>
        [Output("innerHttpsPort")]
        public Output<int> InnerHttpsPort { get; private set; } = null!;

        /// <summary>
        /// Private network access subdomain name.
        /// </summary>
        [Output("internalSubDomain")]
        public Output<string> InternalSubDomain { get; private set; } = null!;

        /// <summary>
        /// IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
        /// </summary>
        [Output("ipVersion")]
        public Output<string?> IpVersion { get; private set; } = null!;

        /// <summary>
        /// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
        /// indicates access over private network, and `OUTER` indicates access over public network.
        /// </summary>
        [Output("netTypes")]
        public Output<ImmutableArray<string>> NetTypes { get; private set; } = null!;

        /// <summary>
        /// Public network access subdomain name.
        /// </summary>
        [Output("outerSubDomain")]
        public Output<string> OuterSubDomain { get; private set; } = null!;

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Output("preLimit")]
        public Output<int> PreLimit { get; private set; } = null!;

        /// <summary>
        /// Service frontend request type. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Output("releaseLimit")]
        public Output<int> ReleaseLimit { get; private set; } = null!;

        /// <summary>
        /// Custom service description.
        /// </summary>
        [Output("serviceDesc")]
        public Output<string?> ServiceDesc { get; private set; } = null!;

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Output("testLimit")]
        public Output<int> TestLimit { get; private set; } = null!;

        /// <summary>
        /// A list of attach usage plans.
        /// </summary>
        [Output("usagePlanLists")]
        public Output<ImmutableArray<Outputs.ServiceUsagePlanList>> UsagePlanLists { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:APIGateway/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:APIGateway/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
        /// </summary>
        [Input("exclusiveSetName")]
        public Input<string>? ExclusiveSetName { get; set; }

        /// <summary>
        /// IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        [Input("netTypes", required: true)]
        private InputList<string>? _netTypes;

        /// <summary>
        /// Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
        /// indicates access over private network, and `OUTER` indicates access over public network.
        /// </summary>
        public InputList<string> NetTypes
        {
            get => _netTypes ?? (_netTypes = new InputList<string>());
            set => _netTypes = value;
        }

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("preLimit")]
        public Input<int>? PreLimit { get; set; }

        /// <summary>
        /// Service frontend request type. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("releaseLimit")]
        public Input<int>? ReleaseLimit { get; set; }

        /// <summary>
        /// Custom service description.
        /// </summary>
        [Input("serviceDesc")]
        public Input<string>? ServiceDesc { get; set; }

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("testLimit")]
        public Input<int>? TestLimit { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("apiLists")]
        private InputList<Inputs.ServiceApiListGetArgs>? _apiLists;

        /// <summary>
        /// A list of APIs.
        /// </summary>
        public InputList<Inputs.ServiceApiListGetArgs> ApiLists
        {
            get => _apiLists ?? (_apiLists = new InputList<Inputs.ServiceApiListGetArgs>());
            set => _apiLists = value;
        }

        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
        /// </summary>
        [Input("exclusiveSetName")]
        public Input<string>? ExclusiveSetName { get; set; }

        /// <summary>
        /// Port number for http access over private network.
        /// </summary>
        [Input("innerHttpPort")]
        public Input<int>? InnerHttpPort { get; set; }

        /// <summary>
        /// Port number for https access over private network.
        /// </summary>
        [Input("innerHttpsPort")]
        public Input<int>? InnerHttpsPort { get; set; }

        /// <summary>
        /// Private network access subdomain name.
        /// </summary>
        [Input("internalSubDomain")]
        public Input<string>? InternalSubDomain { get; set; }

        /// <summary>
        /// IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        [Input("netTypes")]
        private InputList<string>? _netTypes;

        /// <summary>
        /// Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
        /// indicates access over private network, and `OUTER` indicates access over public network.
        /// </summary>
        public InputList<string> NetTypes
        {
            get => _netTypes ?? (_netTypes = new InputList<string>());
            set => _netTypes = value;
        }

        /// <summary>
        /// Public network access subdomain name.
        /// </summary>
        [Input("outerSubDomain")]
        public Input<string>? OuterSubDomain { get; set; }

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("preLimit")]
        public Input<int>? PreLimit { get; set; }

        /// <summary>
        /// Service frontend request type. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("releaseLimit")]
        public Input<int>? ReleaseLimit { get; set; }

        /// <summary>
        /// Custom service description.
        /// </summary>
        [Input("serviceDesc")]
        public Input<string>? ServiceDesc { get; set; }

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
        /// </summary>
        [Input("testLimit")]
        public Input<int>? TestLimit { get; set; }

        [Input("usagePlanLists")]
        private InputList<Inputs.ServiceUsagePlanListGetArgs>? _usagePlanLists;

        /// <summary>
        /// A list of attach usage plans.
        /// </summary>
        public InputList<Inputs.ServiceUsagePlanListGetArgs> UsagePlanLists
        {
            get => _usagePlanLists ?? (_usagePlanLists = new InputList<Inputs.ServiceUsagePlanListGetArgs>());
            set => _usagePlanLists = value;
        }

        public ServiceState()
        {
        }
    }
}
