// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Inputs
{

    public sealed class PolicyGroupReceiverArgs : Pulumi.ResourceArgs
    {
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("needSendNotice")]
        public Input<int>? NeedSendNotice { get; set; }

        [Input("notifyWays")]
        private InputList<string>? _notifyWays;
        public InputList<string> NotifyWays
        {
            get => _notifyWays ?? (_notifyWays = new InputList<string>());
            set => _notifyWays = value;
        }

        [Input("personInterval")]
        public Input<int>? PersonInterval { get; set; }

        [Input("receiveLanguage")]
        public Input<string>? ReceiveLanguage { get; set; }

        [Input("receiverGroupLists")]
        private InputList<int>? _receiverGroupLists;
        public InputList<int> ReceiverGroupLists
        {
            get => _receiverGroupLists ?? (_receiverGroupLists = new InputList<int>());
            set => _receiverGroupLists = value;
        }

        [Input("receiverType")]
        public Input<string>? ReceiverType { get; set; }

        [Input("receiverUserLists")]
        private InputList<int>? _receiverUserLists;
        public InputList<int> ReceiverUserLists
        {
            get => _receiverUserLists ?? (_receiverUserLists = new InputList<int>());
            set => _receiverUserLists = value;
        }

        [Input("recoverNotifies")]
        private InputList<string>? _recoverNotifies;
        public InputList<string> RecoverNotifies
        {
            get => _recoverNotifies ?? (_recoverNotifies = new InputList<string>());
            set => _recoverNotifies = value;
        }

        [Input("roundInterval")]
        public Input<int>? RoundInterval { get; set; }

        [Input("roundNumber")]
        public Input<int>? RoundNumber { get; set; }

        [Input("sendFors")]
        private InputList<string>? _sendFors;
        public InputList<string> SendFors
        {
            get => _sendFors ?? (_sendFors = new InputList<string>());
            set => _sendFors = value;
        }

        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("uidLists")]
        private InputList<int>? _uidLists;
        public InputList<int> UidLists
        {
            get => _uidLists ?? (_uidLists = new InputList<int>());
            set => _uidLists = value;
        }

        public PolicyGroupReceiverArgs()
        {
        }
    }
}
