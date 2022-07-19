// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cloud.Outputs
{

    [OutputType]
    public sealed class TCRRepositoriesRepositoryListResult
    {
        public readonly string BriefDesc;
        public readonly string CreateTime;
        public readonly string Description;
        public readonly bool IsPublic;
        public readonly string Name;
        public readonly string NamespaceName;
        public readonly string UpdateTime;
        public readonly string Url;

        [OutputConstructor]
        private TCRRepositoriesRepositoryListResult(
            string briefDesc,

            string createTime,

            string description,

            bool isPublic,

            string name,

            string namespaceName,

            string updateTime,

            string url)
        {
            BriefDesc = briefDesc;
            CreateTime = createTime;
            Description = description;
            IsPublic = isPublic;
            Name = name;
            NamespaceName = namespaceName;
            UpdateTime = updateTime;
            Url = url;
        }
    }
}
