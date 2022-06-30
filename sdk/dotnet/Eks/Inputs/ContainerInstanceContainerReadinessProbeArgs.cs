// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eks.Inputs
{

    public sealed class ContainerInstanceContainerReadinessProbeArgs : Pulumi.ResourceArgs
    {
        [Input("execCommands")]
        private InputList<string>? _execCommands;
        public InputList<string> ExecCommands
        {
            get => _execCommands ?? (_execCommands = new InputList<string>());
            set => _execCommands = value;
        }

        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        [Input("httpGetPath")]
        public Input<string>? HttpGetPath { get; set; }

        [Input("httpGetPort")]
        public Input<int>? HttpGetPort { get; set; }

        [Input("httpGetScheme")]
        public Input<string>? HttpGetScheme { get; set; }

        [Input("initDelaySeconds")]
        public Input<int>? InitDelaySeconds { get; set; }

        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        [Input("tcpSocketPort")]
        public Input<int>? TcpSocketPort { get; set; }

        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        public ContainerInstanceContainerReadinessProbeArgs()
        {
        }
    }
}