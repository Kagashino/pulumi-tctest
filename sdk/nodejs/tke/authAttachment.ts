// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AuthAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AuthAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthAttachmentState, opts?: pulumi.CustomResourceOptions): AuthAttachment {
        return new AuthAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tke/authAttachment:AuthAttachment';

    /**
     * Returns true if the given object is an instance of AuthAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthAttachment.__pulumiType;
    }

    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
     * '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    public readonly autoCreateDiscoveryAnonymousAuth!: pulumi.Output<boolean | undefined>;
    /**
     * ID of clusters.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Specify service-account-issuer.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    public readonly jwksUri!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthAttachmentArgs | AuthAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthAttachmentState | undefined;
            resourceInputs["autoCreateDiscoveryAnonymousAuth"] = state ? state.autoCreateDiscoveryAnonymousAuth : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["jwksUri"] = state ? state.jwksUri : undefined;
        } else {
            const args = argsOrState as AuthAttachmentArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.issuer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuer'");
            }
            resourceInputs["autoCreateDiscoveryAnonymousAuth"] = args ? args.autoCreateDiscoveryAnonymousAuth : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["jwksUri"] = args ? args.jwksUri : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthAttachment resources.
 */
export interface AuthAttachmentState {
    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
     * '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    autoCreateDiscoveryAnonymousAuth?: pulumi.Input<boolean>;
    /**
     * ID of clusters.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Specify service-account-issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    jwksUri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthAttachment resource.
 */
export interface AuthAttachmentArgs {
    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
     * '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    autoCreateDiscoveryAnonymousAuth?: pulumi.Input<boolean>;
    /**
     * ID of clusters.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Specify service-account-issuer.
     */
    issuer: pulumi.Input<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    jwksUri?: pulumi.Input<string>;
}
