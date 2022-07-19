// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigExtra struct {
	pulumi.CustomResourceState

	// Collection configuration flag.
	ConfigFlag pulumi.StringOutput `pulumi:"configFlag"`
	// Container file path info.
	ContainerFile ConfigExtraContainerFilePtrOutput `pulumi:"containerFile"`
	// Container stdout info.
	ContainerStdout ConfigExtraContainerStdoutPtrOutput `pulumi:"containerStdout"`
	// Collection path blocklist.
	ExcludePaths ConfigExtraExcludePathArrayOutput `pulumi:"excludePaths"`
	// Extraction rule. If ExtractRule is set, LogType must be set.
	ExtractRule ConfigExtraExtractRulePtrOutput `pulumi:"extractRule"`
	// Binding group id.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// Binding group ids.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// Node file config info.
	HostFile ConfigExtraHostFilePtrOutput `pulumi:"hostFile"`
	// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
	// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
	// Default value: minimalist_log.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// Logset Id.
	LogsetId pulumi.StringOutput `pulumi:"logsetId"`
	// Logset Name.
	LogsetName pulumi.StringOutput `pulumi:"logsetName"`
	// Collection configuration name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Log topic ID (TopicId) of collection configuration.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
	// Topic Name.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// Type. Valid values: container_stdout; container_file; host_file.
	Type pulumi.StringOutput `pulumi:"type"`
	// Custom collection rule, which is a serialized JSON string.
	UserDefineRule pulumi.StringPtrOutput `pulumi:"userDefineRule"`
}

