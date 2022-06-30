// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

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
	case "tencentcloud:Gaap/certificate:Certificate":
		r = &Certificate{}
	case "tencentcloud:Gaap/domainErrorPageInfo:DomainErrorPageInfo":
		r = &DomainErrorPageInfo{}
	case "tencentcloud:Gaap/httpDomain:HttpDomain":
		r = &HttpDomain{}
	case "tencentcloud:Gaap/httpRule:HttpRule":
		r = &HttpRule{}
	case "tencentcloud:Gaap/layer4Listener:Layer4Listener":
		r = &Layer4Listener{}
	case "tencentcloud:Gaap/layer7Listener:Layer7Listener":
		r = &Layer7Listener{}
	case "tencentcloud:Gaap/proxy:Proxy":
		r = &Proxy{}
	case "tencentcloud:Gaap/realserver:Realserver":
		r = &Realserver{}
	case "tencentcloud:Gaap/securityPolicy:SecurityPolicy":
		r = &SecurityPolicy{}
	case "tencentcloud:Gaap/securityRule:SecurityRule":
		r = &SecurityRule{}
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
		"Gaap/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/domainErrorPageInfo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/httpDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/httpRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/layer4Listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/layer7Listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/proxy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/realserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/securityPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Gaap/securityRule",
		&module{version},
	)
}
