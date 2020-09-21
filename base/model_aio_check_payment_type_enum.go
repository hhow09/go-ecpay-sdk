/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// AioCheckPaymentTypeEnum **交易類型** 請固定填入 `aio`
type AioCheckPaymentTypeEnum string

// List of AioCheckPaymentTypeEnum
const (
	AIOCHECKPAYMENTTYPEENUM_AIO AioCheckPaymentTypeEnum = "aio"
)

// Ptr returns reference to AioCheckPaymentTypeEnum value
func (v AioCheckPaymentTypeEnum) Ptr() *AioCheckPaymentTypeEnum {
	return &v
}

type NullableAioCheckPaymentTypeEnum struct {
	value *AioCheckPaymentTypeEnum
	isSet bool
}

func (v NullableAioCheckPaymentTypeEnum) Get() *AioCheckPaymentTypeEnum {
	return v.value
}

func (v *NullableAioCheckPaymentTypeEnum) Set(val *AioCheckPaymentTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckPaymentTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckPaymentTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckPaymentTypeEnum(val *AioCheckPaymentTypeEnum) *NullableAioCheckPaymentTypeEnum {
	return &NullableAioCheckPaymentTypeEnum{value: val, isSet: true}
}

func (v NullableAioCheckPaymentTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckPaymentTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
