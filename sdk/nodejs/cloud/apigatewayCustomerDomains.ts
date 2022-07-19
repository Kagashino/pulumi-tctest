// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function apigatewayCustomerDomains(args: ApigatewayCustomerDomainsArgs, opts?: pulumi.InvokeOptions): Promise<ApigatewayCustomerDomainsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cloud/aPIGatewayCustomerDomains:APIGatewayCustomerDomains", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking APIGatewayCustomerDomains.
 */
export interface ApigatewayCustomerDomainsArgs {
    resultOutputFile?: string;
    serviceId: string;
}

/**
 * A collection of values returned by APIGatewayCustomerDomains.
 */
export interface ApigatewayCustomerDomainsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Cloud.APIGatewayCustomerDomainsList[];
    readonly resultOutputFile?: string;
    readonly serviceId: string;
}

export function apigatewayCustomerDomainsOutput(args: ApigatewayCustomerDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ApigatewayCustomerDomainsResult> {
    return pulumi.output(args).apply(a => apigatewayCustomerDomains(a, opts))
}

/**
 * A collection of arguments for invoking APIGatewayCustomerDomains.
 */
export interface ApigatewayCustomerDomainsOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
}
