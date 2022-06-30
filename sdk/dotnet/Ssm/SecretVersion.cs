// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssm
{
    [TencentcloudResourceType("tencentcloud:Ssm/secretVersion:SecretVersion")]
    public partial class SecretVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The base64-encoded binary secret. secret_binary and secret_string must be set only one, and the maximum support is 4096
        /// bytes. When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Output("secretBinary")]
        public Output<string?> SecretBinary { get; private set; } = null!;

        /// <summary>
        /// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain
        /// English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Output("secretName")]
        public Output<string> SecretName { get; private set; } = null!;

        /// <summary>
        /// The string text of secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes.
        /// When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Output("secretString")]
        public Output<string?> SecretString { get; private set; } = null!;

        /// <summary>
        /// Version of secret. The maximum length is 64 bytes. The version_id can only contain English letters, numbers, underscore
        /// and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;


        /// <summary>
        /// Create a SecretVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretVersion(string name, SecretVersionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssm/secretVersion:SecretVersion", name, args ?? new SecretVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretVersion(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssm/secretVersion:SecretVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretVersion Get(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretVersion(name, id, state, options);
        }
    }

    public sealed class SecretVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64-encoded binary secret. secret_binary and secret_string must be set only one, and the maximum support is 4096
        /// bytes. When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Input("secretBinary")]
        public Input<string>? SecretBinary { get; set; }

        /// <summary>
        /// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain
        /// English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        /// <summary>
        /// The string text of secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes.
        /// When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Input("secretString")]
        public Input<string>? SecretString { get; set; }

        /// <summary>
        /// Version of secret. The maximum length is 64 bytes. The version_id can only contain English letters, numbers, underscore
        /// and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public SecretVersionArgs()
        {
        }
    }

    public sealed class SecretVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64-encoded binary secret. secret_binary and secret_string must be set only one, and the maximum support is 4096
        /// bytes. When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Input("secretBinary")]
        public Input<string>? SecretBinary { get; set; }

        /// <summary>
        /// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain
        /// English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        /// <summary>
        /// The string text of secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes.
        /// When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        [Input("secretString")]
        public Input<string>? SecretString { get; set; }

        /// <summary>
        /// Version of secret. The maximum length is 64 bytes. The version_id can only contain English letters, numbers, underscore
        /// and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public SecretVersionState()
        {
        }
    }
}
