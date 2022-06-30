// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function users(args: UsersArgs, opts?: pulumi.InvokeOptions): Promise<UsersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ckafka/users:Users", {
        "accountName": args.accountName,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Users.
 */
export interface UsersArgs {
    accountName?: string;
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by Users.
 */
export interface UsersResult {
    readonly accountName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
    readonly userLists: outputs.Ckafka.UsersUserList[];
}

export function usersOutput(args: UsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<UsersResult> {
    return pulumi.output(args).apply(a => users(a, opts))
}

/**
 * A collection of arguments for invoking Users.
 */
export interface UsersOutputArgs {
    accountName?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
