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
    public sealed class ProductEventListResult
    {
        public readonly ImmutableArray<Outputs.ProductEventListAdditionMsgResult> AdditionMsgs;
        public readonly ImmutableArray<Outputs.ProductEventListDimensionResult> Dimensions;
        public readonly string EventCname;
        public readonly string EventEname;
        public readonly int EventId;
        public readonly string EventName;
        public readonly ImmutableArray<Outputs.ProductEventListGroupInfoResult> GroupInfos;
        public readonly string InstanceId;
        public readonly string InstanceName;
        public readonly int IsAlarmConfig;
        public readonly string ProductCname;
        public readonly string ProductEname;
        public readonly string ProductName;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly int StartTime;
        public readonly string Status;
        public readonly int SupportAlarm;
        public readonly string Type;
        public readonly int UpdateTime;

        [OutputConstructor]
        private ProductEventListResult(
            ImmutableArray<Outputs.ProductEventListAdditionMsgResult> additionMsgs,

            ImmutableArray<Outputs.ProductEventListDimensionResult> dimensions,

            string eventCname,

            string eventEname,

            int eventId,

            string eventName,

            ImmutableArray<Outputs.ProductEventListGroupInfoResult> groupInfos,

            string instanceId,

            string instanceName,

            int isAlarmConfig,

            string productCname,

            string productEname,

            string productName,

            string projectId,

            string region,

            int startTime,

            string status,

            int supportAlarm,

            string type,

            int updateTime)
        {
            AdditionMsgs = additionMsgs;
            Dimensions = dimensions;
            EventCname = eventCname;
            EventEname = eventEname;
            EventId = eventId;
            EventName = eventName;
            GroupInfos = groupInfos;
            InstanceId = instanceId;
            InstanceName = instanceName;
            IsAlarmConfig = isAlarmConfig;
            ProductCname = productCname;
            ProductEname = productEname;
            ProductName = productName;
            ProjectId = projectId;
            Region = region;
            StartTime = startTime;
            Status = status;
            SupportAlarm = supportAlarm;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
