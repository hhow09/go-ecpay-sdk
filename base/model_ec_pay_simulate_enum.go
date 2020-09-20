/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.9
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// ECPaySimulateEnum the model 'ECPaySimulateEnum'
type ECPaySimulateEnum int32

// List of ECPaySimulateEnum
const (
	ECPAYSIMULATEENUM_REAL_PAID ECPaySimulateEnum = 0
	ECPAYSIMULATEENUM_SIMULATED ECPaySimulateEnum = 1
)

// Ptr returns reference to ECPaySimulateEnum value
func (v ECPaySimulateEnum) Ptr() *ECPaySimulateEnum {
	return &v
}

type NullableECPaySimulateEnum struct {
	value *ECPaySimulateEnum
	isSet bool
}

func (v NullableECPaySimulateEnum) Get() *ECPaySimulateEnum {
	return v.value
}

func (v *NullableECPaySimulateEnum) Set(val *ECPaySimulateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableECPaySimulateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableECPaySimulateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECPaySimulateEnum(val *ECPaySimulateEnum) *NullableECPaySimulateEnum {
	return &NullableECPaySimulateEnum{value: val, isSet: true}
}

func (v NullableECPaySimulateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECPaySimulateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
