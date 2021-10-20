package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type outerStructWithNestedPrivateField struct {
	InnerStructWithPrivateField innerStructWithPrivateField `validate:"required"`
}

type outerStructWithNestedPublicField struct {
	InnerStructWithPublicField innerStructWithPublicField `validate:"required"`
}

type innerStructWithPublicField struct {
	PublicField string
}

type innerStructWithPrivateField struct {
	privateField string
}

func TestHasValue_InnerStructWithPrivateField(t *testing.T) {
	for _, currCase := range []struct {
		name        string
		outerStruct outerStructWithNestedPrivateField
		expectedErr bool
	}{
		// failing
		{
			name:        "missing outer struct with nested private field",
			expectedErr: true,
		},
		// failing
		{
			name:        "missing inner struct with private field",
			outerStruct: outerStructWithNestedPrivateField{},
			expectedErr: true,
		},
		// failing
		{
			name: "zero value inner struct with private field",
			outerStruct: outerStructWithNestedPrivateField{
				InnerStructWithPrivateField: innerStructWithPrivateField{},
			},
			expectedErr: true,
		},
		{
			name:        "zero value inner struct private field",
			outerStruct: outerStructWithNestedPrivateField{innerStructWithPrivateField{""}},
			expectedErr: false,
		},
		{
			name:        "valid outer and inner struct with private field",
			outerStruct: outerStructWithNestedPrivateField{innerStructWithPrivateField{"private-field"}},
			expectedErr: false,
		},
	} {
		t.Run(currCase.name, func(t *testing.T) {
			err := New().Struct(currCase.outerStruct)
			fmt.Printf("error: %v\n", err)
			if currCase.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHasValue_InnerStructWithPublicField(t *testing.T) {
	for _, currCase := range []struct {
		name        string
		outerStruct outerStructWithNestedPublicField
		expectedErr bool
	}{
		// failing
		{
			name:        "missing outer struct with nested public field",
			expectedErr: true,
		},
		// failing
		{
			name:        "missing inner struct with public field",
			outerStruct: outerStructWithNestedPublicField{},
			expectedErr: true,
		},
		// failing
		{
			name:        "zero value inner struct with public field",
			outerStruct: outerStructWithNestedPublicField{innerStructWithPublicField{}},
			expectedErr: true,
		},
		{
			name:        "zero value inner struct public field",
			outerStruct: outerStructWithNestedPublicField{innerStructWithPublicField{""}},
			expectedErr: false,
		},
		{
			name:        "valid outer and inner struct with public field",
			outerStruct: outerStructWithNestedPublicField{innerStructWithPublicField{"public-field"}},
			expectedErr: false,
		},
	} {
		t.Run(currCase.name, func(t *testing.T) {
			err := New().Struct(currCase.outerStruct)
			fmt.Printf("error: %v\n", err)
			if currCase.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
