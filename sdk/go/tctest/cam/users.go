// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Users(ctx *pulumi.Context, args *UsersArgs, opts ...pulumi.InvokeOption) (*UsersResult, error) {
	var rv UsersResult
	err := ctx.Invoke("tctest:Cam/users:Users", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Users.
type UsersArgs struct {
	ConsoleLogin     *bool   `pulumi:"consoleLogin"`
	CountryCode      *string `pulumi:"countryCode"`
	Email            *string `pulumi:"email"`
	Name             *string `pulumi:"name"`
	PhoneNum         *string `pulumi:"phoneNum"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Uid              *int    `pulumi:"uid"`
	Uin              *int    `pulumi:"uin"`
}

// A collection of values returned by Users.
type UsersResult struct {
	ConsoleLogin *bool   `pulumi:"consoleLogin"`
	CountryCode  *string `pulumi:"countryCode"`
	Email        *string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id               string          `pulumi:"id"`
	Name             *string         `pulumi:"name"`
	PhoneNum         *string         `pulumi:"phoneNum"`
	Remark           *string         `pulumi:"remark"`
	ResultOutputFile *string         `pulumi:"resultOutputFile"`
	Uid              *int            `pulumi:"uid"`
	Uin              *int            `pulumi:"uin"`
	UserLists        []UsersUserList `pulumi:"userLists"`
}

func UsersOutput(ctx *pulumi.Context, args UsersOutputArgs, opts ...pulumi.InvokeOption) UsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UsersResult, error) {
			args := v.(UsersArgs)
			r, err := Users(ctx, &args, opts...)
			var s UsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(UsersResultOutput)
}

// A collection of arguments for invoking Users.
type UsersOutputArgs struct {
	ConsoleLogin     pulumi.BoolPtrInput   `pulumi:"consoleLogin"`
	CountryCode      pulumi.StringPtrInput `pulumi:"countryCode"`
	Email            pulumi.StringPtrInput `pulumi:"email"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PhoneNum         pulumi.StringPtrInput `pulumi:"phoneNum"`
	Remark           pulumi.StringPtrInput `pulumi:"remark"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Uid              pulumi.IntPtrInput    `pulumi:"uid"`
	Uin              pulumi.IntPtrInput    `pulumi:"uin"`
}

func (UsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UsersArgs)(nil)).Elem()
}

// A collection of values returned by Users.
type UsersResultOutput struct{ *pulumi.OutputState }

func (UsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsersResult)(nil)).Elem()
}

func (o UsersResultOutput) ToUsersResultOutput() UsersResultOutput {
	return o
}

func (o UsersResultOutput) ToUsersResultOutputWithContext(ctx context.Context) UsersResultOutput {
	return o
}

func (o UsersResultOutput) ConsoleLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UsersResult) *bool { return v.ConsoleLogin }).(pulumi.BoolPtrOutput)
}

func (o UsersResultOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.CountryCode }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o UsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v UsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o UsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) PhoneNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.PhoneNum }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UsersResult) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

func (o UsersResultOutput) Uin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UsersResult) *int { return v.Uin }).(pulumi.IntPtrOutput)
}

func (o UsersResultOutput) UserLists() UsersUserListArrayOutput {
	return o.ApplyT(func(v UsersResult) []UsersUserList { return v.UserLists }).(UsersUserListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(UsersResultOutput{})
}