// NewConfigExtra registers a new resource with the given unique name, arguments, and options.
func NewConfigExtra(ctx *pulumi.Context,
	name string, args *ConfigExtraArgs, opts ...pulumi.ResourceOption) (*ConfigExtra, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigFlag == nil {
		return nil, errors.New("invalid value for required argument 'ConfigFlag'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	if args.LogsetId == nil {
		return nil, errors.New("invalid value for required argument 'LogsetId'")
	}
	if args.LogsetName == nil {
		return nil, errors.New("invalid value for required argument 'LogsetName'")
	}
	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ConfigExtra
	err := ctx.RegisterResource("tctest:Cls/configExtra:ConfigExtra", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigExtra gets an existing ConfigExtra resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigExtra(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigExtraState, opts ...pulumi.ResourceOption) (*ConfigExtra, error) {
	var resource ConfigExtra
	err := ctx.ReadResource("tctest:Cls/configExtra:ConfigExtra", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigExtra resources.
type configExtraState struct {
	// Collection configuration flag.
	ConfigFlag *string `pulumi:"configFlag"`
	// Container file path info.
	ContainerFile *ConfigExtraContainerFile `pulumi:"containerFile"`
	// Container stdout info.
	ContainerStdout *ConfigExtraContainerStdout `pulumi:"containerStdout"`
	// Collection path blocklist.
	ExcludePaths []ConfigExtraExcludePath `pulumi:"excludePaths"`
	// Extraction rule. If ExtractRule is set, LogType must be set.
	ExtractRule *ConfigExtraExtractRule `pulumi:"extractRule"`
	// Binding group id.
	GroupId *string `pulumi:"groupId"`
	// Binding group ids.
	GroupIds []string `pulumi:"groupIds"`
	// Node file config info.
	HostFile *ConfigExtraHostFile `pulumi:"hostFile"`
	// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
	// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
	// Default value: minimalist_log.
	LogType *string `pulumi:"logType"`
	// Logset Id.
	LogsetId *string `pulumi:"logsetId"`
	// Logset Name.
	LogsetName *string `pulumi:"logsetName"`
	// Collection configuration name.
	Name *string `pulumi:"name"`
	// Log topic ID (TopicId) of collection configuration.
	TopicId *string `pulumi:"topicId"`
	// Topic Name.
	TopicName *string `pulumi:"topicName"`
	// Type. Valid values: container_stdout; container_file; host_file.
	Type *string `pulumi:"type"`
	// Custom collection rule, which is a serialized JSON string.
	UserDefineRule *string `pulumi:"userDefineRule"`
}

type ConfigExtraState struct {
	// Collection configuration flag.
	ConfigFlag pulumi.StringPtrInput
	// Container file path info.
	ContainerFile ConfigExtraContainerFilePtrInput
	// Container stdout info.
	ContainerStdout ConfigExtraContainerStdoutPtrInput
	// Collection path blocklist.
	ExcludePaths ConfigExtraExcludePathArrayInput
	// Extraction rule. If ExtractRule is set, LogType must be set.
	ExtractRule ConfigExtraExtractRulePtrInput
	// Binding group id.
	GroupId pulumi.StringPtrInput
	// Binding group ids.
	GroupIds pulumi.StringArrayInput
	// Node file config info.
	HostFile ConfigExtraHostFilePtrInput
	// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
	// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
	// Default value: minimalist_log.
	LogType pulumi.StringPtrInput
	// Logset Id.
	LogsetId pulumi.StringPtrInput
	// Logset Name.
	LogsetName pulumi.StringPtrInput
	// Collection configuration name.
	Name pulumi.StringPtrInput
	// Log topic ID (TopicId) of collection configuration.
	TopicId pulumi.StringPtrInput
	// Topic Name.
	TopicName pulumi.StringPtrInput
	// Type. Valid values: container_stdout; container_file; host_file.
	Type pulumi.StringPtrInput
	// Custom collection rule, which is a serialized JSON string.
	UserDefineRule pulumi.StringPtrInput
}

func (ConfigExtraState) ElementType() reflect.Type {
	return reflect.TypeOf((*configExtraState)(nil)).Elem()
}

type configExtraArgs struct {
	// Collection configuration flag.
	ConfigFlag string `pulumi:"configFlag"`
	// Container file path info.
	ContainerFile *ConfigExtraContainerFile `pulumi:"containerFile"`
	// Container stdout info.
	ContainerStdout *ConfigExtraContainerStdout `pulumi:"containerStdout"`
	// Collection path blocklist.
	ExcludePaths []ConfigExtraExcludePath `pulumi:"excludePaths"`
	// Extraction rule. If ExtractRule is set, LogType must be set.
	ExtractRule *ConfigExtraExtractRule `pulumi:"extractRule"`
	// Binding group id.
	GroupId *string `pulumi:"groupId"`
	// Binding group ids.
	GroupIds []string `pulumi:"groupIds"`
	// Node file config info.
	HostFile *ConfigExtraHostFile `pulumi:"hostFile"`
	// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
	// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
	// Default value: minimalist_log.
	LogType string `pulumi:"logType"`
	// Logset Id.
	LogsetId string `pulumi:"logsetId"`
	// Logset Name.
	LogsetName string `pulumi:"logsetName"`
	// Collection configuration name.
	Name *string `pulumi:"name"`
	// Log topic ID (TopicId) of collection configuration.
	TopicId string `pulumi:"topicId"`
	// Topic Name.
	TopicName string `pulumi:"topicName"`
	// Type. Valid values: container_stdout; container_file; host_file.
	Type string `pulumi:"type"`
	// Custom collection rule, which is a serialized JSON string.
	UserDefineRule *string `pulumi:"userDefineRule"`
}

// The set of arguments for constructing a ConfigExtra resource.
type ConfigExtraArgs struct {
	// Collection configuration flag.
	ConfigFlag pulumi.StringInput
	// Container file path info.
	ContainerFile ConfigExtraContainerFilePtrInput
	// Container stdout info.
	ContainerStdout ConfigExtraContainerStdoutPtrInput
	// Collection path blocklist.
	ExcludePaths ConfigExtraExcludePathArrayInput
	// Extraction rule. If ExtractRule is set, LogType must be set.
	ExtractRule ConfigExtraExtractRulePtrInput
	// Binding group id.
	GroupId pulumi.StringPtrInput
	// Binding group ids.
	GroupIds pulumi.StringArrayInput
	// Node file config info.
	HostFile ConfigExtraHostFilePtrInput
	// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
	// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
	// Default value: minimalist_log.
	LogType pulumi.StringInput
	// Logset Id.
	LogsetId pulumi.StringInput
	// Logset Name.
	LogsetName pulumi.StringInput
	// Collection configuration name.
	Name pulumi.StringPtrInput
	// Log topic ID (TopicId) of collection configuration.
	TopicId pulumi.StringInput
	// Topic Name.
	TopicName pulumi.StringInput
	// Type. Valid values: container_stdout; container_file; host_file.
	Type pulumi.StringInput
	// Custom collection rule, which is a serialized JSON string.
	UserDefineRule pulumi.StringPtrInput
}

func (ConfigExtraArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configExtraArgs)(nil)).Elem()
}

type ConfigExtraInput interface {
	pulumi.Input

	ToConfigExtraOutput() ConfigExtraOutput
	ToConfigExtraOutputWithContext(ctx context.Context) ConfigExtraOutput
}

func (*ConfigExtra) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigExtra)(nil)).Elem()
}

