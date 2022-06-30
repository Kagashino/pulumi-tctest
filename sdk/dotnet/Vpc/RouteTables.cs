// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc
{
    public static class RouteTables
    {
        public static Task<RouteTablesResult> InvokeAsync(RouteTablesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<RouteTablesResult>("tencentcloud:Vpc/routeTables:RouteTables", args ?? new RouteTablesArgs(), options.WithDefaults());

        public static Output<RouteTablesResult> Invoke(RouteTablesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<RouteTablesResult>("tencentcloud:Vpc/routeTables:RouteTables", args ?? new RouteTablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class RouteTablesArgs : Pulumi.InvokeArgs
    {
        [Input("associationMain")]
        public bool? AssociationMain { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("routeTableId")]
        public string? RouteTableId { get; set; }

        [Input("tagKey")]
        public string? TagKey { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public RouteTablesArgs()
        {
        }
    }

    public sealed class RouteTablesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("associationMain")]
        public Input<bool>? AssociationMain { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public RouteTablesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class RouteTablesResult
    {
        public readonly bool? AssociationMain;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.RouteTablesInstanceListResult> InstanceLists;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly string? RouteTableId;
        public readonly string? TagKey;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? VpcId;

        [OutputConstructor]
        private RouteTablesResult(
            bool? associationMain,

            string id,

            ImmutableArray<Outputs.RouteTablesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile,

            string? routeTableId,

            string? tagKey,

            ImmutableDictionary<string, object>? tags,

            string? vpcId)
        {
            AssociationMain = associationMain;
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
            RouteTableId = routeTableId;
            TagKey = tagKey;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
