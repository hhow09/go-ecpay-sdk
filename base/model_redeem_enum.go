/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// RedeemEnum **信用卡是否使用紅利折抵**   設為 Y 時，當綠界特店選擇信用卡付款時，會進入紅利折抵的交易流程。   注意事項：   紅利折抵請參考信用卡紅利折抵辦法
type RedeemEnum string

// List of RedeemEnum
const (
	REDEEMENUM_Y RedeemEnum = "Y"
)

// Ptr returns reference to RedeemEnum value
func (v RedeemEnum) Ptr() *RedeemEnum {
	return &v
}

type NullableRedeemEnum struct {
	value *RedeemEnum
	isSet bool
}

func (v NullableRedeemEnum) Get() *RedeemEnum {
	return v.value
}

func (v *NullableRedeemEnum) Set(val *RedeemEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemEnum(val *RedeemEnum) *NullableRedeemEnum {
	return &NullableRedeemEnum{value: val, isSet: true}
}

func (v NullableRedeemEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
