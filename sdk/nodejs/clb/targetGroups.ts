// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function targetGroups(args?: TargetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<TargetGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Clb/targetGroups:TargetGroups", {
        "resultOutputFile": args.resultOutputFile,
        "targetGroupId": args.targetGroupId,
        "targetGroupName": args.targetGroupName,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking TargetGroups.
 */
export interface TargetGroupsArgs {
    resultOutputFile?: string;
    targetGroupId?: string;
    targetGroupName?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by TargetGroups.
 */
export interface TargetGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Clb.TargetGroupsList[];
    readonly resultOutputFile?: string;
    readonly targetGroupId?: string;
    readonly targetGroupName?: string;
    readonly vpcId?: string;
}

export function targetGroupsOutput(args?: TargetGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TargetGroupsResult> {
    return pulumi.output(args).apply(a => targetGroups(a, opts))
}

/**
 * A collection of arguments for invoking TargetGroups.
 */
export interface TargetGroupsOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    targetGroupId?: pulumi.Input<string>;
    targetGroupName?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
