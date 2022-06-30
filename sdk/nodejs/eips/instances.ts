// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function instances(args?: InstancesArgs, opts?: pulumi.InvokeOptions): Promise<InstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Eips/instances:Instances", {
        "eipId": args.eipId,
        "eipName": args.eipName,
        "publicIp": args.publicIp,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    eipId?: string;
    eipName?: string;
    publicIp?: string;
    resultOutputFile?: string;
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    readonly eipId?: string;
    readonly eipLists: outputs.Eips.InstancesEipList[];
    readonly eipName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly publicIp?: string;
    readonly resultOutputFile?: string;
    readonly tags?: {[key: string]: any};
}

export function instancesOutput(args?: InstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstancesResult> {
    return pulumi.output(args).apply(a => instances(a, opts))
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesOutputArgs {
    eipId?: pulumi.Input<string>;
    eipName?: pulumi.Input<string>;
    publicIp?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}
