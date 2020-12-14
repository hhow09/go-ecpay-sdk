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

// CreditCardPeriodInfo struct for CreditCardPeriodInfo
type CreditCardPeriodInfo struct {
	// **特店編號(由綠界提供)**
	MerchantID string `json:"MerchantID"`
	// **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。
	MerchantTradeNo string `json:"MerchantTradeNo"`
	// **綠界的交易編號** 首次授權所產生的綠界交易編號
	TradeNo string `json:"TradeNo"`
	// **交易狀態**    回傳值為 1 時代表授權成功，其餘為失敗，失敗代碼請參考交易訊息代碼一覽表
	RtnCode    int                  `json:"RtnCode"`
	PeriodType CreditPeriodTypeEnum `json:"PeriodType"`
	// **執行頻率**  訂單建立時所設定的執行頻率
	Frequency int `json:"Frequency"`
	// **執行次數**  訂單建立時所設定的執行頻率
	ExecTimes int `json:"ExecTimes"`
	// **每次授權金額**  訂單建立時的每次要授權金額
	PeriodAmount int `json:"PeriodAmount"`
	// **授權金額** 所授權的金額
	Amount int `json:"amount"`
	// **授權交易單號** 所授權的交易單號
	Gwsr int `json:"gwsr"`
	// **授權成功處理時間** 格式為 yyyy/MM/dd HH:mm:ss
	ProcessDate ECPayDateTime `json:"process_date"`
	// **授權碼** 授權碼
	AuthCode string `json:"auth_code"`
	// **卡片的末 4 碼** 卡片的末四碼
	Card4no string `json:"card4no"`
	// **卡片的前 6 碼** 卡片的前六碼
	Card6no string `json:"card6no"`
	// **已成功授權次數合計** 目前已成功授權的次數
	TotalSuccessTimes int `json:"TotalSuccessTimes"`
	// **已成功授權總金額** 目前已成功授權的金額合計
	TotalSuccessAmount int             `json:"TotalSuccessAmount"`
	ExecStatus         ExecStatusEnum  `json:"ExecStatus"`
	ExecLog            []ExecLogRecord `json:"ExecLog"`
}

// NewCreditCardPeriodInfo instantiates a new CreditCardPeriodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardPeriodInfo(merchantID string, merchantTradeNo string, tradeNo string, rtnCode int, periodType CreditPeriodTypeEnum, frequency int, execTimes int, periodAmount int, amount int, gwsr int, processDate ECPayDateTime, authCode string, card4no string, card6no string, totalSuccessTimes int, totalSuccessAmount int, execStatus ExecStatusEnum, execLog []ExecLogRecord) *CreditCardPeriodInfo {
	this := CreditCardPeriodInfo{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.TradeNo = tradeNo
	this.RtnCode = rtnCode
	this.PeriodType = periodType
	this.Frequency = frequency
	this.ExecTimes = execTimes
	this.PeriodAmount = periodAmount
	this.Amount = amount
	this.Gwsr = gwsr
	this.ProcessDate = processDate
	this.AuthCode = authCode
	this.Card4no = card4no
	this.Card6no = card6no
	this.TotalSuccessTimes = totalSuccessTimes
	this.TotalSuccessAmount = totalSuccessAmount
	this.ExecStatus = execStatus
	this.ExecLog = execLog
	return &this
}

// NewCreditCardPeriodInfoWithDefaults instantiates a new CreditCardPeriodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardPeriodInfoWithDefaults() *CreditCardPeriodInfo {
	this := CreditCardPeriodInfo{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *CreditCardPeriodInfo) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *CreditCardPeriodInfo) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *CreditCardPeriodInfo) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *CreditCardPeriodInfo) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetTradeNo returns the TradeNo field value
func (o *CreditCardPeriodInfo) GetTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeNo, true
}

// SetTradeNo sets field value
func (o *CreditCardPeriodInfo) SetTradeNo(v string) {
	o.TradeNo = v
}

// GetRtnCode returns the RtnCode field value
func (o *CreditCardPeriodInfo) GetRtnCode() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RtnCode
}

// GetRtnCodeOk returns a tuple with the RtnCode field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetRtnCodeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnCode, true
}

// SetRtnCode sets field value
func (o *CreditCardPeriodInfo) SetRtnCode(v int) {
	o.RtnCode = v
}

// GetPeriodType returns the PeriodType field value
func (o *CreditCardPeriodInfo) GetPeriodType() CreditPeriodTypeEnum {
	if o == nil {
		var ret CreditPeriodTypeEnum
		return ret
	}

	return o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodType, true
}

// SetPeriodType sets field value
func (o *CreditCardPeriodInfo) SetPeriodType(v CreditPeriodTypeEnum) {
	o.PeriodType = v
}

// GetFrequency returns the Frequency field value
func (o *CreditCardPeriodInfo) GetFrequency() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetFrequencyOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *CreditCardPeriodInfo) SetFrequency(v int) {
	o.Frequency = v
}

// GetExecTimes returns the ExecTimes field value
func (o *CreditCardPeriodInfo) GetExecTimes() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.ExecTimes
}

// GetExecTimesOk returns a tuple with the ExecTimes field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetExecTimesOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecTimes, true
}

// SetExecTimes sets field value
func (o *CreditCardPeriodInfo) SetExecTimes(v int) {
	o.ExecTimes = v
}

// GetPeriodAmount returns the PeriodAmount field value
func (o *CreditCardPeriodInfo) GetPeriodAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.PeriodAmount
}

// GetPeriodAmountOk returns a tuple with the PeriodAmount field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetPeriodAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodAmount, true
}

// SetPeriodAmount sets field value
func (o *CreditCardPeriodInfo) SetPeriodAmount(v int) {
	o.PeriodAmount = v
}

