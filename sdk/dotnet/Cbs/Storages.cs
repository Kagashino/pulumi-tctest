// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cbs
{
    public static class Storages
    {
        public static Task<StoragesResult> InvokeAsync(StoragesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<StoragesResult>("tencentcloud:Cbs/storages:Storages", args ?? new StoragesArgs(), options.WithDefaults());

        public static Output<StoragesResult> Invoke(StoragesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<StoragesResult>("tencentcloud:Cbs/storages:Storages", args ?? new StoragesInvokeArgs(), options.WithDefaults());
    }


    public sealed class StoragesArgs : Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("chargeTypes")]
        private List<string>? _chargeTypes;
        public List<string> ChargeTypes
        {
            get => _chargeTypes ?? (_chargeTypes = new List<string>());
            set => _chargeTypes = value;
        }

        [Input("instanceIps")]
        private List<string>? _instanceIps;
        public List<string> InstanceIps
        {
            get => _instanceIps ?? (_instanceIps = new List<string>());
            set => _instanceIps = value;
        }

        [Input("instanceNames")]
        private List<string>? _instanceNames;
        public List<string> InstanceNames
        {
            get => _instanceNames ?? (_instanceNames = new List<string>());
            set => _instanceNames = value;
        }

        [Input("portable")]
        public bool? Portable { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("storageId")]
        public string? StorageId { get; set; }

        [Input("storageName")]
        public string? StorageName { get; set; }

        [Input("storageStates")]
        private List<string>? _storageStates;
        public List<string> StorageStates
        {
            get => _storageStates ?? (_storageStates = new List<string>());
            set => _storageStates = value;
        }

        [Input("storageType")]
        public string? StorageType { get; set; }

        [Input("storageUsage")]
        public string? StorageUsage { get; set; }

        [Input("tagKeys")]
        private List<string>? _tagKeys;
        public List<string> TagKeys
        {
            get => _tagKeys ?? (_tagKeys = new List<string>());
            set => _tagKeys = value;
        }

        [Input("tagValues")]
        private List<string>? _tagValues;
        public List<string> TagValues
        {
            get => _tagValues ?? (_tagValues = new List<string>());
            set => _tagValues = value;
        }

        public StoragesArgs()
        {
        }
    }

    public sealed class StoragesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("chargeTypes")]
        private InputList<string>? _chargeTypes;
        public InputList<string> ChargeTypes
        {
            get => _chargeTypes ?? (_chargeTypes = new InputList<string>());
            set => _chargeTypes = value;
        }

        [Input("instanceIps")]
        private InputList<string>? _instanceIps;
        public InputList<string> InstanceIps
        {
            get => _instanceIps ?? (_instanceIps = new InputList<string>());
            set => _instanceIps = value;
        }

        [Input("instanceNames")]
        private InputList<string>? _instanceNames;
        public InputList<string> InstanceNames
        {
            get => _instanceNames ?? (_instanceNames = new InputList<string>());
            set => _instanceNames = value;
        }

        [Input("portable")]
        public Input<bool>? Portable { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("storageId")]
        public Input<string>? StorageId { get; set; }

        [Input("storageName")]
        public Input<string>? StorageName { get; set; }

        [Input("storageStates")]
        private InputList<string>? _storageStates;
        public InputList<string> StorageStates
        {
            get => _storageStates ?? (_storageStates = new InputList<string>());
            set => _storageStates = value;
        }

        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("storageUsage")]
        public Input<string>? StorageUsage { get; set; }

        [Input("tagKeys")]
        private InputList<string>? _tagKeys;
        public InputList<string> TagKeys
        {
            get => _tagKeys ?? (_tagKeys = new InputList<string>());
            set => _tagKeys = value;
        }

        [Input("tagValues")]
        private InputList<string>? _tagValues;
        public InputList<string> TagValues
        {
            get => _tagValues ?? (_tagValues = new InputList<string>());
            set => _tagValues = value;
        }

        public StoragesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class StoragesResult
    {
        public readonly string? AvailabilityZone;
        public readonly ImmutableArray<string> ChargeTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIps;
        public readonly ImmutableArray<string> InstanceNames;
        public readonly bool? Portable;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;
        public readonly string? StorageId;
        public readonly ImmutableArray<Outputs.StoragesStorageListResult> StorageLists;
        public readonly string? StorageName;
        public readonly ImmutableArray<string> StorageStates;
        public readonly string? StorageType;
        public readonly string? StorageUsage;
        public readonly ImmutableArray<string> TagKeys;
        public readonly ImmutableArray<string> TagValues;

        [OutputConstructor]
        private StoragesResult(
            string? availabilityZone,

            ImmutableArray<string> chargeTypes,

            string id,

            ImmutableArray<string> instanceIps,

            ImmutableArray<string> instanceNames,

            bool? portable,

            int? projectId,

            string? resultOutputFile,

            string? storageId,

            ImmutableArray<Outputs.StoragesStorageListResult> storageLists,

            string? storageName,

            ImmutableArray<string> storageStates,

            string? storageType,

            string? storageUsage,

            ImmutableArray<string> tagKeys,

            ImmutableArray<string> tagValues)
        {
            AvailabilityZone = availabilityZone;
            ChargeTypes = chargeTypes;
            Id = id;
            InstanceIps = instanceIps;
            InstanceNames = instanceNames;
            Portable = portable;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
            StorageId = storageId;
            StorageLists = storageLists;
            StorageName = storageName;
            StorageStates = storageStates;
            StorageType = storageType;
            StorageUsage = storageUsage;
            TagKeys = tagKeys;
            TagValues = tagValues;
        }
    }
}
