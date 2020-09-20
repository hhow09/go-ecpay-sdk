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

// AioCheckOutCreditOnceOption struct for AioCheckOutCreditOnceOption
type AioCheckOutCreditOnceOption struct {
	Redeem   *RedeemEnum   `json:"Redeem,omitempty"`
	UnionPay *UnionPayEnum `json:"UnionPay,omitempty"`
}

// NewAioCheckOutCreditOnceOption instantiates a new AioCheckOutCreditOnceOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCreditOnceOption() *AioCheckOutCreditOnceOption {
	this := AioCheckOutCreditOnceOption{}
	var unionPay UnionPayEnum = UNIONPAYENUM_NOT_SPECIFIED
	this.UnionPay = &unionPay
	return &this
}

// NewAioCheckOutCreditOnceOptionWithDefaults instantiates a new AioCheckOutCreditOnceOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCreditOnceOptionWithDefaults() *AioCheckOutCreditOnceOption {
	this := AioCheckOutCreditOnceOption{}
	var unionPay UnionPayEnum = UNIONPAYENUM_NOT_SPECIFIED
	this.UnionPay = &unionPay
	return &this
}

// GetRedeem returns the Redeem field value if set, zero value otherwise.
func (o *AioCheckOutCreditOnceOption) GetRedeem() RedeemEnum {
	if o == nil || o.Redeem == nil {
		var ret RedeemEnum
		return ret
	}
	return *o.Redeem
}

// GetRedeemOk returns a tuple with the Redeem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOnceOption) GetRedeemOk() (*RedeemEnum, bool) {
	if o == nil || o.Redeem == nil {
		return nil, false
	}
	return o.Redeem, true
}

// HasRedeem returns a boolean if a field has been set.
func (o *AioCheckOutCreditOnceOption) HasRedeem() bool {
	if o != nil && o.Redeem != nil {
		return true
	}

	return false
}

// SetRedeem gets a reference to the given RedeemEnum and assigns it to the Redeem field.
func (o *AioCheckOutCreditOnceOption) SetRedeem(v RedeemEnum) {
	o.Redeem = &v
}

// GetUnionPay returns the UnionPay field value if set, zero value otherwise.
func (o *AioCheckOutCreditOnceOption) GetUnionPay() UnionPayEnum {
	if o == nil || o.UnionPay == nil {
		var ret UnionPayEnum
		return ret
	}
	return *o.UnionPay
}

// GetUnionPayOk returns a tuple with the UnionPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOnceOption) GetUnionPayOk() (*UnionPayEnum, bool) {
	if o == nil || o.UnionPay == nil {
		return nil, false
	}
	return o.UnionPay, true
}

// HasUnionPay returns a boolean if a field has been set.
func (o *AioCheckOutCreditOnceOption) HasUnionPay() bool {
	if o != nil && o.UnionPay != nil {
		return true
	}

	return false
}

// SetUnionPay gets a reference to the given UnionPayEnum and assigns it to the UnionPay field.
func (o *AioCheckOutCreditOnceOption) SetUnionPay(v UnionPayEnum) {
	o.UnionPay = &v
}

func (o AioCheckOutCreditOnceOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Redeem != nil {
		toSerialize["Redeem"] = o.Redeem
	}
	if o.UnionPay != nil {
		toSerialize["UnionPay"] = o.UnionPay
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCreditOnceOption struct {
	value *AioCheckOutCreditOnceOption
	isSet bool
}

func (v NullableAioCheckOutCreditOnceOption) Get() *AioCheckOutCreditOnceOption {
	return v.value
}

func (v *NullableAioCheckOutCreditOnceOption) Set(val *AioCheckOutCreditOnceOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCreditOnceOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCreditOnceOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCreditOnceOption(val *AioCheckOutCreditOnceOption) *NullableAioCheckOutCreditOnceOption {
	return &NullableAioCheckOutCreditOnceOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCreditOnceOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCreditOnceOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
