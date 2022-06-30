// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./domainInstance";
export * from "./record";

// Import resources to register:
import { DomainInstance } from "./domainInstance";
import { Record } from "./record";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Dnspod/domainInstance:DomainInstance":
                return new DomainInstance(name, <any>undefined, { urn })
            case "tencentcloud:Dnspod/record:Record":
                return new Record(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Dnspod/domainInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dnspod/record", _module)
