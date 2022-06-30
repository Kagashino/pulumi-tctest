// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./record";
export * from "./zone";

// Import resources to register:
import { Record } from "./record";
import { Zone } from "./zone";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:PrivateDns/record:Record":
                return new Record(name, <any>undefined, { urn })
            case "tencentcloud:PrivateDns/zone:Zone":
                return new Zone(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "PrivateDns/record", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "PrivateDns/zone", _module)
