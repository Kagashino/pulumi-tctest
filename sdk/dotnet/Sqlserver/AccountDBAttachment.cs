// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    [TencentcloudResourceType("tencentcloud:Sqlserver/accountDBAttachment:AccountDBAttachment")]
    public partial class AccountDBAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// SQL Server account name.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// SQL Server DB name.
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// SQL Server instance ID that the account belongs to.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
        /// </summary>
        [Output("privilege")]
        public Output<string> Privilege { get; private set; } = null!;


        /// <summary>
        /// Create a AccountDBAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountDBAttachment(string name, AccountDBAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/accountDBAttachment:AccountDBAttachment", name, args ?? new AccountDBAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountDBAttachment(string name, Input<string> id, AccountDBAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/accountDBAttachment:AccountDBAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountDBAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountDBAttachment Get(string name, Input<string> id, AccountDBAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountDBAttachment(name, id, state, options);
        }
    }

    public sealed class AccountDBAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SQL Server account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// SQL Server DB name.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// SQL Server instance ID that the account belongs to.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
        /// </summary>
        [Input("privilege", required: true)]
        public Input<string> Privilege { get; set; } = null!;

        public AccountDBAttachmentArgs()
        {
        }
    }

    public sealed class AccountDBAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SQL Server account name.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// SQL Server DB name.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// SQL Server instance ID that the account belongs to.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
        /// </summary>
        [Input("privilege")]
        public Input<string>? Privilege { get; set; }

        public AccountDBAttachmentState()
        {
        }
    }
}