/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// ExecLogRecord **每次授權明細**   每一次的授權紀錄。   例如，若已成功授權 3 次(TotalSuccessTimes 為 3)，則會顯示 3 筆授權紀錄。   欄位內容詳見下列定期定額授權明細之參數說明。
type ExecLogRecord struct {
	// **交易狀態**   回傳值為 1 時代表授權成功，其餘為失敗，失敗代碼請參考交易訊息代碼一覽
	RtnCode int `json:"RtnCode" form:"RtnCode"`
	// **本次授權金額**  所授權的金額
	Amount int `json:"amount" form:"amount"`
	// **授權交易單號**  所授權的交易單號
	Gwsr int `json:"gwsr" form:"gwsr"`
	// **處理時間** 格式為 yyyy/MM/dd HH:mm:ss
	ProcessDate ECPayDateTime `json:"process_date" form:"process_date"`
	// **授權碼** 授權碼
	AuthCode string `json:"auth_code" form:"auth_code"`
	// **綠界的交易編號** 請保存綠界的交易編號與特店交易編號 `MerchantTradeNo` 的關連。
	TradeNo string `json:"TradeNo" form:"TradeNo"`
}

// NewExecLogRecord instantiates a new ExecLogRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecLogRecord(rtnCode int, amount int, gwsr int, processDate ECPayDateTime, authCode string, tradeNo string) *ExecLogRecord {
	this := ExecLogRecord{}
	this.RtnCode = rtnCode
	this.Amount = amount
	this.Gwsr = gwsr
	this.ProcessDate = processDate
	this.AuthCode = authCode
	this.TradeNo = tradeNo
	return &this
}

// NewExecLogRecordWithDefaults instantiates a new ExecLogRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecLogRecordWithDefaults() *ExecLogRecord {
	this := ExecLogRecord{}
	return &this
}

// GetRtnCode returns the RtnCode field value
func (o *ExecLogRecord) GetRtnCode() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RtnCode
}

// GetRtnCodeOk returns a tuple with the RtnCode field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetRtnCodeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnCode, true
}

// SetRtnCode sets field value
func (o *ExecLogRecord) SetRtnCode(v int) {
	o.RtnCode = v
}

// GetAmount returns the Amount field value
func (o *ExecLogRecord) GetAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ExecLogRecord) SetAmount(v int) {
	o.Amount = v
}

// GetGwsr returns the Gwsr field value
func (o *ExecLogRecord) GetGwsr() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Gwsr
}

// GetGwsrOk returns a tuple with the Gwsr field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetGwsrOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gwsr, true
}

// SetGwsr sets field value
func (o *ExecLogRecord) SetGwsr(v int) {
	o.Gwsr = v
}

// GetProcessDate returns the ProcessDate field value
func (o *ExecLogRecord) GetProcessDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.ProcessDate
}

// GetProcessDateOk returns a tuple with the ProcessDate field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetProcessDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessDate, true
}

// SetProcessDate sets field value
func (o *ExecLogRecord) SetProcessDate(v ECPayDateTime) {
	o.ProcessDate = v
}

// GetAuthCode returns the AuthCode field value
func (o *ExecLogRecord) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *ExecLogRecord) SetAuthCode(v string) {
	o.AuthCode = v
}

// GetTradeNo returns the TradeNo field value
func (o *ExecLogRecord) GetTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value
// and a boolean to check if the value has been set.
func (o *ExecLogRecord) GetTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeNo, true
}

// SetTradeNo sets field value
func (o *ExecLogRecord) SetTradeNo(v string) {
	o.TradeNo = v
}

func (o ExecLogRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RtnCode"] = o.RtnCode
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["gwsr"] = o.Gwsr
	}
	if true {
		toSerialize["process_date"] = o.ProcessDate
	}
	if true {
		toSerialize["auth_code"] = o.AuthCode
	}
	if true {
		toSerialize["TradeNo"] = o.TradeNo
	}
	return json.Marshal(toSerialize)
}

type NullableExecLogRecord struct {
	value *ExecLogRecord
	isSet bool
}

func (v NullableExecLogRecord) Get() *ExecLogRecord {
	return v.value
}

func (v *NullableExecLogRecord) Set(val *ExecLogRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableExecLogRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableExecLogRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecLogRecord(val *ExecLogRecord) *NullableExecLogRecord {
	return &NullableExecLogRecord{value: val, isSet: true}
}

func (v NullableExecLogRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecLogRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
