// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tencentcloud

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/Kagashino/tctest/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "tencentcloud"
	// modules:
	mainMod = "index" // the tencentcloud module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	tcProvider := tencentcloud.Provider().(*schema.Provider)
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(tcProvider)

	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "tencentcloud",
		DisplayName: "Tencentcloud",
		Publisher: "Kagashino",
		LogoURL: "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing tencentcloud cloud resources.",
		Keywords:   []string{"pulumi", "tencentcloud", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "github.com/kagashino/tctest",
		GitHubOrg: "",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TENCENTCLOUD_REGION"},
				},
			},
			"secret_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TENCENTCLOUD_REGION"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TENCENTCLOUD_REGION"},
				},
			},
			"security_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TENCENTCLOUD_SECURITY_TOKEN"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            ResourceInfo,
		DataSources:          DataSourceInfo,
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
