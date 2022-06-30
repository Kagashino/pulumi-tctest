// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Association extends pulumi.CustomResource {
    /**
     * Get an existing Association resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssociationState, opts?: pulumi.CustomResourceOptions): Association {
        return new Association(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Eip/association:Association';

    /**
     * Returns true if the given object is an instance of Association.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Association {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Association.__pulumiType;
    }

    /**
     * The ID of EIP.
     */
    public readonly eipId!: pulumi.Output<string>;
    /**
     * The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
     * `private_ip fields`.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
     */
    public readonly privateIp!: pulumi.Output<string>;

    /**
     * Create a Association resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssociationArgs | AssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssociationState | undefined;
            resourceInputs["eipId"] = state ? state.eipId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
        } else {
            const args = argsOrState as AssociationArgs | undefined;
            if ((!args || args.eipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eipId'");
            }
            resourceInputs["eipId"] = args ? args.eipId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["privateIp"] = args ? args.privateIp : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Association.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Association resources.
 */
export interface AssociationState {
    /**
     * The ID of EIP.
     */
    eipId?: pulumi.Input<string>;
    /**
     * The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
     * `private_ip fields`.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
     */
    privateIp?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Association resource.
 */
export interface AssociationArgs {
    /**
     * The ID of EIP.
     */
    eipId: pulumi.Input<string>;
    /**
     * The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
     * `private_ip fields`.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
     */
    privateIp?: pulumi.Input<string>;
}
