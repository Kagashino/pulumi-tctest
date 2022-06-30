// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainsDomainListHttpsConfigResult
    {
        public readonly string Http2Switch;
        public readonly string HttpsSwitch;
        public readonly string OcspStaplingSwitch;
        public readonly string SpdySwitch;
        public readonly string VerifyClient;

        [OutputConstructor]
        private DomainsDomainListHttpsConfigResult(
            string http2Switch,

            string httpsSwitch,

            string ocspStaplingSwitch,

            string spdySwitch,

            string verifyClient)
        {
            Http2Switch = http2Switch;
            HttpsSwitch = httpsSwitch;
            OcspStaplingSwitch = ocspStaplingSwitch;
            SpdySwitch = spdySwitch;
            VerifyClient = verifyClient;
        }
    }
}
