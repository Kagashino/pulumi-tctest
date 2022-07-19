// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class Config extends pulumi.CustomResource {
    /**
     * Get an existing Config resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigState, opts?: pulumi.CustomResourceOptions): Config {
        return new Config(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Cls/config:Config';

    /**
     * Returns true if the given object is an instance of Config.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Config {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Config.__pulumiType;
    }

    /**
     * Collection path blocklist.
     */
    public readonly excludePaths!: pulumi.Output<outputs.Cls.ConfigExcludePath[] | undefined>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    public readonly extractRule!: pulumi.Output<outputs.Cls.ConfigExtractRule>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
     * minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
     * Default value: minimalist_log.
     */
    public readonly logType!: pulumi.Output<string | undefined>;
    /**
     * Collection configuration name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    public readonly output!: pulumi.Output<string | undefined>;
    /**
     * Log collection path containing the filename.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    public readonly userDefineRule!: pulumi.Output<string | undefined>;

    /**
     * Create a Config resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigArgs | ConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigState | undefined;
            resourceInputs["excludePaths"] = state ? state.excludePaths : undefined;
            resourceInputs["extractRule"] = state ? state.extractRule : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["output"] = state ? state.output : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["userDefineRule"] = state ? state.userDefineRule : undefined;
        } else {
            const args = argsOrState as ConfigArgs | undefined;
            if ((!args || args.extractRule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extractRule'");
            }
            resourceInputs["excludePaths"] = args ? args.excludePaths : undefined;
            resourceInputs["extractRule"] = args ? args.extractRule : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["output"] = args ? args.output : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["userDefineRule"] = args ? args.userDefineRule : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Config.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Config resources.
 */
export interface ConfigState {
    /**
     * Collection path blocklist.
     */
    excludePaths?: pulumi.Input<pulumi.Input<inputs.Cls.ConfigExcludePath>[]>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    extractRule?: pulumi.Input<inputs.Cls.ConfigExtractRule>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
     * minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
     * Default value: minimalist_log.
     */
    logType?: pulumi.Input<string>;
    /**
     * Collection configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    output?: pulumi.Input<string>;
    /**
     * Log collection path containing the filename.
     */
    path?: pulumi.Input<string>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    userDefineRule?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Config resource.
 */
export interface ConfigArgs {
    /**
     * Collection path blocklist.
     */
    excludePaths?: pulumi.Input<pulumi.Input<inputs.Cls.ConfigExcludePath>[]>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    extractRule: pulumi.Input<inputs.Cls.ConfigExtractRule>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
     * minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
     * Default value: minimalist_log.
     */
    logType?: pulumi.Input<string>;
    /**
     * Collection configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    output?: pulumi.Input<string>;
    /**
     * Log collection path containing the filename.
     */
    path?: pulumi.Input<string>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    userDefineRule?: pulumi.Input<string>;
}
