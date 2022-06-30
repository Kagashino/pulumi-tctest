// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tke.Outputs
{

    [OutputType]
    public sealed class ClusterEventPersistence
    {
        public readonly bool Enabled;
        public readonly string? LogSetId;
        public readonly string? TopicId;

        [OutputConstructor]
        private ClusterEventPersistence(
            bool enabled,

            string? logSetId,

            string? topicId)
        {
            Enabled = enabled;
            LogSetId = logSetId;
            TopicId = topicId;
        }
    }
}
