// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cls/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * Whether to enable automatic split. Default value: true.
     */
    public readonly autoSplit!: pulumi.Output<boolean | undefined>;
    /**
     * Logset ID.
     */
    public readonly logsetId!: pulumi.Output<string>;
    /**
     * Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
     */
    public readonly maxSplitPartitions!: pulumi.Output<number | undefined>;
    /**
     * Number of log topic partitions. Default value: 1. Maximum value: 10.
     */
    public readonly partitionCount!: pulumi.Output<number | undefined>;
    /**
     * Lifecycle in days. Value range: 1~366. Default value: 30.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
     * passed in, please contact the customer service to add the log topic to the allowlist first..
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Log topic name.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            resourceInputs["autoSplit"] = state ? state.autoSplit : undefined;
            resourceInputs["logsetId"] = state ? state.logsetId : undefined;
            resourceInputs["maxSplitPartitions"] = state ? state.maxSplitPartitions : undefined;
            resourceInputs["partitionCount"] = state ? state.partitionCount : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if ((!args || args.logsetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logsetId'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["autoSplit"] = args ? args.autoSplit : undefined;
            resourceInputs["logsetId"] = args ? args.logsetId : undefined;
            resourceInputs["maxSplitPartitions"] = args ? args.maxSplitPartitions : undefined;
            resourceInputs["partitionCount"] = args ? args.partitionCount : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * Whether to enable automatic split. Default value: true.
     */
    autoSplit?: pulumi.Input<boolean>;
    /**
     * Logset ID.
     */
    logsetId?: pulumi.Input<string>;
    /**
     * Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
     */
    maxSplitPartitions?: pulumi.Input<number>;
    /**
     * Number of log topic partitions. Default value: 1. Maximum value: 10.
     */
    partitionCount?: pulumi.Input<number>;
    /**
     * Lifecycle in days. Value range: 1~366. Default value: 30.
     */
    period?: pulumi.Input<number>;
    /**
     * Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
     * passed in, please contact the customer service to add the log topic to the allowlist first..
     */
    storageType?: pulumi.Input<string>;
    /**
     * Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Log topic name.
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * Whether to enable automatic split. Default value: true.
     */
    autoSplit?: pulumi.Input<boolean>;
    /**
     * Logset ID.
     */
    logsetId: pulumi.Input<string>;
    /**
     * Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
     */
    maxSplitPartitions?: pulumi.Input<number>;
    /**
     * Number of log topic partitions. Default value: 1. Maximum value: 10.
     */
    partitionCount?: pulumi.Input<number>;
    /**
     * Lifecycle in days. Value range: 1~366. Default value: 30.
     */
    period?: pulumi.Input<number>;
    /**
     * Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
     * passed in, please contact the customer service to add the log topic to the allowlist first..
     */
    storageType?: pulumi.Input<string>;
    /**
     * Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Log topic name.
     */
    topicName: pulumi.Input<string>;
}
