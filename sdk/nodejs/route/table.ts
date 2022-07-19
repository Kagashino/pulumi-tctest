// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function table(args: TableArgs, opts?: pulumi.InvokeOptions): Promise<TableResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Route/table:Table", {
        "name": args.name,
        "routeTableId": args.routeTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking Table.
 */
export interface TableArgs {
    name?: string;
    routeTableId: string;
}

/**
 * A collection of values returned by Table.
 */
export interface TableResult {
    readonly createTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly routeTableId: string;
    readonly routes: outputs.Route.TableRoute[];
    readonly subnetNum: number;
    readonly vpcId: string;
}

export function tableOutput(args: TableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TableResult> {
    return pulumi.output(args).apply(a => table(a, opts))
}

/**
 * A collection of arguments for invoking Table.
 */
export interface TableOutputArgs {
    name?: pulumi.Input<string>;
    routeTableId: pulumi.Input<string>;
}
