// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cos.Outputs
{

    [OutputType]
    public sealed class BucketsBucketListLifecycleRuleNonCurrentTransitionResult
    {
        public readonly int NonCurrentDays;
        public readonly string StorageClass;

        [OutputConstructor]
        private BucketsBucketListLifecycleRuleNonCurrentTransitionResult(
            int nonCurrentDays,

            string storageClass)
        {
            NonCurrentDays = nonCurrentDays;
            StorageClass = storageClass;
        }
    }
}
