// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class AlarmPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AlarmPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmPolicyState, opts?: pulumi.CustomResourceOptions): AlarmPolicy {
        return new AlarmPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/alarmPolicy:AlarmPolicy';

    /**
     * Returns true if the given object is an instance of AlarmPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlarmPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlarmPolicy.__pulumiType;
    }

    /**
     * A list of metric trigger condition.
     */
    public readonly conditions!: pulumi.Output<outputs.Monitor.AlarmPolicyConditions>;
    /**
     * ID of trigger condition template.
     */
    public readonly conditonTemplateId!: pulumi.Output<number | undefined>;
    /**
     * The alarm policy create time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether to enable, default is `1`.
     */
    public readonly enable!: pulumi.Output<number | undefined>;
    /**
     * A list of event trigger condition.
     */
    public readonly eventConditions!: pulumi.Output<outputs.Monitor.AlarmPolicyEventCondition[]>;
    /**
     * The type of monitor.
     */
    public readonly monitorType!: pulumi.Output<string>;
    /**
     * The type of alarm.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * List of notification rule IDs.
     */
    public readonly noticeIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of policy.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * Policy tag to bind object.
     */
    public readonly policyTags!: pulumi.Output<outputs.Monitor.AlarmPolicyPolicyTag[] | undefined>;
    /**
     * Project ID. For products with different projects, a value other than -1 must be passed in.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * The remark of policy group.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Triggered task list.
     */
    public readonly triggerTasks!: pulumi.Output<outputs.Monitor.AlarmPolicyTriggerTask[] | undefined>;
    /**
     * The alarm policy update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a AlarmPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmPolicyArgs | AlarmPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmPolicyState | undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["conditonTemplateId"] = state ? state.conditonTemplateId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["eventConditions"] = state ? state.eventConditions : undefined;
            resourceInputs["monitorType"] = state ? state.monitorType : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["noticeIds"] = state ? state.noticeIds : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["policyTags"] = state ? state.policyTags : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["triggerTasks"] = state ? state.triggerTasks : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AlarmPolicyArgs | undefined;
            if ((!args || args.monitorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorType'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["conditonTemplateId"] = args ? args.conditonTemplateId : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["eventConditions"] = args ? args.eventConditions : undefined;
            resourceInputs["monitorType"] = args ? args.monitorType : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["noticeIds"] = args ? args.noticeIds : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["policyTags"] = args ? args.policyTags : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["triggerTasks"] = args ? args.triggerTasks : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlarmPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlarmPolicy resources.
 */
export interface AlarmPolicyState {
    /**
     * A list of metric trigger condition.
     */
    conditions?: pulumi.Input<inputs.Monitor.AlarmPolicyConditions>;
    /**
     * ID of trigger condition template.
     */
    conditonTemplateId?: pulumi.Input<number>;
    /**
     * The alarm policy create time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether to enable, default is `1`.
     */
    enable?: pulumi.Input<number>;
    /**
     * A list of event trigger condition.
     */
    eventConditions?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyEventCondition>[]>;
    /**
     * The type of monitor.
     */
    monitorType?: pulumi.Input<string>;
    /**
     * The type of alarm.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of notification rule IDs.
     */
    noticeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of policy.
     */
    policyName?: pulumi.Input<string>;
    /**
     * Policy tag to bind object.
     */
    policyTags?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyPolicyTag>[]>;
    /**
     * Project ID. For products with different projects, a value other than -1 must be passed in.
     */
    projectId?: pulumi.Input<number>;
    /**
     * The remark of policy group.
     */
    remark?: pulumi.Input<string>;
    /**
     * Triggered task list.
     */
    triggerTasks?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyTriggerTask>[]>;
    /**
     * The alarm policy update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlarmPolicy resource.
 */
export interface AlarmPolicyArgs {
    /**
     * A list of metric trigger condition.
     */
    conditions?: pulumi.Input<inputs.Monitor.AlarmPolicyConditions>;
    /**
     * ID of trigger condition template.
     */
    conditonTemplateId?: pulumi.Input<number>;
    /**
     * Whether to enable, default is `1`.
     */
    enable?: pulumi.Input<number>;
    /**
     * A list of event trigger condition.
     */
    eventConditions?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyEventCondition>[]>;
    /**
     * The type of monitor.
     */
    monitorType: pulumi.Input<string>;
    /**
     * The type of alarm.
     */
    namespace: pulumi.Input<string>;
    /**
     * List of notification rule IDs.
     */
    noticeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of policy.
     */
    policyName: pulumi.Input<string>;
    /**
     * Policy tag to bind object.
     */
    policyTags?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyPolicyTag>[]>;
    /**
     * Project ID. For products with different projects, a value other than -1 must be passed in.
     */
    projectId?: pulumi.Input<number>;
    /**
     * The remark of policy group.
     */
    remark?: pulumi.Input<string>;
    /**
     * Triggered task list.
     */
    triggerTasks?: pulumi.Input<pulumi.Input<inputs.Monitor.AlarmPolicyTriggerTask>[]>;
}
