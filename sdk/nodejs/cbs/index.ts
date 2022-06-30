// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./snapshot";
export * from "./snapshotPolicies";
export * from "./snapshotPolicy";
export * from "./snapshotPolicyAttachment";
export * from "./snapshots";
export * from "./storage";
export * from "./storageAttachment";
export * from "./storages";

// Import resources to register:
import { Snapshot } from "./snapshot";
import { SnapshotPolicy } from "./snapshotPolicy";
import { SnapshotPolicyAttachment } from "./snapshotPolicyAttachment";
import { Storage } from "./storage";
import { StorageAttachment } from "./storageAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cbs/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy":
                return new SnapshotPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment":
                return new SnapshotPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storage:Storage":
                return new Storage(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storageAttachment:StorageAttachment":
                return new StorageAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshot", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storage", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageAttachment", _module)