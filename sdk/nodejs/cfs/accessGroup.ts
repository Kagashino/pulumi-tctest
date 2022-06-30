// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AccessGroup extends pulumi.CustomResource {
    /**
     * Get an existing AccessGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessGroupState, opts?: pulumi.CustomResourceOptions): AccessGroup {
        return new AccessGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cfs/accessGroup:AccessGroup';

    /**
     * Returns true if the given object is an instance of AccessGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessGroup.__pulumiType;
    }

    /**
     * Create time of the access group.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the access group, and max length is 255.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the access group, and max length is 64.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AccessGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccessGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessGroupArgs | AccessGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessGroupState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AccessGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessGroup resources.
 */
export interface AccessGroupState {
    /**
     * Create time of the access group.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the access group, and max length is 255.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the access group, and max length is 64.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessGroup resource.
 */
export interface AccessGroupArgs {
    /**
     * Description of the access group, and max length is 255.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the access group, and max length is 64.
     */
    name?: pulumi.Input<string>;
}
