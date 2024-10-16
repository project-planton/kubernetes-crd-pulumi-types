// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KeycloakList is a list of Keycloak
type KeycloakList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of keycloaks. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KeycloakTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewKeycloakList registers a new resource with the given unique name, arguments, and options.
func NewKeycloakList(ctx *pulumi.Context,
	name string, args *KeycloakListArgs, opts ...pulumi.ResourceOption) (*KeycloakList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("k8s.keycloak.org/v2alpha1")
	args.Kind = pulumi.StringPtr("KeycloakList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KeycloakList
	err := ctx.RegisterResource("kubernetes:k8s.keycloak.org/v2alpha1:KeycloakList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeycloakList gets an existing KeycloakList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeycloakList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeycloakListState, opts ...pulumi.ResourceOption) (*KeycloakList, error) {
	var resource KeycloakList
	err := ctx.ReadResource("kubernetes:k8s.keycloak.org/v2alpha1:KeycloakList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeycloakList resources.
type keycloakListState struct {
}

type KeycloakListState struct {
}

func (KeycloakListState) ElementType() reflect.Type {
	return reflect.TypeOf((*keycloakListState)(nil)).Elem()
}

type keycloakListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of keycloaks. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []KeycloakType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a KeycloakList resource.
type KeycloakListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of keycloaks. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items KeycloakTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (KeycloakListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keycloakListArgs)(nil)).Elem()
}

type KeycloakListInput interface {
	pulumi.Input

	ToKeycloakListOutput() KeycloakListOutput
	ToKeycloakListOutputWithContext(ctx context.Context) KeycloakListOutput
}

func (*KeycloakList) ElementType() reflect.Type {
	return reflect.TypeOf((**KeycloakList)(nil)).Elem()
}

func (i *KeycloakList) ToKeycloakListOutput() KeycloakListOutput {
	return i.ToKeycloakListOutputWithContext(context.Background())
}

func (i *KeycloakList) ToKeycloakListOutputWithContext(ctx context.Context) KeycloakListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakListOutput)
}

// KeycloakListArrayInput is an input type that accepts KeycloakListArray and KeycloakListArrayOutput values.
// You can construct a concrete instance of `KeycloakListArrayInput` via:
//
//	KeycloakListArray{ KeycloakListArgs{...} }
type KeycloakListArrayInput interface {
	pulumi.Input

	ToKeycloakListArrayOutput() KeycloakListArrayOutput
	ToKeycloakListArrayOutputWithContext(context.Context) KeycloakListArrayOutput
}

type KeycloakListArray []KeycloakListInput

func (KeycloakListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeycloakList)(nil)).Elem()
}

func (i KeycloakListArray) ToKeycloakListArrayOutput() KeycloakListArrayOutput {
	return i.ToKeycloakListArrayOutputWithContext(context.Background())
}

func (i KeycloakListArray) ToKeycloakListArrayOutputWithContext(ctx context.Context) KeycloakListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakListArrayOutput)
}

// KeycloakListMapInput is an input type that accepts KeycloakListMap and KeycloakListMapOutput values.
// You can construct a concrete instance of `KeycloakListMapInput` via:
//
//	KeycloakListMap{ "key": KeycloakListArgs{...} }
type KeycloakListMapInput interface {
	pulumi.Input

	ToKeycloakListMapOutput() KeycloakListMapOutput
	ToKeycloakListMapOutputWithContext(context.Context) KeycloakListMapOutput
}

type KeycloakListMap map[string]KeycloakListInput

func (KeycloakListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeycloakList)(nil)).Elem()
}

func (i KeycloakListMap) ToKeycloakListMapOutput() KeycloakListMapOutput {
	return i.ToKeycloakListMapOutputWithContext(context.Background())
}

func (i KeycloakListMap) ToKeycloakListMapOutputWithContext(ctx context.Context) KeycloakListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeycloakListMapOutput)
}

type KeycloakListOutput struct{ *pulumi.OutputState }

func (KeycloakListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeycloakList)(nil)).Elem()
}

func (o KeycloakListOutput) ToKeycloakListOutput() KeycloakListOutput {
	return o
}

func (o KeycloakListOutput) ToKeycloakListOutputWithContext(ctx context.Context) KeycloakListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o KeycloakListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KeycloakList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of keycloaks. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o KeycloakListOutput) Items() KeycloakTypeArrayOutput {
	return o.ApplyT(func(v *KeycloakList) KeycloakTypeArrayOutput { return v.Items }).(KeycloakTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KeycloakListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KeycloakList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o KeycloakListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *KeycloakList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type KeycloakListArrayOutput struct{ *pulumi.OutputState }

func (KeycloakListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeycloakList)(nil)).Elem()
}

func (o KeycloakListArrayOutput) ToKeycloakListArrayOutput() KeycloakListArrayOutput {
	return o
}

func (o KeycloakListArrayOutput) ToKeycloakListArrayOutputWithContext(ctx context.Context) KeycloakListArrayOutput {
	return o
}

func (o KeycloakListArrayOutput) Index(i pulumi.IntInput) KeycloakListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeycloakList {
		return vs[0].([]*KeycloakList)[vs[1].(int)]
	}).(KeycloakListOutput)
}

type KeycloakListMapOutput struct{ *pulumi.OutputState }

func (KeycloakListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeycloakList)(nil)).Elem()
}

func (o KeycloakListMapOutput) ToKeycloakListMapOutput() KeycloakListMapOutput {
	return o
}

func (o KeycloakListMapOutput) ToKeycloakListMapOutputWithContext(ctx context.Context) KeycloakListMapOutput {
	return o
}

func (o KeycloakListMapOutput) MapIndex(k pulumi.StringInput) KeycloakListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeycloakList {
		return vs[0].(map[string]*KeycloakList)[vs[1].(string)]
	}).(KeycloakListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakListInput)(nil)).Elem(), &KeycloakList{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakListArrayInput)(nil)).Elem(), KeycloakListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeycloakListMapInput)(nil)).Elem(), KeycloakListMap{})
	pulumi.RegisterOutputType(KeycloakListOutput{})
	pulumi.RegisterOutputType(KeycloakListArrayOutput{})
	pulumi.RegisterOutputType(KeycloakListMapOutput{})
}
