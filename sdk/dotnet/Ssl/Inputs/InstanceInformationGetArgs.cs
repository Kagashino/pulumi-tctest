// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl.Inputs
{

    public sealed class InstanceInformationGetArgs : Pulumi.ResourceArgs
    {
        [Input("adminEmail", required: true)]
        public Input<string> AdminEmail { get; set; } = null!;

        [Input("adminFirstName", required: true)]
        public Input<string> AdminFirstName { get; set; } = null!;

        [Input("adminLastName", required: true)]
        public Input<string> AdminLastName { get; set; } = null!;

        [Input("adminPhoneNum", required: true)]
        public Input<string> AdminPhoneNum { get; set; } = null!;

        [Input("adminPosition", required: true)]
        public Input<string> AdminPosition { get; set; } = null!;

        [Input("certificateDomain", required: true)]
        public Input<string> CertificateDomain { get; set; } = null!;

        [Input("contactEmail", required: true)]
        public Input<string> ContactEmail { get; set; } = null!;

        [Input("contactFirstName", required: true)]
        public Input<string> ContactFirstName { get; set; } = null!;

        [Input("contactLastName", required: true)]
        public Input<string> ContactLastName { get; set; } = null!;

        [Input("contactNumber", required: true)]
        public Input<string> ContactNumber { get; set; } = null!;

        [Input("contactPosition", required: true)]
        public Input<string> ContactPosition { get; set; } = null!;

        [Input("csrContent")]
        public Input<string>? CsrContent { get; set; }

        [Input("csrType")]
        public Input<string>? CsrType { get; set; }

        [Input("domainLists")]
        private InputList<string>? _domainLists;
        public InputList<string> DomainLists
        {
            get => _domainLists ?? (_domainLists = new InputList<string>());
            set => _domainLists = value;
        }

        [Input("keyPassword")]
        public Input<string>? KeyPassword { get; set; }

        [Input("organizationAddress", required: true)]
        public Input<string> OrganizationAddress { get; set; } = null!;

        [Input("organizationCity", required: true)]
        public Input<string> OrganizationCity { get; set; } = null!;

        [Input("organizationCountry", required: true)]
        public Input<string> OrganizationCountry { get; set; } = null!;

        [Input("organizationDivision", required: true)]
        public Input<string> OrganizationDivision { get; set; } = null!;

        [Input("organizationName", required: true)]
        public Input<string> OrganizationName { get; set; } = null!;

        [Input("organizationRegion", required: true)]
        public Input<string> OrganizationRegion { get; set; } = null!;

        [Input("phoneAreaCode", required: true)]
        public Input<string> PhoneAreaCode { get; set; } = null!;

        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        [Input("postalCode", required: true)]
        public Input<string> PostalCode { get; set; } = null!;

        [Input("verifyType", required: true)]
        public Input<string> VerifyType { get; set; } = null!;

        public InstanceInformationGetArgs()
        {
        }
    }
}