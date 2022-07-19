// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function listenerRules(args: ListenerRulesArgs, opts?: pulumi.InvokeOptions): Promise<ListenerRulesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Clb/listenerRules:ListenerRules", {
        "clbId": args.clbId,
        "domain": args.domain,
        "listenerId": args.listenerId,
        "resultOutputFile": args.resultOutputFile,
        "ruleId": args.ruleId,
        "scheduler": args.scheduler,
        "url": args.url,
    }, opts);
}

/**
 * A collection of arguments for invoking ListenerRules.
 */
export interface ListenerRulesArgs {
    clbId: string;
    domain?: string;
    listenerId: string;
    resultOutputFile?: string;
    ruleId?: string;
    scheduler?: string;
    url?: string;
}

/**
 * A collection of values returned by ListenerRules.
 */
export interface ListenerRulesResult {
    readonly clbId: string;
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listenerId: string;
    readonly resultOutputFile?: string;
    readonly ruleId?: string;
    readonly ruleLists: outputs.Clb.ListenerRulesRuleList[];
    readonly scheduler?: string;
    readonly url?: string;
}

export function listenerRulesOutput(args: ListenerRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListenerRulesResult> {
    return pulumi.output(args).apply(a => listenerRules(a, opts))
}

/**
 * A collection of arguments for invoking ListenerRules.
 */
export interface ListenerRulesOutputArgs {
    clbId: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    listenerId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    ruleId?: pulumi.Input<string>;
    scheduler?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}
