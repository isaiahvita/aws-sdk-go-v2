// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
)

func ExampleEntityTypeFilters_outputUsage() {
	var union types.EntityTypeFilters
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EntityTypeFiltersMemberAmiProductFilters:
		_ = v.Value // Value is types.AmiProductFilters

	case *types.EntityTypeFiltersMemberContainerProductFilters:
		_ = v.Value // Value is types.ContainerProductFilters

	case *types.EntityTypeFiltersMemberDataProductFilters:
		_ = v.Value // Value is types.DataProductFilters

	case *types.EntityTypeFiltersMemberMachineLearningProductFilters:
		_ = v.Value // Value is types.MachineLearningProductFilters

	case *types.EntityTypeFiltersMemberOfferFilters:
		_ = v.Value // Value is types.OfferFilters

	case *types.EntityTypeFiltersMemberResaleAuthorizationFilters:
		_ = v.Value // Value is types.ResaleAuthorizationFilters

	case *types.EntityTypeFiltersMemberSaaSProductFilters:
		_ = v.Value // Value is types.SaaSProductFilters

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AmiProductFilters
var _ *types.ContainerProductFilters
var _ *types.SaaSProductFilters
var _ *types.OfferFilters
var _ *types.MachineLearningProductFilters
var _ *types.ResaleAuthorizationFilters
var _ *types.DataProductFilters

func ExampleEntityTypeSort_outputUsage() {
	var union types.EntityTypeSort
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EntityTypeSortMemberAmiProductSort:
		_ = v.Value // Value is types.AmiProductSort

	case *types.EntityTypeSortMemberContainerProductSort:
		_ = v.Value // Value is types.ContainerProductSort

	case *types.EntityTypeSortMemberDataProductSort:
		_ = v.Value // Value is types.DataProductSort

	case *types.EntityTypeSortMemberMachineLearningProductSort:
		_ = v.Value // Value is types.MachineLearningProductSort

	case *types.EntityTypeSortMemberOfferSort:
		_ = v.Value // Value is types.OfferSort

	case *types.EntityTypeSortMemberResaleAuthorizationSort:
		_ = v.Value // Value is types.ResaleAuthorizationSort

	case *types.EntityTypeSortMemberSaaSProductSort:
		_ = v.Value // Value is types.SaaSProductSort

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MachineLearningProductSort
var _ *types.SaaSProductSort
var _ *types.AmiProductSort
var _ *types.OfferSort
var _ *types.DataProductSort
var _ *types.ResaleAuthorizationSort
var _ *types.ContainerProductSort
