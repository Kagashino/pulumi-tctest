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
    public sealed class CCPolicyV2CcPrecisionPolicyPolicy
    {
        public readonly string FieldName;
        public readonly string FieldType;
        public readonly string Value;
        public readonly string ValueOperator;

        [OutputConstructor]
        private CCPolicyV2CcPrecisionPolicyPolicy(
            string fieldName,

            string fieldType,

            string value,

            string valueOperator)
        {
            FieldName = fieldName;
            FieldType = fieldType;
            Value = value;
            ValueOperator = valueOperator;
        }
    }
}
