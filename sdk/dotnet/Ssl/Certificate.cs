// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    [TencentcloudResourceType("tencentcloud:Ssl/certificate:Certificate")]
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// Beginning time of the SSL certificate.
        /// </summary>
        [Output("beginTime")]
        public Output<string> BeginTime { get; private set; } = null!;

        /// <summary>
        /// Content of the SSL certificate. Not allowed newline at the start and end.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Creation time of the SSL certificate.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Primary domain of the SSL certificate.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Ending time of the SSL certificate.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// Name of the SSL certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Certificate authority.
        /// </summary>
        [Output("productZhName")]
        public Output<string> ProductZhName { get; private set; } = null!;

        /// <summary>
        /// Project ID of the SSL certificate. Default is `0`.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Status of the SSL certificate.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// ALL domains included in the SSL certificate. Including the primary domain name.
        /// </summary>
        [Output("subjectNames")]
        public Output<ImmutableArray<string>> SubjectNames { get; private set; } = null!;

        /// <summary>
        /// Type of the SSL certificate. Valid values: `CA` and `SVR`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/certificate:Certificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content of the SSL certificate. Not allowed newline at the start and end.
        /// </summary>
        [Input("cert", required: true)]
        public Input<string> Cert { get; set; } = null!;

        /// <summary>
        /// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Name of the SSL certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project ID of the SSL certificate. Default is `0`.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Type of the SSL certificate. Valid values: `CA` and `SVR`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Beginning time of the SSL certificate.
        /// </summary>
        [Input("beginTime")]
        public Input<string>? BeginTime { get; set; }

        /// <summary>
        /// Content of the SSL certificate. Not allowed newline at the start and end.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Creation time of the SSL certificate.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Primary domain of the SSL certificate.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Ending time of the SSL certificate.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Name of the SSL certificate.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Certificate authority.
        /// </summary>
        [Input("productZhName")]
        public Input<string>? ProductZhName { get; set; }

        /// <summary>
        /// Project ID of the SSL certificate. Default is `0`.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Status of the SSL certificate.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("subjectNames")]
        private InputList<string>? _subjectNames;

        /// <summary>
        /// ALL domains included in the SSL certificate. Including the primary domain name.
        /// </summary>
        public InputList<string> SubjectNames
        {
            get => _subjectNames ?? (_subjectNames = new InputList<string>());
            set => _subjectNames = value;
        }

        /// <summary>
        /// Type of the SSL certificate. Valid values: `CA` and `SVR`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateState()
        {
        }
    }
}
