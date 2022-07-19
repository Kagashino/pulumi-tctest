// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	// Description of EKS cluster.
	ClusterDesc pulumi.StringPtrOutput `pulumi:"clusterDesc"`
	// Name of EKS cluster.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// List of cluster custom DNS Server info.
	DnsServers ClusterDnsServerArrayOutput `pulumi:"dnsServers"`
	// Indicates whether to enable dns in user cluster, default value is `true`.
	EnableVpcCoreDns pulumi.BoolPtrOutput `pulumi:"enableVpcCoreDns"`
	// Extend parameters.
	ExtraParam pulumi.MapOutput `pulumi:"extraParam"`
	// Cluster internal access LoadBalancer info.
	InternalLb ClusterInternalLbPtrOutput `pulumi:"internalLb"`
	// Kubernetes version of EKS cluster.
	K8sVersion pulumi.StringOutput `pulumi:"k8sVersion"`
	// EKS cluster kubeconfig.
	KubeConfig pulumi.StringOutput `pulumi:"kubeConfig"`
	// Delete CBS after EKS cluster remove.
	NeedDeleteCbs pulumi.BoolPtrOutput `pulumi:"needDeleteCbs"`
	// Cluster public access LoadBalancer info.
	PublicLb ClusterPublicLbPtrOutput `pulumi:"publicLb"`
	// Subnet id of service.
	ServiceSubnetId pulumi.StringPtrOutput `pulumi:"serviceSubnetId"`
	// Subnet Ids for EKS cluster.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Tags of EKS cluster.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Vpc Id of EKS cluster.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.K8sVersion == nil {
		return nil, errors.New("invalid value for required argument 'K8sVersion'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Cluster
	err := ctx.RegisterResource("tctest:Eks/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("tctest:Eks/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Description of EKS cluster.
	ClusterDesc *string `pulumi:"clusterDesc"`
	// Name of EKS cluster.
	ClusterName *string `pulumi:"clusterName"`
	// List of cluster custom DNS Server info.
	DnsServers []ClusterDnsServer `pulumi:"dnsServers"`
	// Indicates whether to enable dns in user cluster, default value is `true`.
	EnableVpcCoreDns *bool `pulumi:"enableVpcCoreDns"`
	// Extend parameters.
	ExtraParam map[string]interface{} `pulumi:"extraParam"`
	// Cluster internal access LoadBalancer info.
	InternalLb *ClusterInternalLb `pulumi:"internalLb"`
	// Kubernetes version of EKS cluster.
	K8sVersion *string `pulumi:"k8sVersion"`
	// EKS cluster kubeconfig.
	KubeConfig *string `pulumi:"kubeConfig"`
	// Delete CBS after EKS cluster remove.
	NeedDeleteCbs *bool `pulumi:"needDeleteCbs"`
	// Cluster public access LoadBalancer info.
	PublicLb *ClusterPublicLb `pulumi:"publicLb"`
	// Subnet id of service.
	ServiceSubnetId *string `pulumi:"serviceSubnetId"`
	// Subnet Ids for EKS cluster.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tags of EKS cluster.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id of EKS cluster.
	VpcId *string `pulumi:"vpcId"`
}

type ClusterState struct {
	// Description of EKS cluster.
	ClusterDesc pulumi.StringPtrInput
	// Name of EKS cluster.
	ClusterName pulumi.StringPtrInput
	// List of cluster custom DNS Server info.
	DnsServers ClusterDnsServerArrayInput
	// Indicates whether to enable dns in user cluster, default value is `true`.
	EnableVpcCoreDns pulumi.BoolPtrInput
	// Extend parameters.
	ExtraParam pulumi.MapInput
	// Cluster internal access LoadBalancer info.
	InternalLb ClusterInternalLbPtrInput
	// Kubernetes version of EKS cluster.
	K8sVersion pulumi.StringPtrInput
	// EKS cluster kubeconfig.
	KubeConfig pulumi.StringPtrInput
	// Delete CBS after EKS cluster remove.
	NeedDeleteCbs pulumi.BoolPtrInput
	// Cluster public access LoadBalancer info.
	PublicLb ClusterPublicLbPtrInput
	// Subnet id of service.
	ServiceSubnetId pulumi.StringPtrInput
	// Subnet Ids for EKS cluster.
	SubnetIds pulumi.StringArrayInput
	// Tags of EKS cluster.
	Tags pulumi.MapInput
	// Vpc Id of EKS cluster.
	VpcId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Description of EKS cluster.
	ClusterDesc *string `pulumi:"clusterDesc"`
	// Name of EKS cluster.
	ClusterName string `pulumi:"clusterName"`
	// List of cluster custom DNS Server info.
	DnsServers []ClusterDnsServer `pulumi:"dnsServers"`
	// Indicates whether to enable dns in user cluster, default value is `true`.
	EnableVpcCoreDns *bool `pulumi:"enableVpcCoreDns"`
	// Extend parameters.
	ExtraParam map[string]interface{} `pulumi:"extraParam"`
	// Cluster internal access LoadBalancer info.
	InternalLb *ClusterInternalLb `pulumi:"internalLb"`
	// Kubernetes version of EKS cluster.
	K8sVersion string `pulumi:"k8sVersion"`
	// Delete CBS after EKS cluster remove.
	NeedDeleteCbs *bool `pulumi:"needDeleteCbs"`
	// Cluster public access LoadBalancer info.
	PublicLb *ClusterPublicLb `pulumi:"publicLb"`
	// Subnet id of service.
	ServiceSubnetId *string `pulumi:"serviceSubnetId"`
	// Subnet Ids for EKS cluster.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tags of EKS cluster.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id of EKS cluster.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Description of EKS cluster.
	ClusterDesc pulumi.StringPtrInput
	// Name of EKS cluster.
	ClusterName pulumi.StringInput
	// List of cluster custom DNS Server info.
	DnsServers ClusterDnsServerArrayInput
	// Indicates whether to enable dns in user cluster, default value is `true`.
	EnableVpcCoreDns pulumi.BoolPtrInput
	// Extend parameters.
	ExtraParam pulumi.MapInput
	// Cluster internal access LoadBalancer info.
	InternalLb ClusterInternalLbPtrInput
	// Kubernetes version of EKS cluster.
	K8sVersion pulumi.StringInput
	// Delete CBS after EKS cluster remove.
	NeedDeleteCbs pulumi.BoolPtrInput
	// Cluster public access LoadBalancer info.
	PublicLb ClusterPublicLbPtrInput
	// Subnet id of service.
	ServiceSubnetId pulumi.StringPtrInput
	// Subnet Ids for EKS cluster.
	SubnetIds pulumi.StringArrayInput
	// Tags of EKS cluster.
	Tags pulumi.MapInput
	// Vpc Id of EKS cluster.
	VpcId pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// Description of EKS cluster.
func (o ClusterOutput) ClusterDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterDesc }).(pulumi.StringPtrOutput)
}

