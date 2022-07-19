// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RoleSSO extends pulumi.CustomResource {
    /**
     * Get an existing RoleSSO resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleSSOState, opts?: pulumi.CustomResourceOptions): RoleSSO {
        return new RoleSSO(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Cam/roleSSO:RoleSSO';

    /**
     * Returns true if the given object is an instance of RoleSSO.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleSSO {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleSSO.__pulumiType;
    }

    /**
     * Client ids.
     */
    public readonly clientIds!: pulumi.Output<string[]>;
    /**
     * The description of resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Sign the public key.
     */
    public readonly identityKey!: pulumi.Output<string>;
    /**
     * Identity provider URL.
     */
    public readonly identityUrl!: pulumi.Output<string>;
    /**
     * The name of resource.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a RoleSSO resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleSSOArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleSSOArgs | RoleSSOState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleSSOState | undefined;
            resourceInputs["clientIds"] = state ? state.clientIds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["identityKey"] = state ? state.identityKey : undefined;
            resourceInputs["identityUrl"] = state ? state.identityUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as RoleSSOArgs | undefined;
            if ((!args || args.clientIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientIds'");
            }
            if ((!args || args.identityKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityKey'");
            }
            if ((!args || args.identityUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityUrl'");
            }
            resourceInputs["clientIds"] = args ? args.clientIds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identityKey"] = args ? args.identityKey : undefined;
            resourceInputs["identityUrl"] = args ? args.identityUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RoleSSO.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoleSSO resources.
 */
export interface RoleSSOState {
    /**
     * Client ids.
     */
    clientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Sign the public key.
     */
    identityKey?: pulumi.Input<string>;
    /**
     * Identity provider URL.
     */
    identityUrl?: pulumi.Input<string>;
    /**
     * The name of resource.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoleSSO resource.
 */
export interface RoleSSOArgs {
    /**
     * Client ids.
     */
    clientIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Sign the public key.
     */
    identityKey: pulumi.Input<string>;
    /**
     * Identity provider URL.
     */
    identityUrl: pulumi.Input<string>;
    /**
     * The name of resource.
     */
    name?: pulumi.Input<string>;
}
