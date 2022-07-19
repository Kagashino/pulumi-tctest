// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AddonAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AddonAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddonAttachmentState, opts?: pulumi.CustomResourceOptions): AddonAttachment {
        return new AddonAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Tke/addonAttachment:AddonAttachment';

    /**
     * Returns true if the given object is an instance of AddonAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddonAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddonAttachment.__pulumiType;
    }

    /**
     * ID of cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Name of addon.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
     */
    public readonly requestBody!: pulumi.Output<string | undefined>;
    /**
     * Addon response body.
     */
    public /*out*/ readonly responseBody!: pulumi.Output<string>;
    /**
     * Addon current status.
     */
    public /*out*/ readonly status!: pulumi.Output<{[key: string]: any}>;
    /**
     * Values the addon passthroughs. Conflict with `request_body`.
     */
    public readonly values!: pulumi.Output<string[]>;
    /**
     * Addon version, default latest version. Conflict with `request_body`.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a AddonAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddonAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddonAttachmentArgs | AddonAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddonAttachmentState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["requestBody"] = state ? state.requestBody : undefined;
            resourceInputs["responseBody"] = state ? state.responseBody : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["values"] = state ? state.values : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AddonAttachmentArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["requestBody"] = args ? args.requestBody : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["responseBody"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AddonAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddonAttachment resources.
 */
export interface AddonAttachmentState {
    /**
     * ID of cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Name of addon.
     */
    name?: pulumi.Input<string>;
    /**
     * Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
     */
    requestBody?: pulumi.Input<string>;
    /**
     * Addon response body.
     */
    responseBody?: pulumi.Input<string>;
    /**
     * Addon current status.
     */
    status?: pulumi.Input<{[key: string]: any}>;
    /**
     * Values the addon passthroughs. Conflict with `request_body`.
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Addon version, default latest version. Conflict with `request_body`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AddonAttachment resource.
 */
export interface AddonAttachmentArgs {
    /**
     * ID of cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of addon.
     */
    name?: pulumi.Input<string>;
    /**
     * Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
     */
    requestBody?: pulumi.Input<string>;
    /**
     * Values the addon passthroughs. Conflict with `request_body`.
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Addon version, default latest version. Conflict with `request_body`.
     */
    version?: pulumi.Input<string>;
}
