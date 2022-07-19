// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
    public static readonly __pulumiType = 'tctest:Image/instance:Instance';

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
     * Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
     * image.
     */
    public readonly dataDiskIds!: pulumi.Output<string[]>;
    /**
     * Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
     * will be made after shutdown.
     */
    public readonly forcePoweroff!: pulumi.Output<boolean | undefined>;
    /**
     * Image Description.
     */
    public readonly imageDescription!: pulumi.Output<string | undefined>;
    /**
     * Image name.
     */
    public readonly imageName!: pulumi.Output<string>;
    /**
     * Cloud server instance ID.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
     * passed in simultaneously with InstanceId.
     */
    public readonly snapshotIds!: pulumi.Output<string[] | undefined>;
    /**
     * Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
     * Syspre function.
     */
    public readonly sysprep!: pulumi.Output<boolean | undefined>;

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
            resourceInputs["dataDiskIds"] = state ? state.dataDiskIds : undefined;
            resourceInputs["forcePoweroff"] = state ? state.forcePoweroff : undefined;
            resourceInputs["imageDescription"] = state ? state.imageDescription : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["snapshotIds"] = state ? state.snapshotIds : undefined;
            resourceInputs["sysprep"] = state ? state.sysprep : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            resourceInputs["dataDiskIds"] = args ? args.dataDiskIds : undefined;
            resourceInputs["forcePoweroff"] = args ? args.forcePoweroff : undefined;
            resourceInputs["imageDescription"] = args ? args.imageDescription : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["snapshotIds"] = args ? args.snapshotIds : undefined;
            resourceInputs["sysprep"] = args ? args.sysprep : undefined;
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
     * Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
     * image.
     */
    dataDiskIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
     * will be made after shutdown.
     */
    forcePoweroff?: pulumi.Input<boolean>;
    /**
     * Image Description.
     */
    imageDescription?: pulumi.Input<string>;
    /**
     * Image name.
     */
    imageName?: pulumi.Input<string>;
    /**
     * Cloud server instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
     * passed in simultaneously with InstanceId.
     */
    snapshotIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
     * Syspre function.
     */
    sysprep?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
     * image.
     */
    dataDiskIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
     * will be made after shutdown.
     */
    forcePoweroff?: pulumi.Input<boolean>;
    /**
     * Image Description.
     */
    imageDescription?: pulumi.Input<string>;
    /**
     * Image name.
     */
    imageName: pulumi.Input<string>;
    /**
     * Cloud server instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
     * passed in simultaneously with InstanceId.
     */
    snapshotIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
     * Syspre function.
     */
    sysprep?: pulumi.Input<boolean>;
}
