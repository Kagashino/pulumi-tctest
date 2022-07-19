// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func APIGatewayCustomerDomains(ctx *pulumi.Context, args *APIGatewayCustomerDomainsArgs, opts ...pulumi.InvokeOption) (*APIGatewayCustomerDomainsResult, error) {
	var rv APIGatewayCustomerDomainsResult
	err := ctx.Invoke("tctest:Cloud/aPIGatewayCustomerDomains:APIGatewayCustomerDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking APIGatewayCustomerDomains.
type APIGatewayCustomerDomainsArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	ServiceId        string  `pulumi:"serviceId"`
}

// A collection of values returned by APIGatewayCustomerDomains.
type APIGatewayCustomerDomainsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                          `pulumi:"id"`
	Lists            []APIGatewayCustomerDomainsList `pulumi:"lists"`
	ResultOutputFile *string                         `pulumi:"resultOutputFile"`
	ServiceId        string                          `pulumi:"serviceId"`
}

func APIGatewayCustomerDomainsOutput(ctx *pulumi.Context, args APIGatewayCustomerDomainsOutputArgs, opts ...pulumi.InvokeOption) APIGatewayCustomerDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (APIGatewayCustomerDomainsResult, error) {
			args := v.(APIGatewayCustomerDomainsArgs)
			r, err := APIGatewayCustomerDomains(ctx, &args, opts...)
			var s APIGatewayCustomerDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(APIGatewayCustomerDomainsResultOutput)
}

// A collection of arguments for invoking APIGatewayCustomerDomains.
type APIGatewayCustomerDomainsOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	ServiceId        pulumi.StringInput    `pulumi:"serviceId"`
}

func (APIGatewayCustomerDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*APIGatewayCustomerDomainsArgs)(nil)).Elem()
}

// A collection of values returned by APIGatewayCustomerDomains.
type APIGatewayCustomerDomainsResultOutput struct{ *pulumi.OutputState }

func (APIGatewayCustomerDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIGatewayCustomerDomainsResult)(nil)).Elem()
}

func (o APIGatewayCustomerDomainsResultOutput) ToAPIGatewayCustomerDomainsResultOutput() APIGatewayCustomerDomainsResultOutput {
	return o
}

func (o APIGatewayCustomerDomainsResultOutput) ToAPIGatewayCustomerDomainsResultOutputWithContext(ctx context.Context) APIGatewayCustomerDomainsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o APIGatewayCustomerDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v APIGatewayCustomerDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o APIGatewayCustomerDomainsResultOutput) Lists() APIGatewayCustomerDomainsListArrayOutput {
	return o.ApplyT(func(v APIGatewayCustomerDomainsResult) []APIGatewayCustomerDomainsList { return v.Lists }).(APIGatewayCustomerDomainsListArrayOutput)
}

func (o APIGatewayCustomerDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v APIGatewayCustomerDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o APIGatewayCustomerDomainsResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v APIGatewayCustomerDomainsResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(APIGatewayCustomerDomainsResultOutput{})
}
