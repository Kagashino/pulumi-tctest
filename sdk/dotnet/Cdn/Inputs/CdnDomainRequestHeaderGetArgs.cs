// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class CdnDomainRequestHeaderGetArgs : Pulumi.ResourceArgs
    {
        [Input("headerRules")]
        private InputList<Inputs.CdnDomainRequestHeaderHeaderRuleGetArgs>? _headerRules;
        public InputList<Inputs.CdnDomainRequestHeaderHeaderRuleGetArgs> HeaderRules
        {
            get => _headerRules ?? (_headerRules = new InputList<Inputs.CdnDomainRequestHeaderHeaderRuleGetArgs>());
            set => _headerRules = value;
        }

        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public CdnDomainRequestHeaderGetArgs()
        {
        }
    }
}
