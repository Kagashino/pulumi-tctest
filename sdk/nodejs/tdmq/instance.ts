// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The Dedicated Cluster Id.
     */
    public readonly bindClusterId!: pulumi.Output<number | undefined>;
    /**
     * The name of tdmq cluster to be created.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Description of the tdmq cluster.
     */
    public readonly remark!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["bindClusterId"] = state ? state.bindClusterId : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            resourceInputs["bindClusterId"] = args ? args.bindClusterId : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The Dedicated Cluster Id.
     */
    bindClusterId?: pulumi.Input<number>;
    /**
     * The name of tdmq cluster to be created.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Description of the tdmq cluster.
     */
    remark?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The Dedicated Cluster Id.
     */
    bindClusterId?: pulumi.Input<number>;
    /**
     * The name of tdmq cluster to be created.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Description of the tdmq cluster.
     */
    remark?: pulumi.Input<string>;
}
