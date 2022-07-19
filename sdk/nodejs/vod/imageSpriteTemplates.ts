// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function imageSpriteTemplates(args?: ImageSpriteTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<ImageSpriteTemplatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Vod/imageSpriteTemplates:ImageSpriteTemplates", {
        "definition": args.definition,
        "resultOutputFile": args.resultOutputFile,
        "subAppId": args.subAppId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking ImageSpriteTemplates.
 */
export interface ImageSpriteTemplatesArgs {
    definition?: string;
    resultOutputFile?: string;
    subAppId?: number;
    type?: string;
}

/**
 * A collection of values returned by ImageSpriteTemplates.
 */
export interface ImageSpriteTemplatesResult {
    readonly definition?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly subAppId?: number;
    readonly templateLists: outputs.Vod.ImageSpriteTemplatesTemplateList[];
    readonly type?: string;
}

export function imageSpriteTemplatesOutput(args?: ImageSpriteTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ImageSpriteTemplatesResult> {
    return pulumi.output(args).apply(a => imageSpriteTemplates(a, opts))
}

/**
 * A collection of arguments for invoking ImageSpriteTemplates.
 */
export interface ImageSpriteTemplatesOutputArgs {
    definition?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subAppId?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}
