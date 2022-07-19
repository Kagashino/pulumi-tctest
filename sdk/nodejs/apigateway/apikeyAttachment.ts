// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class APIKeyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing APIKeyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: APIKeyAttachmentState, opts?: pulumi.CustomResourceOptions): APIKeyAttachment {
        return new APIKeyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:APIGateway/aPIKeyAttachment:APIKeyAttachment';

    /**
     * Returns true if the given object is an instance of APIKeyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is APIKeyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === APIKeyAttachment.__pulumiType;
    }

    /**
     * ID of API key.
     */
    public readonly apiKeyId!: pulumi.Output<string>;
    /**
     * ID of the usage plan.
     */
    public readonly usagePlanId!: pulumi.Output<string>;

    /**
     * Create a APIKeyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: APIKeyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: APIKeyAttachmentArgs | APIKeyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as APIKeyAttachmentState | undefined;
            resourceInputs["apiKeyId"] = state ? state.apiKeyId : undefined;
            resourceInputs["usagePlanId"] = state ? state.usagePlanId : undefined;
        } else {
            const args = argsOrState as APIKeyAttachmentArgs | undefined;
            if ((!args || args.apiKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKeyId'");
            }
            if ((!args || args.usagePlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'usagePlanId'");
            }
            resourceInputs["apiKeyId"] = args ? args.apiKeyId : undefined;
            resourceInputs["usagePlanId"] = args ? args.usagePlanId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(APIKeyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering APIKeyAttachment resources.
 */
export interface APIKeyAttachmentState {
    /**
     * ID of API key.
     */
    apiKeyId?: pulumi.Input<string>;
    /**
     * ID of the usage plan.
     */
    usagePlanId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a APIKeyAttachment resource.
 */
export interface APIKeyAttachmentArgs {
    /**
     * ID of API key.
     */
    apiKeyId: pulumi.Input<string>;
    /**
     * ID of the usage plan.
     */
    usagePlanId: pulumi.Input<string>;
}
