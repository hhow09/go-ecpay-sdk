/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// BindingCardEnum **記憶卡號**   使用記憶信用卡   使用：請傳 `1`     不使用：請傳 `0`
type BindingCardEnum int

// List of BindingCardEnum
const (
	BINDINGCARDENUM_BINDING     BindingCardEnum = 1
	BINDINGCARDENUM_NOT_BINDING BindingCardEnum = 0
)

// Ptr returns reference to BindingCardEnum value
func (v BindingCardEnum) Ptr() *BindingCardEnum {
	return &v
}

type NullableBindingCardEnum struct {
	value *BindingCardEnum
	isSet bool
}

func (v NullableBindingCardEnum) Get() *BindingCardEnum {
	return v.value
}

func (v *NullableBindingCardEnum) Set(val *BindingCardEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBindingCardEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBindingCardEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindingCardEnum(val *BindingCardEnum) *NullableBindingCardEnum {
	return &NullableBindingCardEnum{value: val, isSet: true}
}

func (v NullableBindingCardEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindingCardEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
