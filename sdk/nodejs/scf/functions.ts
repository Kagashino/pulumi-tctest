// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function functions(args?: FunctionsArgs, opts?: pulumi.InvokeOptions): Promise<FunctionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Scf/functions:Functions", {
        "description": args.description,
        "name": args.name,
        "namespace": args.namespace,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking Functions.
 */
export interface FunctionsArgs {
    description?: string;
    name?: string;
    namespace?: string;
    resultOutputFile?: string;
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by Functions.
 */
export interface FunctionsResult {
    readonly description?: string;
    readonly functions: outputs.Scf.FunctionsFunction[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly namespace?: string;
    readonly resultOutputFile?: string;
    readonly tags?: {[key: string]: any};
}

export function functionsOutput(args?: FunctionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FunctionsResult> {
    return pulumi.output(args).apply(a => functions(a, opts))
}

/**
 * A collection of arguments for invoking Functions.
 */
export interface FunctionsOutputArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}
