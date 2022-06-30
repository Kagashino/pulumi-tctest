// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class TcrToken extends pulumi.CustomResource {
    /**
     * Get an existing TcrToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TcrTokenState, opts?: pulumi.CustomResourceOptions): TcrToken {
        return new TcrToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/tcrToken:TcrToken';

    /**
     * Returns true if the given object is an instance of TcrToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TcrToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TcrToken.__pulumiType;
    }

    /**
     * Create time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the token. Valid length is [0~255].
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicate to enable this token or not.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the TCR instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The content of the token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
     */
    public /*out*/ readonly tokenId!: pulumi.Output<string>;
    /**
     * User name of the token.
     */
    public /*out*/ readonly userName!: pulumi.Output<string>;

    /**
     * Create a TcrToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TcrTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TcrTokenArgs | TcrTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TcrTokenState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["tokenId"] = state ? state.tokenId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as TcrTokenArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["tokenId"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TcrToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TcrToken resources.
 */
export interface TcrTokenState {
    /**
     * Create time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the token. Valid length is [0~255].
     */
    description?: pulumi.Input<string>;
    /**
     * Indicate to enable this token or not.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * ID of the TCR instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The content of the token.
     */
    token?: pulumi.Input<string>;
    /**
     * Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
     */
    tokenId?: pulumi.Input<string>;
    /**
     * User name of the token.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TcrToken resource.
 */
export interface TcrTokenArgs {
    /**
     * Description of the token. Valid length is [0~255].
     */
    description?: pulumi.Input<string>;
    /**
     * Indicate to enable this token or not.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * ID of the TCR instance.
     */
    instanceId: pulumi.Input<string>;
}