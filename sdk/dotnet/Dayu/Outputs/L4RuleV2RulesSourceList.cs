// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Dayu.Outputs
{

    [OutputType]
    public sealed class L4RuleV2RulesSourceList
    {
        public readonly string Source;
        public readonly int Weight;

        [OutputConstructor]
        private L4RuleV2RulesSourceList(
            string source,

            int weight)
        {
            Source = source;
            Weight = weight;
        }
    }
}
