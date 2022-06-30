// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cam
{
    public static class Roles
    {
        public static Task<RolesResult> InvokeAsync(RolesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<RolesResult>("tencentcloud:Cam/roles:Roles", args ?? new RolesArgs(), options.WithDefaults());

        public static Output<RolesResult> Invoke(RolesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<RolesResult>("tencentcloud:Cam/roles:Roles", args ?? new RolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class RolesArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("roleId")]
        public string? RoleId { get; set; }

        public RolesArgs()
        {
        }
    }

    public sealed class RolesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public RolesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class RolesResult
    {
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly string? RoleId;
        public readonly ImmutableArray<Outputs.RolesRoleListResult> RoleLists;

        [OutputConstructor]
        private RolesResult(
            string? description,

            string id,

            string? name,

            string? resultOutputFile,

            string? roleId,

            ImmutableArray<Outputs.RolesRoleListResult> roleLists)
        {
            Description = description;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
            RoleId = roleId;
            RoleLists = roleLists;
        }
    }
}