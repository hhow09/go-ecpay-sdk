/*
 * ECPay API
 *
 * 幸運日 API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// PaymentType the model 'PaymentType'
type PaymentType string

// List of PaymentType
const (
	PAYMENTTYPE_WEB_ATM_TAISHIN    PaymentType = "WebATM_TAISHIN"
	PAYMENTTYPE_WEB_ATM_ESUN       PaymentType = "WebATM_ESUN"
	PAYMENTTYPE_WEB_ATM_BOT        PaymentType = "WebATM_BOT"
	PAYMENTTYPE_WEB_ATM_FUBON      PaymentType = "WebATM_FUBON"
	PAYMENTTYPE_WEB_ATM_CHINATRUST PaymentType = "WebATM_CHINATRUST"
	PAYMENTTYPE_WEB_ATM_FIRST      PaymentType = "WebATM_FIRST"
	PAYMENTTYPE_WEB_ATM_CATHAY     PaymentType = "WebATM_CATHAY"
	PAYMENTTYPE_WEB_ATM_MEGA       PaymentType = "WebATM_MEGA"
	PAYMENTTYPE_WEB_ATM_LAND       PaymentType = "WebATM_LAND"
	PAYMENTTYPE_WEB_ATM_TACHONG    PaymentType = "WebATM_TACHONG"
	PAYMENTTYPE_WEB_ATM_SINOPAC    PaymentType = "WebATM_SINOPAC"
	PAYMENTTYPE_ATM_TAISHIN        PaymentType = "ATM_TAISHIN"
	PAYMENTTYPE_ATM_ESUN           PaymentType = "ATM_ESUN"
	PAYMENTTYPE_ATM_BOT            PaymentType = "ATM_BOT"
	PAYMENTTYPE_ATM_FUBON          PaymentType = "ATM_FUBON"
	PAYMENTTYPE_ATM_CHINATRUST     PaymentType = "ATM_CHINATRUST"
	PAYMENTTYPE_ATM_FIRST          PaymentType = "ATM_FIRST"
	PAYMENTTYPE_ATM_LAND           PaymentType = "ATM_LAND"
	PAYMENTTYPE_ATM_CATHAY         PaymentType = "ATM_CATHAY"
	PAYMENTTYPE_ATM_TACHONG        PaymentType = "ATM_TACHONG"
	PAYMENTTYPE_CVS_CVS            PaymentType = "CVS_CVS"
	PAYMENTTYPE_CVS_OK             PaymentType = "CVS_OK"
	PAYMENTTYPE_CVS_FAMILY         PaymentType = "CVS_FAMILY"
	PAYMENTTYPE_CVS_HILIFE         PaymentType = "CVS_HILIFE"
	PAYMENTTYPE_CVS_IBON           PaymentType = "CVS_IBON"
	PAYMENTTYPE_BARCODE_BARCODE    PaymentType = "BARCODE_BARCODE"
	PAYMENTTYPE_CREDIT_CREDIT_CARD PaymentType = "Credit_CreditCard"
)

// Ptr returns reference to PaymentType value
func (v PaymentType) Ptr() *PaymentType {
	return &v
}

type NullablePaymentType struct {
	value *PaymentType
	isSet bool
}

func (v NullablePaymentType) Get() *PaymentType {
	return v.value
}

func (v *NullablePaymentType) Set(val *PaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentType(val *PaymentType) *NullablePaymentType {
	return &NullablePaymentType{value: val, isSet: true}
}

func (v NullablePaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
