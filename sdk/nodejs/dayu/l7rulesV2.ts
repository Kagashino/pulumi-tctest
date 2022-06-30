// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function l7rulesV2(args: L7rulesV2Args, opts?: pulumi.InvokeOptions): Promise<L7rulesV2Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dayu/l7RulesV2:L7RulesV2", {
        "business": args.business,
        "domain": args.domain,
        "ip": args.ip,
        "limit": args.limit,
        "offset": args.offset,
        "protocol": args.protocol,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking L7RulesV2.
 */
export interface L7rulesV2Args {
    business: string;
    domain?: string;
    ip?: string;
    limit?: number;
    offset?: number;
    protocol?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by L7RulesV2.
 */
export interface L7rulesV2Result {
    readonly business: string;
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ip?: string;
    readonly limit?: number;
    readonly lists: outputs.Dayu.L7RulesV2List[];
    readonly offset?: number;
    readonly protocol?: string;
    readonly resultOutputFile?: string;
}

export function l7rulesV2Output(args: L7rulesV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<L7rulesV2Result> {
    return pulumi.output(args).apply(a => l7rulesV2(a, opts))
}

/**
 * A collection of arguments for invoking L7RulesV2.
 */
export interface L7rulesV2OutputArgs {
    business: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    limit?: pulumi.Input<number>;
    offset?: pulumi.Input<number>;
    protocol?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
