/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// AioCheckOutCreditInstallmentOption struct for AioCheckOutCreditInstallmentOption
type AioCheckOutCreditInstallmentOption struct {
	// **刷卡分期期數**    提供刷卡分期期數   信用卡分期可用參數為:3,6,12,18,24   注意事項：   使用的期數必須先透過申請開通後方能使用，並以申請開通的期數為主。
	CreditInstallment string `json:"CreditInstallment"`
}

// NewAioCheckOutCreditInstallmentOption instantiates a new AioCheckOutCreditInstallmentOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCreditInstallmentOption(creditInstallment string) *AioCheckOutCreditInstallmentOption {
	this := AioCheckOutCreditInstallmentOption{}
	this.CreditInstallment = creditInstallment
	return &this
}

// NewAioCheckOutCreditInstallmentOptionWithDefaults instantiates a new AioCheckOutCreditInstallmentOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCreditInstallmentOptionWithDefaults() *AioCheckOutCreditInstallmentOption {
	this := AioCheckOutCreditInstallmentOption{}
	return &this
}

// GetCreditInstallment returns the CreditInstallment field value
func (o *AioCheckOutCreditInstallmentOption) GetCreditInstallment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreditInstallment
}

// GetCreditInstallmentOk returns a tuple with the CreditInstallment field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditInstallmentOption) GetCreditInstallmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditInstallment, true
}

// SetCreditInstallment sets field value
func (o *AioCheckOutCreditInstallmentOption) SetCreditInstallment(v string) {
	o.CreditInstallment = v
}

func (o AioCheckOutCreditInstallmentOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CreditInstallment"] = o.CreditInstallment
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCreditInstallmentOption struct {
	value *AioCheckOutCreditInstallmentOption
	isSet bool
}

func (v NullableAioCheckOutCreditInstallmentOption) Get() *AioCheckOutCreditInstallmentOption {
	return v.value
}

func (v *NullableAioCheckOutCreditInstallmentOption) Set(val *AioCheckOutCreditInstallmentOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCreditInstallmentOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCreditInstallmentOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCreditInstallmentOption(val *AioCheckOutCreditInstallmentOption) *NullableAioCheckOutCreditInstallmentOption {
	return &NullableAioCheckOutCreditInstallmentOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCreditInstallmentOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCreditInstallmentOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
