// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class CCHttpPolicy extends pulumi.CustomResource {
    /**
     * Get an existing CCHttpPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CCHttpPolicyState, opts?: pulumi.CustomResourceOptions): CCHttpPolicy {
        return new CCHttpPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Dayu/cCHttpPolicy:CCHttpPolicy';

    /**
     * Returns true if the given object is an instance of CCHttpPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CCHttpPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CCHttpPolicy.__pulumiType;
    }

    /**
     * Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Create time of the CC self-define http policy.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
     */
    public readonly frequency!: pulumi.Output<number>;
    /**
     * Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
     * set one.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Name of the CC self-define http policy. Length should between 1 and 20.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Id of the CC self-define http policy.
     */
    public /*out*/ readonly policyId!: pulumi.Output<string>;
    /**
     * ID of the resource that the CC self-define http policy works for.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
     * `net`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
     */
    public readonly ruleLists!: pulumi.Output<outputs.Dayu.CCHttpPolicyRuleList[] | undefined>;
    /**
     * Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
     * one.
     */
    public readonly smode!: pulumi.Output<string | undefined>;
    /**
     * Indicate the CC self-define http policy takes effect or not.
     */
    public readonly switch!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CCHttpPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CCHttpPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CCHttpPolicyArgs | CCHttpPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CCHttpPolicyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["ruleLists"] = state ? state.ruleLists : undefined;
            resourceInputs["smode"] = state ? state.smode : undefined;
            resourceInputs["switch"] = state ? state.switch : undefined;
        } else {
            const args = argsOrState as CCHttpPolicyArgs | undefined;
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["ruleLists"] = args ? args.ruleLists : undefined;
            resourceInputs["smode"] = args ? args.smode : undefined;
            resourceInputs["switch"] = args ? args.switch : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["policyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CCHttpPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CCHttpPolicy resources.
 */
export interface CCHttpPolicyState {
    /**
     * Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
     */
    action?: pulumi.Input<string>;
    /**
     * Create time of the CC self-define http policy.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
     */
    frequency?: pulumi.Input<number>;
    /**
     * Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
     * set one.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the CC self-define http policy. Length should between 1 and 20.
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the CC self-define http policy.
     */
    policyId?: pulumi.Input<string>;
    /**
     * ID of the resource that the CC self-define http policy works for.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
     * `net`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
     */
    ruleLists?: pulumi.Input<pulumi.Input<inputs.Dayu.CCHttpPolicyRuleList>[]>;
    /**
     * Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
     * one.
     */
    smode?: pulumi.Input<string>;
    /**
     * Indicate the CC self-define http policy takes effect or not.
     */
    switch?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CCHttpPolicy resource.
 */
export interface CCHttpPolicyArgs {
    /**
     * Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
     */
    action?: pulumi.Input<string>;
    /**
     * Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
     */
    frequency?: pulumi.Input<number>;
    /**
     * Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
     * set one.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the CC self-define http policy. Length should between 1 and 20.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the resource that the CC self-define http policy works for.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
     * `net`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
     */
    ruleLists?: pulumi.Input<pulumi.Input<inputs.Dayu.CCHttpPolicyRuleList>[]>;
    /**
     * Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
     * one.
     */
    smode?: pulumi.Input<string>;
    /**
     * Indicate the CC self-define http policy takes effect or not.
     */
    switch?: pulumi.Input<boolean>;
}
