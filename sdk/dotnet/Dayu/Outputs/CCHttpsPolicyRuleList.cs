// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class CCHttpsPolicyRuleList
    {
        public readonly string Operator;
        public readonly string Skey;
        public readonly string Value;

        [OutputConstructor]
        private CCHttpsPolicyRuleList(
            string @operator,

            string skey,

            string value)
        {
            Operator = @operator;
            Skey = skey;
            Value = value;
        }
    }
}