// Name of EKS cluster.
func (o ClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// List of cluster custom DNS Server info.
func (o ClusterOutput) DnsServers() ClusterDnsServerArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterDnsServerArrayOutput { return v.DnsServers }).(ClusterDnsServerArrayOutput)
}

// Indicates whether to enable dns in user cluster, default value is `true`.
func (o ClusterOutput) EnableVpcCoreDns() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableVpcCoreDns }).(pulumi.BoolPtrOutput)
}

// Extend parameters.
func (o ClusterOutput) ExtraParam() pulumi.MapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.MapOutput { return v.ExtraParam }).(pulumi.MapOutput)
}

// Cluster internal access LoadBalancer info.
func (o ClusterOutput) InternalLb() ClusterInternalLbPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterInternalLbPtrOutput { return v.InternalLb }).(ClusterInternalLbPtrOutput)
}

// Kubernetes version of EKS cluster.
func (o ClusterOutput) K8sVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.K8sVersion }).(pulumi.StringOutput)
}

// EKS cluster kubeconfig.
func (o ClusterOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.KubeConfig }).(pulumi.StringOutput)
}

// Delete CBS after EKS cluster remove.
func (o ClusterOutput) NeedDeleteCbs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.NeedDeleteCbs }).(pulumi.BoolPtrOutput)
}

// Cluster public access LoadBalancer info.
func (o ClusterOutput) PublicLb() ClusterPublicLbPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterPublicLbPtrOutput { return v.PublicLb }).(ClusterPublicLbPtrOutput)
}

// Subnet id of service.
func (o ClusterOutput) ServiceSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ServiceSubnetId }).(pulumi.StringPtrOutput)
}

// Subnet Ids for EKS cluster.
func (o ClusterOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Tags of EKS cluster.
func (o ClusterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Vpc Id of EKS cluster.
func (o ClusterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
