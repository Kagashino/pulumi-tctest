// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function connections(args?: ConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<ConnectionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Vpn/connections:Connections", {
        "customerGatewayId": args.customerGatewayId,
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking Connections.
 */
export interface ConnectionsArgs {
    customerGatewayId?: string;
    id?: string;
    name?: string;
    resultOutputFile?: string;
    tags?: {[key: string]: any};
    vpcId?: string;
    vpnGatewayId?: string;
}

/**
 * A collection of values returned by Connections.
 */
export interface ConnectionsResult {
    readonly connectionLists: outputs.Vpn.ConnectionsConnectionList[];
    readonly customerGatewayId?: string;
    readonly id?: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcId?: string;
    readonly vpnGatewayId?: string;
}

export function connectionsOutput(args?: ConnectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ConnectionsResult> {
    return pulumi.output(args).apply(a => connections(a, opts))
}

/**
 * A collection of arguments for invoking Connections.
 */
export interface ConnectionsOutputArgs {
    customerGatewayId?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vpcId?: pulumi.Input<string>;
    vpnGatewayId?: pulumi.Input<string>;
}
