// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function apigatewayUsagePlans(args?: ApigatewayUsagePlansArgs, opts?: pulumi.InvokeOptions): Promise<ApigatewayUsagePlansResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cloud/aPIGatewayUsagePlans:APIGatewayUsagePlans", {
        "resultOutputFile": args.resultOutputFile,
        "usagePlanId": args.usagePlanId,
        "usagePlanName": args.usagePlanName,
    }, opts);
}

/**
 * A collection of arguments for invoking APIGatewayUsagePlans.
 */
export interface ApigatewayUsagePlansArgs {
    resultOutputFile?: string;
    usagePlanId?: string;
    usagePlanName?: string;
}

/**
 * A collection of values returned by APIGatewayUsagePlans.
 */
export interface ApigatewayUsagePlansResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Cloud.APIGatewayUsagePlansList[];
    readonly resultOutputFile?: string;
    readonly usagePlanId?: string;
    readonly usagePlanName?: string;
}

export function apigatewayUsagePlansOutput(args?: ApigatewayUsagePlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ApigatewayUsagePlansResult> {
    return pulumi.output(args).apply(a => apigatewayUsagePlans(a, opts))
}

/**
 * A collection of arguments for invoking APIGatewayUsagePlans.
 */
export interface ApigatewayUsagePlansOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    usagePlanId?: pulumi.Input<string>;
    usagePlanName?: pulumi.Input<string>;
}
