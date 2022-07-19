// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cbs
{
    public static class SnapshotPolicies
    {
        public static Task<SnapshotPoliciesResult> InvokeAsync(SnapshotPoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SnapshotPoliciesResult>("tctest:Cbs/snapshotPolicies:SnapshotPolicies", args ?? new SnapshotPoliciesArgs(), options.WithDefaults());

        public static Output<SnapshotPoliciesResult> Invoke(SnapshotPoliciesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SnapshotPoliciesResult>("tctest:Cbs/snapshotPolicies:SnapshotPolicies", args ?? new SnapshotPoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class SnapshotPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("snapshotPolicyId")]
        public string? SnapshotPolicyId { get; set; }

        [Input("snapshotPolicyName")]
        public string? SnapshotPolicyName { get; set; }

        public SnapshotPoliciesArgs()
        {
        }
    }

    public sealed class SnapshotPoliciesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("snapshotPolicyId")]
        public Input<string>? SnapshotPolicyId { get; set; }

        [Input("snapshotPolicyName")]
        public Input<string>? SnapshotPolicyName { get; set; }

        public SnapshotPoliciesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SnapshotPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string? SnapshotPolicyId;
        public readonly ImmutableArray<Outputs.SnapshotPoliciesSnapshotPolicyListResult> SnapshotPolicyLists;
        public readonly string? SnapshotPolicyName;

        [OutputConstructor]
        private SnapshotPoliciesResult(
            string id,

            string? resultOutputFile,

            string? snapshotPolicyId,

            ImmutableArray<Outputs.SnapshotPoliciesSnapshotPolicyListResult> snapshotPolicyLists,

            string? snapshotPolicyName)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SnapshotPolicyId = snapshotPolicyId;
            SnapshotPolicyLists = snapshotPolicyLists;
            SnapshotPolicyName = snapshotPolicyName;
        }
    }
}
