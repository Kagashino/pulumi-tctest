// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Tables(ctx *pulumi.Context, args *TablesArgs, opts ...pulumi.InvokeOption) (*TablesResult, error) {
	var rv TablesResult
	err := ctx.Invoke("tctest:Tcaplus/tables:Tables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Tables.
type TablesArgs struct {
	ClusterId        string  `pulumi:"clusterId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	TableId          *string `pulumi:"tableId"`
	TableName        *string `pulumi:"tableName"`
	TablegroupId     *string `pulumi:"tablegroupId"`
}

// A collection of values returned by Tables.
type TablesResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string       `pulumi:"id"`
	Lists            []TablesList `pulumi:"lists"`
	ResultOutputFile *string      `pulumi:"resultOutputFile"`
	TableId          *string      `pulumi:"tableId"`
	TableName        *string      `pulumi:"tableName"`
	TablegroupId     *string      `pulumi:"tablegroupId"`
}

func TablesOutput(ctx *pulumi.Context, args TablesOutputArgs, opts ...pulumi.InvokeOption) TablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TablesResult, error) {
			args := v.(TablesArgs)
			r, err := Tables(ctx, &args, opts...)
			var s TablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TablesResultOutput)
}

// A collection of arguments for invoking Tables.
type TablesOutputArgs struct {
	ClusterId        pulumi.StringInput    `pulumi:"clusterId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	TableId          pulumi.StringPtrInput `pulumi:"tableId"`
	TableName        pulumi.StringPtrInput `pulumi:"tableName"`
	TablegroupId     pulumi.StringPtrInput `pulumi:"tablegroupId"`
}

func (TablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TablesArgs)(nil)).Elem()
}

// A collection of values returned by Tables.
type TablesResultOutput struct{ *pulumi.OutputState }

func (TablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TablesResult)(nil)).Elem()
}

func (o TablesResultOutput) ToTablesResultOutput() TablesResultOutput {
	return o
}

func (o TablesResultOutput) ToTablesResultOutputWithContext(ctx context.Context) TablesResultOutput {
	return o
}

func (o TablesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v TablesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o TablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TablesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o TablesResultOutput) Lists() TablesListArrayOutput {
	return o.ApplyT(func(v TablesResult) []TablesList { return v.Lists }).(TablesListArrayOutput)
}

func (o TablesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TablesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o TablesResultOutput) TableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TablesResult) *string { return v.TableId }).(pulumi.StringPtrOutput)
}

func (o TablesResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TablesResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o TablesResultOutput) TablegroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TablesResult) *string { return v.TablegroupId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TablesResultOutput{})
}
