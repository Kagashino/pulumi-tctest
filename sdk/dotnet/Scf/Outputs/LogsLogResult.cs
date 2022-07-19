// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Scf.Outputs
{

    [OutputType]
    public sealed class LogsLogResult
    {
        public readonly int BillDuration;
        public readonly double Duration;
        public readonly string FunctionName;
        public readonly int InvokeFinished;
        public readonly string Level;
        public readonly string Log;
        public readonly int MemUsage;
        public readonly string RequestId;
        public readonly int RetCode;
        public readonly string RetMsg;
        public readonly string Source;
        public readonly string StartTime;

        [OutputConstructor]
        private LogsLogResult(
            int billDuration,

            double duration,

            string functionName,

            int invokeFinished,

            string level,

            string log,

            int memUsage,

            string requestId,

            int retCode,

            string retMsg,

            string source,

            string startTime)
        {
            BillDuration = billDuration;
            Duration = duration;
            FunctionName = functionName;
            InvokeFinished = invokeFinished;
            Level = level;
            Log = log;
            MemUsage = memUsage;
            RequestId = requestId;
            RetCode = retCode;
            RetMsg = retMsg;
            Source = source;
            StartTime = startTime;
        }
    }
}
