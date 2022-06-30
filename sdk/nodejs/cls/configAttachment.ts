// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ConfigAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ConfigAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigAttachmentState, opts?: pulumi.CustomResourceOptions): ConfigAttachment {
        return new ConfigAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cls/configAttachment:ConfigAttachment';

    /**
     * Returns true if the given object is an instance of ConfigAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigAttachment.__pulumiType;
    }

    /**
     * Collection configuration id.
     */
    public readonly configId!: pulumi.Output<string>;
    /**
     * Machine group id.
     */
    public readonly groupId!: pulumi.Output<string>;

    /**
     * Create a ConfigAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigAttachmentArgs | ConfigAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigAttachmentState | undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
        } else {
            const args = argsOrState as ConfigAttachmentArgs | undefined;
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigAttachment resources.
 */
export interface ConfigAttachmentState {
    /**
     * Collection configuration id.
     */
    configId?: pulumi.Input<string>;
    /**
     * Machine group id.
     */
    groupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigAttachment resource.
 */
export interface ConfigAttachmentArgs {
    /**
     * Collection configuration id.
     */
    configId: pulumi.Input<string>;
    /**
     * Machine group id.
     */
    groupId: pulumi.Input<string>;
}