// GetAmount returns the Amount field value
func (o *CreditCardPeriodInfo) GetAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreditCardPeriodInfo) SetAmount(v int) {
	o.Amount = v
}

// GetGwsr returns the Gwsr field value
func (o *CreditCardPeriodInfo) GetGwsr() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Gwsr
}

// GetGwsrOk returns a tuple with the Gwsr field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetGwsrOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gwsr, true
}

// SetGwsr sets field value
func (o *CreditCardPeriodInfo) SetGwsr(v int) {
	o.Gwsr = v
}

// GetProcessDate returns the ProcessDate field value
func (o *CreditCardPeriodInfo) GetProcessDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.ProcessDate
}

// GetProcessDateOk returns a tuple with the ProcessDate field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetProcessDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessDate, true
}

// SetProcessDate sets field value
func (o *CreditCardPeriodInfo) SetProcessDate(v ECPayDateTime) {
	o.ProcessDate = v
}

// GetAuthCode returns the AuthCode field value
func (o *CreditCardPeriodInfo) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *CreditCardPeriodInfo) SetAuthCode(v string) {
	o.AuthCode = v
}

// GetCard4no returns the Card4no field value
func (o *CreditCardPeriodInfo) GetCard4no() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Card4no
}

// GetCard4noOk returns a tuple with the Card4no field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetCard4noOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Card4no, true
}

// SetCard4no sets field value
func (o *CreditCardPeriodInfo) SetCard4no(v string) {
	o.Card4no = v
}

// GetCard6no returns the Card6no field value
func (o *CreditCardPeriodInfo) GetCard6no() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Card6no
}

// GetCard6noOk returns a tuple with the Card6no field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetCard6noOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Card6no, true
}

// SetCard6no sets field value
func (o *CreditCardPeriodInfo) SetCard6no(v string) {
	o.Card6no = v
}

// GetTotalSuccessTimes returns the TotalSuccessTimes field value
func (o *CreditCardPeriodInfo) GetTotalSuccessTimes() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TotalSuccessTimes
}

// GetTotalSuccessTimesOk returns a tuple with the TotalSuccessTimes field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetTotalSuccessTimesOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSuccessTimes, true
}

// SetTotalSuccessTimes sets field value
func (o *CreditCardPeriodInfo) SetTotalSuccessTimes(v int) {
	o.TotalSuccessTimes = v
}

// GetTotalSuccessAmount returns the TotalSuccessAmount field value
func (o *CreditCardPeriodInfo) GetTotalSuccessAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TotalSuccessAmount
}

// GetTotalSuccessAmountOk returns a tuple with the TotalSuccessAmount field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetTotalSuccessAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSuccessAmount, true
}

// SetTotalSuccessAmount sets field value
func (o *CreditCardPeriodInfo) SetTotalSuccessAmount(v int) {
	o.TotalSuccessAmount = v
}

// GetExecStatus returns the ExecStatus field value
func (o *CreditCardPeriodInfo) GetExecStatus() ExecStatusEnum {
	if o == nil {
		var ret ExecStatusEnum
		return ret
	}

	return o.ExecStatus
}

// GetExecStatusOk returns a tuple with the ExecStatus field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetExecStatusOk() (*ExecStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecStatus, true
}

// SetExecStatus sets field value
func (o *CreditCardPeriodInfo) SetExecStatus(v ExecStatusEnum) {
	o.ExecStatus = v
}

// GetExecLog returns the ExecLog field value
func (o *CreditCardPeriodInfo) GetExecLog() []ExecLogRecord {
	if o == nil {
		var ret []ExecLogRecord
		return ret
	}

	return o.ExecLog
}

// GetExecLogOk returns a tuple with the ExecLog field value
// and a boolean to check if the value has been set.
func (o *CreditCardPeriodInfo) GetExecLogOk() (*[]ExecLogRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecLog, true
}

// SetExecLog sets field value
func (o *CreditCardPeriodInfo) SetExecLog(v []ExecLogRecord) {
	o.ExecLog = v
}

func (o CreditCardPeriodInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if true {
		toSerialize["MerchantTradeNo"] = o.MerchantTradeNo
	}
	if true {
		toSerialize["TradeNo"] = o.TradeNo
	}
	if true {
		toSerialize["RtnCode"] = o.RtnCode
	}
	if true {
		toSerialize["PeriodType"] = o.PeriodType
	}
	if true {
		toSerialize["Frequency"] = o.Frequency
	}
	if true {
		toSerialize["ExecTimes"] = o.ExecTimes
	}
	if true {
		toSerialize["PeriodAmount"] = o.PeriodAmount
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
		toSerialize["card4no"] = o.Card4no
	}
	if true {
		toSerialize["card6no"] = o.Card6no
	}
	if true {
		toSerialize["TotalSuccessTimes"] = o.TotalSuccessTimes
	}
	if true {
		toSerialize["TotalSuccessAmount"] = o.TotalSuccessAmount
	}
	if true {
		toSerialize["ExecStatus"] = o.ExecStatus
	}
	if true {
		toSerialize["ExecLog"] = o.ExecLog
	}
	return json.Marshal(toSerialize)
}

type NullableCreditCardPeriodInfo struct {
	value *CreditCardPeriodInfo
	isSet bool
}

func (v NullableCreditCardPeriodInfo) Get() *CreditCardPeriodInfo {
	return v.value
}

func (v *NullableCreditCardPeriodInfo) Set(val *CreditCardPeriodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardPeriodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardPeriodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardPeriodInfo(val *CreditCardPeriodInfo) *NullableCreditCardPeriodInfo {
	return &NullableCreditCardPeriodInfo{value: val, isSet: true}
}

func (v NullableCreditCardPeriodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardPeriodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
