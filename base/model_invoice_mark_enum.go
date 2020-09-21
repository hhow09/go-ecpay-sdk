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

// InvoiceMarkEnum **電子發票開立註記**   此參數為付款完成後同時開立電子發票。   若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」   注意事項：   正式環境欲使用電子發票功能，須與綠界申請開通，若未開通請致電客服中心 `(02) 2655-1775`。
type InvoiceMarkEnum string

// List of InvoiceMarkEnum
const (
	INVOICEMARKENUM_Y InvoiceMarkEnum = "Y"
	INVOICEMARKENUM_N InvoiceMarkEnum = "N"
)

// Ptr returns reference to InvoiceMarkEnum value
func (v InvoiceMarkEnum) Ptr() *InvoiceMarkEnum {
	return &v
}

type NullableInvoiceMarkEnum struct {
	value *InvoiceMarkEnum
	isSet bool
}

func (v NullableInvoiceMarkEnum) Get() *InvoiceMarkEnum {
	return v.value
}

func (v *NullableInvoiceMarkEnum) Set(val *InvoiceMarkEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceMarkEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceMarkEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceMarkEnum(val *InvoiceMarkEnum) *NullableInvoiceMarkEnum {
	return &NullableInvoiceMarkEnum{value: val, isSet: true}
}

func (v NullableInvoiceMarkEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceMarkEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
