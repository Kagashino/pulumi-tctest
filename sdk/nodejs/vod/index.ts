// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./adaptiveDynamicStreamingTemplate";
export * from "./adaptiveDynamicStreamingTemplates";
export * from "./imageSpriteTemplate";
export * from "./imageSpriteTemplates";
export * from "./procedureTemplate";
export * from "./procedureTemplates";
export * from "./snapshotByTimeOffsetTemplate";
export * from "./snapshotByTimeOffsetTemplates";
export * from "./subApplication";
export * from "./superPlayerConfig";
export * from "./superPlayerConfigs";

// Import resources to register:
import { AdaptiveDynamicStreamingTemplate } from "./adaptiveDynamicStreamingTemplate";
import { ImageSpriteTemplate } from "./imageSpriteTemplate";
import { ProcedureTemplate } from "./procedureTemplate";
import { SnapshotByTimeOffsetTemplate } from "./snapshotByTimeOffsetTemplate";
import { SubApplication } from "./subApplication";
import { SuperPlayerConfig } from "./superPlayerConfig";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate":
                return new AdaptiveDynamicStreamingTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Vod/imageSpriteTemplate:ImageSpriteTemplate":
                return new ImageSpriteTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Vod/procedureTemplate:ProcedureTemplate":
                return new ProcedureTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Vod/snapshotByTimeOffsetTemplate:SnapshotByTimeOffsetTemplate":
                return new SnapshotByTimeOffsetTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Vod/subApplication:SubApplication":
                return new SubApplication(name, <any>undefined, { urn })
            case "tencentcloud:Vod/superPlayerConfig:SuperPlayerConfig":
                return new SuperPlayerConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/adaptiveDynamicStreamingTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/imageSpriteTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/procedureTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/snapshotByTimeOffsetTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/subApplication", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vod/superPlayerConfig", _module)