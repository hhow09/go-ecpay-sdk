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

// AioCheckOutCreditOnetimeOption struct for AioCheckOutCreditOnetimeOption
type AioCheckOutCreditOnetimeOption struct {
	Redeem   *RedeemEnum   `json:"Redeem,omitempty"`
	UnionPay *UnionPayEnum `json:"UnionPay,omitempty"`
}

// NewAioCheckOutCreditOnetimeOption instantiates a new AioCheckOutCreditOnetimeOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCreditOnetimeOption() *AioCheckOutCreditOnetimeOption {
	this := AioCheckOutCreditOnetimeOption{}
	var unionPay UnionPayEnum = UNIONPAYENUM_NOT_SPECIFIED
	this.UnionPay = &unionPay
	return &this
}

// NewAioCheckOutCreditOnetimeOptionWithDefaults instantiates a new AioCheckOutCreditOnetimeOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCreditOnetimeOptionWithDefaults() *AioCheckOutCreditOnetimeOption {
	this := AioCheckOutCreditOnetimeOption{}
	var unionPay UnionPayEnum = UNIONPAYENUM_NOT_SPECIFIED
	this.UnionPay = &unionPay
	return &this
}

// GetRedeem returns the Redeem field value if set, zero value otherwise.
func (o *AioCheckOutCreditOnetimeOption) GetRedeem() RedeemEnum {
	if o == nil || o.Redeem == nil {
		var ret RedeemEnum
		return ret
	}
	return *o.Redeem
}

// GetRedeemOk returns a tuple with the Redeem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOnetimeOption) GetRedeemOk() (*RedeemEnum, bool) {
	if o == nil || o.Redeem == nil {
		return nil, false
	}
	return o.Redeem, true
}

// HasRedeem returns a boolean if a field has been set.
func (o *AioCheckOutCreditOnetimeOption) HasRedeem() bool {
	if o != nil && o.Redeem != nil {
		return true
	}

	return false
}

// SetRedeem gets a reference to the given RedeemEnum and assigns it to the Redeem field.
func (o *AioCheckOutCreditOnetimeOption) SetRedeem(v RedeemEnum) {
	o.Redeem = &v
}

// GetUnionPay returns the UnionPay field value if set, zero value otherwise.
func (o *AioCheckOutCreditOnetimeOption) GetUnionPay() UnionPayEnum {
	if o == nil || o.UnionPay == nil {
		var ret UnionPayEnum
		return ret
	}
	return *o.UnionPay
}

// GetUnionPayOk returns a tuple with the UnionPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOnetimeOption) GetUnionPayOk() (*UnionPayEnum, bool) {
	if o == nil || o.UnionPay == nil {
		return nil, false
	}
	return o.UnionPay, true
}

// HasUnionPay returns a boolean if a field has been set.
func (o *AioCheckOutCreditOnetimeOption) HasUnionPay() bool {
	if o != nil && o.UnionPay != nil {
		return true
	}

	return false
}

// SetUnionPay gets a reference to the given UnionPayEnum and assigns it to the UnionPay field.
func (o *AioCheckOutCreditOnetimeOption) SetUnionPay(v UnionPayEnum) {
	o.UnionPay = &v
}

func (o AioCheckOutCreditOnetimeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Redeem != nil {
		toSerialize["Redeem"] = o.Redeem
	}
	if o.UnionPay != nil {
		toSerialize["UnionPay"] = o.UnionPay
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCreditOnetimeOption struct {
	value *AioCheckOutCreditOnetimeOption
	isSet bool
}

func (v NullableAioCheckOutCreditOnetimeOption) Get() *AioCheckOutCreditOnetimeOption {
	return v.value
}

func (v *NullableAioCheckOutCreditOnetimeOption) Set(val *AioCheckOutCreditOnetimeOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCreditOnetimeOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCreditOnetimeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCreditOnetimeOption(val *AioCheckOutCreditOnetimeOption) *NullableAioCheckOutCreditOnetimeOption {
	return &NullableAioCheckOutCreditOnetimeOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCreditOnetimeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCreditOnetimeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
