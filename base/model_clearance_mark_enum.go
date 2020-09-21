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

// ClearanceMarkEnum **通關方式**   當課稅類別`TaxType`為 `2`(零稅率)時，則該參數請帶 `1`(非經海關出口)或 `2`(經海關出口)。
type ClearanceMarkEnum string

// List of ClearanceMarkEnum
const (
	CLEARANCEMARKENUM_NOT_CUSTOMS ClearanceMarkEnum = "1"
	CLEARANCEMARKENUM_CUSTOMS     ClearanceMarkEnum = "2"
)

// Ptr returns reference to ClearanceMarkEnum value
func (v ClearanceMarkEnum) Ptr() *ClearanceMarkEnum {
	return &v
}

type NullableClearanceMarkEnum struct {
	value *ClearanceMarkEnum
	isSet bool
}

func (v NullableClearanceMarkEnum) Get() *ClearanceMarkEnum {
	return v.value
}

func (v *NullableClearanceMarkEnum) Set(val *ClearanceMarkEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClearanceMarkEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClearanceMarkEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearanceMarkEnum(val *ClearanceMarkEnum) *NullableClearanceMarkEnum {
	return &NullableClearanceMarkEnum{value: val, isSet: true}
}

func (v NullableClearanceMarkEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearanceMarkEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
