// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Inputs
{

    public sealed class ConfigExtraExtractRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("backtracking")]
        public Input<int>? Backtracking { get; set; }

        [Input("beginRegex")]
        public Input<string>? BeginRegex { get; set; }

        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("filterKeyRegexes")]
        private InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexGetArgs>? _filterKeyRegexes;
        public InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexGetArgs> FilterKeyRegexes
        {
            get => _filterKeyRegexes ?? (_filterKeyRegexes = new InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexGetArgs>());
            set => _filterKeyRegexes = value;
        }

        [Input("keys")]
        private InputList<string>? _keys;
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        [Input("logRegex")]
        public Input<string>? LogRegex { get; set; }

        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        [Input("timeKey")]
        public Input<string>? TimeKey { get; set; }

        [Input("unMatchLogKey")]
        public Input<string>? UnMatchLogKey { get; set; }

        [Input("unMatchUpLoadSwitch")]
        public Input<bool>? UnMatchUpLoadSwitch { get; set; }

        public ConfigExtraExtractRuleGetArgs()
        {
        }
    }
}
