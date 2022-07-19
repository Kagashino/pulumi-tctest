// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:APIGateway/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A list of APIs.
     */
    public /*out*/ readonly apiLists!: pulumi.Output<outputs.APIGateway.ServiceApiList[]>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
     */
    public readonly exclusiveSetName!: pulumi.Output<string | undefined>;
    /**
     * Port number for http access over private network.
     */
    public /*out*/ readonly innerHttpPort!: pulumi.Output<number>;
    /**
     * Port number for https access over private network.
     */
    public /*out*/ readonly innerHttpsPort!: pulumi.Output<number>;
    /**
     * Private network access subdomain name.
     */
    public /*out*/ readonly internalSubDomain!: pulumi.Output<string>;
    /**
     * IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
     */
    public readonly ipVersion!: pulumi.Output<string | undefined>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
     * indicates access over private network, and `OUTER` indicates access over public network.
     */
    public readonly netTypes!: pulumi.Output<string[]>;
    /**
     * Public network access subdomain name.
     */
    public /*out*/ readonly outerSubDomain!: pulumi.Output<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    public readonly preLimit!: pulumi.Output<number>;
    /**
     * Service frontend request type. Valid values: `http`, `https`, `http&https`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    public readonly releaseLimit!: pulumi.Output<number>;
    /**
     * Custom service description.
     */
    public readonly serviceDesc!: pulumi.Output<string | undefined>;
    /**
     * Custom service name.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    public readonly testLimit!: pulumi.Output<number>;
    /**
     * A list of attach usage plans.
     */
    public /*out*/ readonly usagePlanLists!: pulumi.Output<outputs.APIGateway.ServiceUsagePlanList[]>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["apiLists"] = state ? state.apiLists : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["exclusiveSetName"] = state ? state.exclusiveSetName : undefined;
            resourceInputs["innerHttpPort"] = state ? state.innerHttpPort : undefined;
            resourceInputs["innerHttpsPort"] = state ? state.innerHttpsPort : undefined;
            resourceInputs["internalSubDomain"] = state ? state.internalSubDomain : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["modifyTime"] = state ? state.modifyTime : undefined;
            resourceInputs["netTypes"] = state ? state.netTypes : undefined;
            resourceInputs["outerSubDomain"] = state ? state.outerSubDomain : undefined;
            resourceInputs["preLimit"] = state ? state.preLimit : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["releaseLimit"] = state ? state.releaseLimit : undefined;
            resourceInputs["serviceDesc"] = state ? state.serviceDesc : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["testLimit"] = state ? state.testLimit : undefined;
            resourceInputs["usagePlanLists"] = state ? state.usagePlanLists : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.netTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'netTypes'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["exclusiveSetName"] = args ? args.exclusiveSetName : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["netTypes"] = args ? args.netTypes : undefined;
            resourceInputs["preLimit"] = args ? args.preLimit : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["releaseLimit"] = args ? args.releaseLimit : undefined;
            resourceInputs["serviceDesc"] = args ? args.serviceDesc : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["testLimit"] = args ? args.testLimit : undefined;
            resourceInputs["apiLists"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["innerHttpPort"] = undefined /*out*/;
            resourceInputs["innerHttpsPort"] = undefined /*out*/;
            resourceInputs["internalSubDomain"] = undefined /*out*/;
            resourceInputs["modifyTime"] = undefined /*out*/;
            resourceInputs["outerSubDomain"] = undefined /*out*/;
            resourceInputs["usagePlanLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * A list of APIs.
     */
    apiLists?: pulumi.Input<pulumi.Input<inputs.APIGateway.ServiceApiList>[]>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
     */
    exclusiveSetName?: pulumi.Input<string>;
    /**
     * Port number for http access over private network.
     */
    innerHttpPort?: pulumi.Input<number>;
    /**
     * Port number for https access over private network.
     */
    innerHttpsPort?: pulumi.Input<number>;
    /**
     * Private network access subdomain name.
     */
    internalSubDomain?: pulumi.Input<string>;
    /**
     * IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
     * indicates access over private network, and `OUTER` indicates access over public network.
     */
    netTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Public network access subdomain name.
     */
    outerSubDomain?: pulumi.Input<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    preLimit?: pulumi.Input<number>;
    /**
     * Service frontend request type. Valid values: `http`, `https`, `http&https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    releaseLimit?: pulumi.Input<number>;
    /**
     * Custom service description.
     */
    serviceDesc?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    testLimit?: pulumi.Input<number>;
    /**
     * A list of attach usage plans.
     */
    usagePlanLists?: pulumi.Input<pulumi.Input<inputs.APIGateway.ServiceUsagePlanList>[]>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.
     */
    exclusiveSetName?: pulumi.Input<string>;
    /**
     * IP version number. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Network type list, which is used to specify the supported network types. Valid values: `INNER`, `OUTER`. `INNER`
     * indicates access over private network, and `OUTER` indicates access over public network.
     */
    netTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    preLimit?: pulumi.Input<number>;
    /**
     * Service frontend request type. Valid values: `http`, `https`, `http&https`.
     */
    protocol: pulumi.Input<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    releaseLimit?: pulumi.Input<number>;
    /**
     * Custom service description.
     */
    serviceDesc?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    serviceName: pulumi.Input<string>;
    /**
     * API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
     */
    testLimit?: pulumi.Input<number>;
}
