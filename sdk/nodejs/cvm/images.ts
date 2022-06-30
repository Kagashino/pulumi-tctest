// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function images(args?: ImagesArgs, opts?: pulumi.InvokeOptions): Promise<ImagesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cvm/images:Images", {
        "imageId": args.imageId,
        "imageNameRegex": args.imageNameRegex,
        "imageTypes": args.imageTypes,
        "osName": args.osName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Images.
 */
export interface ImagesArgs {
    imageId?: string;
    imageNameRegex?: string;
    imageTypes?: string[];
    osName?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by Images.
 */
export interface ImagesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId?: string;
    readonly imageNameRegex?: string;
    readonly imageTypes?: string[];
    readonly images: outputs.Cvm.ImagesImage[];
    readonly osName?: string;
    readonly resultOutputFile?: string;
}

export function imagesOutput(args?: ImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ImagesResult> {
    return pulumi.output(args).apply(a => images(a, opts))
}

/**
 * A collection of arguments for invoking Images.
 */
export interface ImagesOutputArgs {
    imageId?: pulumi.Input<string>;
    imageNameRegex?: pulumi.Input<string>;
    imageTypes?: pulumi.Input<pulumi.Input<string>[]>;
    osName?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
