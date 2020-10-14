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

// UnionPayEnum **銀聯卡交易選項**   可帶入以下選項:   `0`: 消費者於交易頁面可選擇是否使用銀聯交易。   `1`: 只使用銀聯卡交易，且綠界會將交易頁面直接導到銀聯網站。   `2`: 不可使用銀聯卡，綠界會將交易頁面隱藏銀聯選項。   注意事項：   1.若需使用銀聯卡服務，請與綠界提出申請方可使用，測試環境未提供銀聯卡服務。   2.不支援信用卡分期付款及定期定額。   3.不支援信用卡紅利折抵   4.不支援信用卡記憶卡號功能
type UnionPayEnum int

// List of UnionPayEnum
const (
	UNIONPAYENUM_NOT_SPECIFIED UnionPayEnum = 0
	UNIONPAYENUM_SPECIFIED     UnionPayEnum = 1
	UNIONPAYENUM_HIDDEN        UnionPayEnum = 2
)

// Ptr returns reference to UnionPayEnum value
func (v UnionPayEnum) Ptr() *UnionPayEnum {
	return &v
}

type NullableUnionPayEnum struct {
	value *UnionPayEnum
	isSet bool
}

func (v NullableUnionPayEnum) Get() *UnionPayEnum {
	return v.value
}

func (v *NullableUnionPayEnum) Set(val *UnionPayEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUnionPayEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUnionPayEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnionPayEnum(val *UnionPayEnum) *NullableUnionPayEnum {
	return &NullableUnionPayEnum{value: val, isSet: true}
}

func (v NullableUnionPayEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnionPayEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