func (i *ConfigExtra) ToConfigExtraOutput() ConfigExtraOutput {
	return i.ToConfigExtraOutputWithContext(context.Background())
}

func (i *ConfigExtra) ToConfigExtraOutputWithContext(ctx context.Context) ConfigExtraOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigExtraOutput)
}

// ConfigExtraArrayInput is an input type that accepts ConfigExtraArray and ConfigExtraArrayOutput values.
// You can construct a concrete instance of `ConfigExtraArrayInput` via:
//
//          ConfigExtraArray{ ConfigExtraArgs{...} }
type ConfigExtraArrayInput interface {
	pulumi.Input

	ToConfigExtraArrayOutput() ConfigExtraArrayOutput
	ToConfigExtraArrayOutputWithContext(context.Context) ConfigExtraArrayOutput
}

type ConfigExtraArray []ConfigExtraInput

func (ConfigExtraArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigExtra)(nil)).Elem()
}

func (i ConfigExtraArray) ToConfigExtraArrayOutput() ConfigExtraArrayOutput {
	return i.ToConfigExtraArrayOutputWithContext(context.Background())
}

func (i ConfigExtraArray) ToConfigExtraArrayOutputWithContext(ctx context.Context) ConfigExtraArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigExtraArrayOutput)
}

// ConfigExtraMapInput is an input type that accepts ConfigExtraMap and ConfigExtraMapOutput values.
// You can construct a concrete instance of `ConfigExtraMapInput` via:
//
//          ConfigExtraMap{ "key": ConfigExtraArgs{...} }
type ConfigExtraMapInput interface {
	pulumi.Input

	ToConfigExtraMapOutput() ConfigExtraMapOutput
	ToConfigExtraMapOutputWithContext(context.Context) ConfigExtraMapOutput
}

type ConfigExtraMap map[string]ConfigExtraInput

func (ConfigExtraMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigExtra)(nil)).Elem()
}

func (i ConfigExtraMap) ToConfigExtraMapOutput() ConfigExtraMapOutput {
	return i.ToConfigExtraMapOutputWithContext(context.Background())
}

func (i ConfigExtraMap) ToConfigExtraMapOutputWithContext(ctx context.Context) ConfigExtraMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigExtraMapOutput)
}

type ConfigExtraOutput struct{ *pulumi.OutputState }

func (ConfigExtraOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigExtra)(nil)).Elem()
}

func (o ConfigExtraOutput) ToConfigExtraOutput() ConfigExtraOutput {
	return o
}

func (o ConfigExtraOutput) ToConfigExtraOutputWithContext(ctx context.Context) ConfigExtraOutput {
	return o
}

// Collection configuration flag.
func (o ConfigExtraOutput) ConfigFlag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.ConfigFlag }).(pulumi.StringOutput)
}

// Container file path info.
func (o ConfigExtraOutput) ContainerFile() ConfigExtraContainerFilePtrOutput {
	return o.ApplyT(func(v *ConfigExtra) ConfigExtraContainerFilePtrOutput { return v.ContainerFile }).(ConfigExtraContainerFilePtrOutput)
}

