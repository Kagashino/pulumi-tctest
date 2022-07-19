// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function templateGroups(args?: TemplateGroupsArgs, opts?: pulumi.InvokeOptions): Promise<TemplateGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Protocol/templateGroups:TemplateGroups", {
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking TemplateGroups.
 */
export interface TemplateGroupsArgs {
    id?: string;
    name?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by TemplateGroups.
 */
export interface TemplateGroupsResult {
    readonly groupLists: outputs.Protocol.TemplateGroupsGroupList[];
    readonly id?: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
}

export function templateGroupsOutput(args?: TemplateGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TemplateGroupsResult> {
    return pulumi.output(args).apply(a => templateGroups(a, opts))
}

/**
 * A collection of arguments for invoking TemplateGroups.
 */
export interface TemplateGroupsOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
