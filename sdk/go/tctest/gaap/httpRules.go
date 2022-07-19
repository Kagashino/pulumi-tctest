// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func HttpRules(ctx *pulumi.Context, args *HttpRulesArgs, opts ...pulumi.InvokeOption) (*HttpRulesResult, error) {
	var rv HttpRulesResult
	err := ctx.Invoke("tctest:Gaap/httpRules:HttpRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking HttpRules.
type HttpRulesArgs struct {
	Domain           *string `pulumi:"domain"`
	ForwardHost      *string `pulumi:"forwardHost"`
	ListenerId       string  `pulumi:"listenerId"`
	Path             *string `pulumi:"path"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by HttpRules.
type HttpRulesResult struct {
	Domain      *string `pulumi:"domain"`
	ForwardHost *string `pulumi:"forwardHost"`
	// The provider-assigned unique ID for this managed resource.
	Id               string          `pulumi:"id"`
	ListenerId       string          `pulumi:"listenerId"`
	Path             *string         `pulumi:"path"`
	ResultOutputFile *string         `pulumi:"resultOutputFile"`
	Rules            []HttpRulesRule `pulumi:"rules"`
}

func HttpRulesOutput(ctx *pulumi.Context, args HttpRulesOutputArgs, opts ...pulumi.InvokeOption) HttpRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (HttpRulesResult, error) {
			args := v.(HttpRulesArgs)
			r, err := HttpRules(ctx, &args, opts...)
			var s HttpRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(HttpRulesResultOutput)
}

// A collection of arguments for invoking HttpRules.
type HttpRulesOutputArgs struct {
	Domain           pulumi.StringPtrInput `pulumi:"domain"`
	ForwardHost      pulumi.StringPtrInput `pulumi:"forwardHost"`
	ListenerId       pulumi.StringInput    `pulumi:"listenerId"`
	Path             pulumi.StringPtrInput `pulumi:"path"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (HttpRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRulesArgs)(nil)).Elem()
}

// A collection of values returned by HttpRules.
type HttpRulesResultOutput struct{ *pulumi.OutputState }

func (HttpRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRulesResult)(nil)).Elem()
}

func (o HttpRulesResultOutput) ToHttpRulesResultOutput() HttpRulesResultOutput {
	return o
}

func (o HttpRulesResultOutput) ToHttpRulesResultOutputWithContext(ctx context.Context) HttpRulesResultOutput {
	return o
}

func (o HttpRulesResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRulesResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o HttpRulesResultOutput) ForwardHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRulesResult) *string { return v.ForwardHost }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o HttpRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o HttpRulesResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRulesResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

func (o HttpRulesResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRulesResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HttpRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o HttpRulesResultOutput) Rules() HttpRulesRuleArrayOutput {
	return o.ApplyT(func(v HttpRulesResult) []HttpRulesRule { return v.Rules }).(HttpRulesRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(HttpRulesResultOutput{})
}
