// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tke.Inputs
{

    public sealed class ClusterExistInstanceInstancesParaGetArgs : Pulumi.ResourceArgs
    {
        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        public ClusterExistInstanceInstancesParaGetArgs()
        {
        }
    }
}
