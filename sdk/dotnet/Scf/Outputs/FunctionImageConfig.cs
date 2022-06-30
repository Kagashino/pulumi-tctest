// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class FunctionImageConfig
    {
        public readonly string? Args;
        public readonly string? Command;
        public readonly string? EntryPoint;
        public readonly string ImageType;
        public readonly string ImageUri;
        public readonly string? RegistryId;

        [OutputConstructor]
        private FunctionImageConfig(
            string? args,

            string? command,

            string? entryPoint,

            string imageType,

            string imageUri,

            string? registryId)
        {
            Args = args;
            Command = command;
            EntryPoint = entryPoint;
            ImageType = imageType;
            ImageUri = imageUri;
            RegistryId = registryId;
        }
    }
}
