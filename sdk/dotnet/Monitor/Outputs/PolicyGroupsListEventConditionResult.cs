// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class PolicyGroupsListEventConditionResult
    {
        public readonly int AlarmNotifyPeriod;
        public readonly int AlarmNotifyType;
        public readonly int EventId;
        public readonly string EventShowName;
        public readonly int RuleId;

        [OutputConstructor]
        private PolicyGroupsListEventConditionResult(
            int alarmNotifyPeriod,

            int alarmNotifyType,

            int eventId,

            string eventShowName,

            int ruleId)
        {
            AlarmNotifyPeriod = alarmNotifyPeriod;
            AlarmNotifyType = alarmNotifyType;
            EventId = eventId;
            EventShowName = eventShowName;
            RuleId = ruleId;
        }
    }
}
