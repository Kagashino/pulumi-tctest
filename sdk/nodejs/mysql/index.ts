// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./accountPrivilege";
export * from "./backupList";
export * from "./backupPolicy";
export * from "./defaultParams";
export * from "./instance";
export * from "./instances";
export * from "./parameterList";
export * from "./privilege";
export * from "./readonlyInstance";
export * from "./zoneConfig";

// Import resources to register:
import { Account } from "./account";
import { AccountPrivilege } from "./accountPrivilege";
import { BackupPolicy } from "./backupPolicy";
import { Instance } from "./instance";
import { Privilege } from "./privilege";
import { ReadonlyInstance } from "./readonlyInstance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tctest:Mysql/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "tctest:Mysql/accountPrivilege:AccountPrivilege":
                return new AccountPrivilege(name, <any>undefined, { urn })
            case "tctest:Mysql/backupPolicy:BackupPolicy":
                return new BackupPolicy(name, <any>undefined, { urn })
            case "tctest:Mysql/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tctest:Mysql/privilege:Privilege":
                return new Privilege(name, <any>undefined, { urn })
            case "tctest:Mysql/readonlyInstance:ReadonlyInstance":
                return new ReadonlyInstance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tctest", "Mysql/account", _module)
pulumi.runtime.registerResourceModule("tctest", "Mysql/accountPrivilege", _module)
pulumi.runtime.registerResourceModule("tctest", "Mysql/backupPolicy", _module)
pulumi.runtime.registerResourceModule("tctest", "Mysql/instance", _module)
pulumi.runtime.registerResourceModule("tctest", "Mysql/privilege", _module)
pulumi.runtime.registerResourceModule("tctest", "Mysql/readonlyInstance", _module)
