// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ACL extends pulumi.CustomResource {
    /**
     * Get an existing ACL resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ACLState, opts?: pulumi.CustomResourceOptions): ACL {
        return new ACL(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/aCL:ACL';

    /**
     * Returns true if the given object is an instance of ACL.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ACL {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ACL.__pulumiType;
    }

    /**
     * Creation time of ACL.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Egress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    public readonly egresses!: pulumi.Output<string[] | undefined>;
    /**
     * Ingress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    public readonly ingresses!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the network ACL.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the VPC instance.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ACL resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ACLArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ACLArgs | ACLState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ACLState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["egresses"] = state ? state.egresses : undefined;
            resourceInputs["ingresses"] = state ? state.ingresses : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ACLArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["egresses"] = args ? args.egresses : undefined;
            resourceInputs["ingresses"] = args ? args.ingresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ACL.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ACL resources.
 */
export interface ACLState {
    /**
     * Creation time of ACL.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Egress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    egresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ingress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    ingresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the network ACL.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the VPC instance.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ACL resource.
 */
export interface ACLArgs {
    /**
     * Egress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    egresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ingress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of
     * 'action' is `ACCEPT` and `DROP`. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is
     * `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol'
     * is `ICMP` or `ALL`, the 'port' must be `ALL`.
     */
    ingresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the network ACL.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the VPC instance.
     */
    vpcId: pulumi.Input<string>;
}