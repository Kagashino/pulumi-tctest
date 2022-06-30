// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:APIGateway/aPI:API":
		r = &API{}
	case "tencentcloud:APIGateway/aPIKey:APIKey":
		r = &APIKey{}
	case "tencentcloud:APIGateway/aPIKeyAttachment:APIKeyAttachment":
		r = &APIKeyAttachment{}
	case "tencentcloud:APIGateway/customDomain:CustomDomain":
		r = &CustomDomain{}
	case "tencentcloud:APIGateway/iPStrategy:IPStrategy":
		r = &IPStrategy{}
	case "tencentcloud:APIGateway/service:Service":
		r = &Service{}
	case "tencentcloud:APIGateway/serviceRelease:ServiceRelease":
		r = &ServiceRelease{}
	case "tencentcloud:APIGateway/strategyAttachment:StrategyAttachment":
		r = &StrategyAttachment{}
	case "tencentcloud:APIGateway/usagePlan:UsagePlan":
		r = &UsagePlan{}
	case "tencentcloud:APIGateway/usagePlanAttachment:UsagePlanAttachment":
		r = &UsagePlanAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := tencentcloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/aPI",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/aPIKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/aPIKeyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/customDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/iPStrategy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/serviceRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/strategyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/usagePlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"APIGateway/usagePlanAttachment",
		&module{version},
	)
}
