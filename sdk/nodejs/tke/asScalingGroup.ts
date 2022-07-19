// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class AsScalingGroup extends pulumi.CustomResource {
    /**
     * Get an existing AsScalingGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AsScalingGroupState, opts?: pulumi.CustomResourceOptions): AsScalingGroup {
        return new AsScalingGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Tke/asScalingGroup:AsScalingGroup';

    /**
     * Returns true if the given object is an instance of AsScalingGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AsScalingGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AsScalingGroup.__pulumiType;
    }

    /**
     * Auto scaling config parameters.
     */
    public readonly autoScalingConfig!: pulumi.Output<outputs.Tke.AsScalingGroupAutoScalingConfig>;
    /**
     * Auto scaling group parameters.
     */
    public readonly autoScalingGroup!: pulumi.Output<outputs.Tke.AsScalingGroupAutoScalingGroup>;
    /**
     * ID of the cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Custom parameter information related to the node.
     */
    public readonly extraArgs!: pulumi.Output<string[] | undefined>;
    /**
     * Labels of kubernetes AS Group created nodes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    public readonly unschedulable!: pulumi.Output<number | undefined>;

    /**
     * Create a AsScalingGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AsScalingGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AsScalingGroupArgs | AsScalingGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AsScalingGroupState | undefined;
            resourceInputs["autoScalingConfig"] = state ? state.autoScalingConfig : undefined;
            resourceInputs["autoScalingGroup"] = state ? state.autoScalingGroup : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["extraArgs"] = state ? state.extraArgs : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
        } else {
            const args = argsOrState as AsScalingGroupArgs | undefined;
            if ((!args || args.autoScalingConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingConfig'");
            }
            if ((!args || args.autoScalingGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingGroup'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["autoScalingConfig"] = args ? args.autoScalingConfig : undefined;
            resourceInputs["autoScalingGroup"] = args ? args.autoScalingGroup : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["extraArgs"] = args ? args.extraArgs : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AsScalingGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AsScalingGroup resources.
 */
export interface AsScalingGroupState {
    /**
     * Auto scaling config parameters.
     */
    autoScalingConfig?: pulumi.Input<inputs.Tke.AsScalingGroupAutoScalingConfig>;
    /**
     * Auto scaling group parameters.
     */
    autoScalingGroup?: pulumi.Input<inputs.Tke.AsScalingGroupAutoScalingGroup>;
    /**
     * ID of the cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Custom parameter information related to the node.
     */
    extraArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Labels of kubernetes AS Group created nodes.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AsScalingGroup resource.
 */
export interface AsScalingGroupArgs {
    /**
     * Auto scaling config parameters.
     */
    autoScalingConfig: pulumi.Input<inputs.Tke.AsScalingGroupAutoScalingConfig>;
    /**
     * Auto scaling group parameters.
     */
    autoScalingGroup: pulumi.Input<inputs.Tke.AsScalingGroupAutoScalingGroup>;
    /**
     * ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Custom parameter information related to the node.
     */
    extraArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Labels of kubernetes AS Group created nodes.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
}
