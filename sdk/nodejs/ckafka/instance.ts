// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
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
    public static readonly __pulumiType = 'tencentcloud:Ckafka/instance:Instance';

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
     * Instance bandwidth in MBps.
     */
    public readonly bandWidth!: pulumi.Output<number>;
    /**
     * Instance configuration.
     */
    public readonly config!: pulumi.Output<outputs.Ckafka.InstanceConfig | undefined>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through
     * the control. If it is not within the interval, the plan will cause a change when first created.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * Type of disk.
     */
    public readonly diskType!: pulumi.Output<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    public readonly dynamicRetentionConfig!: pulumi.Output<outputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    public readonly kafkaVersion!: pulumi.Output<string>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the
     * default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    public readonly msgRetentionTime!: pulumi.Output<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.
     */
    public readonly multiZoneFlag!: pulumi.Output<boolean | undefined>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed
     * through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    public readonly partition!: pulumi.Output<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Timestamp.
     */
    public readonly publicNetwork!: pulumi.Output<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    public readonly rebalanceTime!: pulumi.Output<number | undefined>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear
     * no automatic renewal (user setting).
     */
    public readonly renewFlag!: pulumi.Output<number>;
    /**
     * Subnet id.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Partition size, the professional version does not need tag.
     */
    public readonly tags!: pulumi.Output<outputs.Ckafka.InstanceTag[] | undefined>;
    /**
     * Vpc id.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Available zone id.
     */
    public readonly zoneId!: pulumi.Output<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.
     */
    public readonly zoneIds!: pulumi.Output<number[] | undefined>;

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
            resourceInputs["bandWidth"] = state ? state.bandWidth : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["diskType"] = state ? state.diskType : undefined;
            resourceInputs["dynamicRetentionConfig"] = state ? state.dynamicRetentionConfig : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["kafkaVersion"] = state ? state.kafkaVersion : undefined;
            resourceInputs["msgRetentionTime"] = state ? state.msgRetentionTime : undefined;
            resourceInputs["multiZoneFlag"] = state ? state.multiZoneFlag : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["publicNetwork"] = state ? state.publicNetwork : undefined;
            resourceInputs["rebalanceTime"] = state ? state.rebalanceTime : undefined;
            resourceInputs["renewFlag"] = state ? state.renewFlag : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["bandWidth"] = args ? args.bandWidth : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["diskType"] = args ? args.diskType : undefined;
            resourceInputs["dynamicRetentionConfig"] = args ? args.dynamicRetentionConfig : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["kafkaVersion"] = args ? args.kafkaVersion : undefined;
            resourceInputs["msgRetentionTime"] = args ? args.msgRetentionTime : undefined;
            resourceInputs["multiZoneFlag"] = args ? args.multiZoneFlag : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["publicNetwork"] = args ? args.publicNetwork : undefined;
            resourceInputs["rebalanceTime"] = args ? args.rebalanceTime : undefined;
            resourceInputs["renewFlag"] = args ? args.renewFlag : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
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
     * Instance bandwidth in MBps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * Instance configuration.
     */
    config?: pulumi.Input<inputs.Ckafka.InstanceConfig>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through
     * the control. If it is not within the interval, the plan will cause a change when first created.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Type of disk.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    dynamicRetentionConfig?: pulumi.Input<inputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the
     * default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    msgRetentionTime?: pulumi.Input<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.
     */
    multiZoneFlag?: pulumi.Input<boolean>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed
     * through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    partition?: pulumi.Input<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    period?: pulumi.Input<number>;
    /**
     * Timestamp.
     */
    publicNetwork?: pulumi.Input<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    rebalanceTime?: pulumi.Input<number>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear
     * no automatic renewal (user setting).
     */
    renewFlag?: pulumi.Input<number>;
    /**
     * Subnet id.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Partition size, the professional version does not need tag.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Ckafka.InstanceTag>[]>;
    /**
     * Vpc id.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id.
     */
    zoneId?: pulumi.Input<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.
     */
    zoneIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Instance bandwidth in MBps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * Instance configuration.
     */
    config?: pulumi.Input<inputs.Ckafka.InstanceConfig>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through
     * the control. If it is not within the interval, the plan will cause a change when first created.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Type of disk.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    dynamicRetentionConfig?: pulumi.Input<inputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the
     * default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    msgRetentionTime?: pulumi.Input<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.
     */
    multiZoneFlag?: pulumi.Input<boolean>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed
     * through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    partition?: pulumi.Input<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    period?: pulumi.Input<number>;
    /**
     * Timestamp.
     */
    publicNetwork?: pulumi.Input<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    rebalanceTime?: pulumi.Input<number>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear
     * no automatic renewal (user setting).
     */
    renewFlag?: pulumi.Input<number>;
    /**
     * Subnet id.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Partition size, the professional version does not need tag.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Ckafka.InstanceTag>[]>;
    /**
     * Vpc id.
     */
    vpcId: pulumi.Input<string>;
    /**
     * Available zone id.
     */
    zoneId: pulumi.Input<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.
     */
    zoneIds?: pulumi.Input<pulumi.Input<number>[]>;
}