// Container stdout info.
func (o ConfigExtraOutput) ContainerStdout() ConfigExtraContainerStdoutPtrOutput {
	return o.ApplyT(func(v *ConfigExtra) ConfigExtraContainerStdoutPtrOutput { return v.ContainerStdout }).(ConfigExtraContainerStdoutPtrOutput)
}

// Collection path blocklist.
func (o ConfigExtraOutput) ExcludePaths() ConfigExtraExcludePathArrayOutput {
	return o.ApplyT(func(v *ConfigExtra) ConfigExtraExcludePathArrayOutput { return v.ExcludePaths }).(ConfigExtraExcludePathArrayOutput)
}

// Extraction rule. If ExtractRule is set, LogType must be set.
func (o ConfigExtraOutput) ExtractRule() ConfigExtraExtractRulePtrOutput {
	return o.ApplyT(func(v *ConfigExtra) ConfigExtraExtractRulePtrOutput { return v.ExtractRule }).(ConfigExtraExtractRulePtrOutput)
}

// Binding group id.
func (o ConfigExtraOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// Binding group ids.
func (o ConfigExtraOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Node file config info.
func (o ConfigExtraOutput) HostFile() ConfigExtraHostFilePtrOutput {
	return o.ApplyT(func(v *ConfigExtra) ConfigExtraHostFilePtrOutput { return v.HostFile }).(ConfigExtraHostFilePtrOutput)
}

// Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format;
// minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format.
// Default value: minimalist_log.
func (o ConfigExtraOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// Logset Id.
func (o ConfigExtraOutput) LogsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.LogsetId }).(pulumi.StringOutput)
}

// Logset Name.
func (o ConfigExtraOutput) LogsetName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.LogsetName }).(pulumi.StringOutput)
}

// Collection configuration name.
func (o ConfigExtraOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Log topic ID (TopicId) of collection configuration.
func (o ConfigExtraOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

// Topic Name.
func (o ConfigExtraOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

// Type. Valid values: container_stdout; container_file; host_file.
func (o ConfigExtraOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Custom collection rule, which is a serialized JSON string.
func (o ConfigExtraOutput) UserDefineRule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigExtra) pulumi.StringPtrOutput { return v.UserDefineRule }).(pulumi.StringPtrOutput)
}

type ConfigExtraArrayOutput struct{ *pulumi.OutputState }

func (ConfigExtraArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigExtra)(nil)).Elem()
}

func (o ConfigExtraArrayOutput) ToConfigExtraArrayOutput() ConfigExtraArrayOutput {
	return o
}

func (o ConfigExtraArrayOutput) ToConfigExtraArrayOutputWithContext(ctx context.Context) ConfigExtraArrayOutput {
	return o
}

func (o ConfigExtraArrayOutput) Index(i pulumi.IntInput) ConfigExtraOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigExtra {
		return vs[0].([]*ConfigExtra)[vs[1].(int)]
	}).(ConfigExtraOutput)
}

type ConfigExtraMapOutput struct{ *pulumi.OutputState }

func (ConfigExtraMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigExtra)(nil)).Elem()
}

func (o ConfigExtraMapOutput) ToConfigExtraMapOutput() ConfigExtraMapOutput {
	return o
}

func (o ConfigExtraMapOutput) ToConfigExtraMapOutputWithContext(ctx context.Context) ConfigExtraMapOutput {
	return o
}

func (o ConfigExtraMapOutput) MapIndex(k pulumi.StringInput) ConfigExtraOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigExtra {
		return vs[0].(map[string]*ConfigExtra)[vs[1].(string)]
	}).(ConfigExtraOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigExtraInput)(nil)).Elem(), &ConfigExtra{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigExtraArrayInput)(nil)).Elem(), ConfigExtraArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigExtraMapInput)(nil)).Elem(), ConfigExtraMap{})
	pulumi.RegisterOutputType(ConfigExtraOutput{})
	pulumi.RegisterOutputType(ConfigExtraArrayOutput{})
	pulumi.RegisterOutputType(ConfigExtraMapOutput{})
}
