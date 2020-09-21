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

// ReturnPaymentTypeEnum **特店選擇的付款方式** 請參考回覆付款方式一覽表
type ReturnPaymentTypeEnum string

// List of ReturnPaymentTypeEnum
const (
	RETURNPAYMENTTYPEENUM_WEB_ATM_TAISHIN    ReturnPaymentTypeEnum = "WebATM_TAISHIN"
	RETURNPAYMENTTYPEENUM_WEB_ATM_ESUN       ReturnPaymentTypeEnum = "WebATM_ESUN"
	RETURNPAYMENTTYPEENUM_WEB_ATM_BOT        ReturnPaymentTypeEnum = "WebATM_BOT"
	RETURNPAYMENTTYPEENUM_WEB_ATM_FUBON      ReturnPaymentTypeEnum = "WebATM_FUBON"
	RETURNPAYMENTTYPEENUM_WEB_ATM_CHINATRUST ReturnPaymentTypeEnum = "WebATM_CHINATRUST"
	RETURNPAYMENTTYPEENUM_WEB_ATM_FIRST      ReturnPaymentTypeEnum = "WebATM_FIRST"
	RETURNPAYMENTTYPEENUM_WEB_ATM_CATHAY     ReturnPaymentTypeEnum = "WebATM_CATHAY"
	RETURNPAYMENTTYPEENUM_WEB_ATM_MEGA       ReturnPaymentTypeEnum = "WebATM_MEGA"
	RETURNPAYMENTTYPEENUM_WEB_ATM_LAND       ReturnPaymentTypeEnum = "WebATM_LAND"
	RETURNPAYMENTTYPEENUM_WEB_ATM_TACHONG    ReturnPaymentTypeEnum = "WebATM_TACHONG"
	RETURNPAYMENTTYPEENUM_WEB_ATM_SINOPAC    ReturnPaymentTypeEnum = "WebATM_SINOPAC"
	RETURNPAYMENTTYPEENUM_ATM_TAISHIN        ReturnPaymentTypeEnum = "ATM_TAISHIN"
	RETURNPAYMENTTYPEENUM_ATM_ESUN           ReturnPaymentTypeEnum = "ATM_ESUN"
	RETURNPAYMENTTYPEENUM_ATM_BOT            ReturnPaymentTypeEnum = "ATM_BOT"
	RETURNPAYMENTTYPEENUM_ATM_FUBON          ReturnPaymentTypeEnum = "ATM_FUBON"
	RETURNPAYMENTTYPEENUM_ATM_CHINATRUST     ReturnPaymentTypeEnum = "ATM_CHINATRUST"
	RETURNPAYMENTTYPEENUM_ATM_FIRST          ReturnPaymentTypeEnum = "ATM_FIRST"
	RETURNPAYMENTTYPEENUM_ATM_LAND           ReturnPaymentTypeEnum = "ATM_LAND"
	RETURNPAYMENTTYPEENUM_ATM_CATHAY         ReturnPaymentTypeEnum = "ATM_CATHAY"
	RETURNPAYMENTTYPEENUM_ATM_TACHONG        ReturnPaymentTypeEnum = "ATM_TACHONG"
	RETURNPAYMENTTYPEENUM_CVS_CVS            ReturnPaymentTypeEnum = "CVS_CVS"
	RETURNPAYMENTTYPEENUM_CVS_OK             ReturnPaymentTypeEnum = "CVS_OK"
	RETURNPAYMENTTYPEENUM_CVS_FAMILY         ReturnPaymentTypeEnum = "CVS_FAMILY"
	RETURNPAYMENTTYPEENUM_CVS_HILIFE         ReturnPaymentTypeEnum = "CVS_HILIFE"
	RETURNPAYMENTTYPEENUM_CVS_IBON           ReturnPaymentTypeEnum = "CVS_IBON"
	RETURNPAYMENTTYPEENUM_BARCODE_BARCODE    ReturnPaymentTypeEnum = "BARCODE_BARCODE"
	RETURNPAYMENTTYPEENUM_CREDIT_CREDIT_CARD ReturnPaymentTypeEnum = "Credit_CreditCard"
)

// Ptr returns reference to ReturnPaymentTypeEnum value
func (v ReturnPaymentTypeEnum) Ptr() *ReturnPaymentTypeEnum {
	return &v
}

type NullableReturnPaymentTypeEnum struct {
	value *ReturnPaymentTypeEnum
	isSet bool
}

func (v NullableReturnPaymentTypeEnum) Get() *ReturnPaymentTypeEnum {
	return v.value
}

func (v *NullableReturnPaymentTypeEnum) Set(val *ReturnPaymentTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnPaymentTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnPaymentTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnPaymentTypeEnum(val *ReturnPaymentTypeEnum) *NullableReturnPaymentTypeEnum {
	return &NullableReturnPaymentTypeEnum{value: val, isSet: true}
}

func (v NullableReturnPaymentTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnPaymentTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
