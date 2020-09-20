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

// AioCheckOutInvoiceOption struct for AioCheckOutInvoiceOption
type AioCheckOutInvoiceOption struct {
	// **特店自訂編號** 此為特店自訂編號，編號均為唯一值不可重複使用。
	RelateNumber string `json:"RelateNumber"`
	// **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。
	CustomerID *string `json:"CustomerID,omitempty"`
	// **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。
	CustomerIdentifier *string `json:"CustomerIdentifier,omitempty"`
	// **客戶名稱**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。
	CustomerName *string `json:"CustomerName,omitempty"`
	// **客戶地址**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。
	CustomerAddr *string `json:"CustomerAddr,omitempty"`
	// **客戶手機號碼**   當客戶電子信箱`CustomerEmail`為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用
	CustomerPhone *string `json:"CustomerPhone,omitempty"`
	// **客戶電子信箱**   當客戶手機號碼`CustomerPhone`為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。
	CustomerEmail *string            `json:"CustomerEmail,omitempty"`
	ClearanceMark *ClearanceMarkEnum `json:"ClearanceMark,omitempty"`
	TaxType       TaxTypeEnum        `json:"TaxType"`
	CarruerType   *CarruerTypeEnum   `json:"CarruerType,omitempty"`
	// **載具編號**   1. 當載具類別 `CarruerType`=``無載具)，請帶空字串。   2. 當載具類別`CarruerType`=`1`(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別`CarruerType`=`2`(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別`CarruerType`=`3`(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核
	CarruerNum *string             `json:"CarruerNum,omitempty"`
	Donation   InvoiceDonationEunm `json:"Donation"`
	// **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 `Donation`= `1` (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。
	LoveCode *string          `json:"LoveCode,omitempty"`
	Print    InvoicePrintEnum `json:"Print"`
	// **商品名稱**   預設不可為空字串且格式為 名稱 1 | 名稱 2 | 名稱 3 | … | 名稱 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。   將參數值以 UrlEncode 方式編碼。
	InvoiceItemName string `json:"InvoiceItemName"`
	// **商品數量**   預設不可為空字串且格式為 數量 1 | 數量 2 | 數量 3 | … | 數量 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。
	InvoiceItemCount string `json:"InvoiceItemCount"`
	// **商品單位**   商品單位若超過二筆以上請以「|」符號區隔單位最大長度為 6 碼。   請將參數做 UrlEncode 方式編碼。
	InvoiceItemWord string `json:"InvoiceItemWord"`
	// **商品價格**   預設不可為空字串且格式為 價格 1 | 價格 2 | 價格 3 | … | 價格 n，當含有二筆或以上的商品價格時，則以「|」符號區隔。
	InvoiceItemPrice string `json:"InvoiceItemPrice"`
	// **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] = 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。
	InvoiceItemTaxType *string `json:"InvoiceItemTaxType,omitempty"`
	// **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。
	InvoiceRemark *string `json:"InvoiceRemark,omitempty"`
	// **延遲天數**   本參數值請帶 0~15(天)，當天數為 0 時，則付款完成後立即開立發票。
	DelayDay string `json:"DelayDay"`
	// **字軌類別**   若為一般稅額時，請帶 07。   預設值：07
	InvType string `json:"InvType"`
}

// NewAioCheckOutInvoiceOption instantiates a new AioCheckOutInvoiceOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutInvoiceOption(relateNumber string, taxType TaxTypeEnum, donation InvoiceDonationEunm, print InvoicePrintEnum, invoiceItemName string, invoiceItemCount string, invoiceItemWord string, invoiceItemPrice string, delayDay string, invType string) *AioCheckOutInvoiceOption {
	this := AioCheckOutInvoiceOption{}
	this.RelateNumber = relateNumber
	this.TaxType = taxType
	this.Donation = donation
	this.Print = print
	this.InvoiceItemName = invoiceItemName
	this.InvoiceItemCount = invoiceItemCount
	this.InvoiceItemWord = invoiceItemWord
	this.InvoiceItemPrice = invoiceItemPrice
	this.DelayDay = delayDay
	this.InvType = invType
	return &this
}

// NewAioCheckOutInvoiceOptionWithDefaults instantiates a new AioCheckOutInvoiceOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutInvoiceOptionWithDefaults() *AioCheckOutInvoiceOption {
	this := AioCheckOutInvoiceOption{}
	return &this
}

// GetRelateNumber returns the RelateNumber field value
func (o *AioCheckOutInvoiceOption) GetRelateNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelateNumber
}

// GetRelateNumberOk returns a tuple with the RelateNumber field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetRelateNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelateNumber, true
}

// SetRelateNumber sets field value
func (o *AioCheckOutInvoiceOption) SetRelateNumber(v string) {
	o.RelateNumber = v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *AioCheckOutInvoiceOption) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetCustomerIdentifier returns the CustomerIdentifier field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerIdentifier() string {
	if o == nil || o.CustomerIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CustomerIdentifier
}

// GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerIdentifierOk() (*string, bool) {
	if o == nil || o.CustomerIdentifier == nil {
		return nil, false
	}
	return o.CustomerIdentifier, true
}

// HasCustomerIdentifier returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerIdentifier() bool {
	if o != nil && o.CustomerIdentifier != nil {
		return true
	}

	return false
}

// SetCustomerIdentifier gets a reference to the given string and assigns it to the CustomerIdentifier field.
func (o *AioCheckOutInvoiceOption) SetCustomerIdentifier(v string) {
	o.CustomerIdentifier = &v
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerName() string {
	if o == nil || o.CustomerName == nil {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerNameOk() (*string, bool) {
	if o == nil || o.CustomerName == nil {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerName() bool {
	if o != nil && o.CustomerName != nil {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *AioCheckOutInvoiceOption) SetCustomerName(v string) {
	o.CustomerName = &v
}

// GetCustomerAddr returns the CustomerAddr field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerAddr() string {
	if o == nil || o.CustomerAddr == nil {
		var ret string
		return ret
	}
	return *o.CustomerAddr
}

// GetCustomerAddrOk returns a tuple with the CustomerAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerAddrOk() (*string, bool) {
	if o == nil || o.CustomerAddr == nil {
		return nil, false
	}
	return o.CustomerAddr, true
}

// HasCustomerAddr returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerAddr() bool {
	if o != nil && o.CustomerAddr != nil {
		return true
	}

	return false
}

// SetCustomerAddr gets a reference to the given string and assigns it to the CustomerAddr field.
func (o *AioCheckOutInvoiceOption) SetCustomerAddr(v string) {
	o.CustomerAddr = &v
}

// GetCustomerPhone returns the CustomerPhone field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerPhone() string {
	if o == nil || o.CustomerPhone == nil {
		var ret string
		return ret
	}
	return *o.CustomerPhone
}

// GetCustomerPhoneOk returns a tuple with the CustomerPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerPhoneOk() (*string, bool) {
	if o == nil || o.CustomerPhone == nil {
		return nil, false
	}
	return o.CustomerPhone, true
}

// HasCustomerPhone returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerPhone() bool {
	if o != nil && o.CustomerPhone != nil {
		return true
	}

	return false
}

// SetCustomerPhone gets a reference to the given string and assigns it to the CustomerPhone field.
func (o *AioCheckOutInvoiceOption) SetCustomerPhone(v string) {
	o.CustomerPhone = &v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *AioCheckOutInvoiceOption) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetClearanceMark returns the ClearanceMark field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetClearanceMark() ClearanceMarkEnum {
	if o == nil || o.ClearanceMark == nil {
		var ret ClearanceMarkEnum
		return ret
	}
	return *o.ClearanceMark
}

// GetClearanceMarkOk returns a tuple with the ClearanceMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetClearanceMarkOk() (*ClearanceMarkEnum, bool) {
	if o == nil || o.ClearanceMark == nil {
		return nil, false
	}
	return o.ClearanceMark, true
}

// HasClearanceMark returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasClearanceMark() bool {
	if o != nil && o.ClearanceMark != nil {
		return true
	}

	return false
}

// SetClearanceMark gets a reference to the given ClearanceMarkEnum and assigns it to the ClearanceMark field.
func (o *AioCheckOutInvoiceOption) SetClearanceMark(v ClearanceMarkEnum) {
	o.ClearanceMark = &v
}

// GetTaxType returns the TaxType field value
func (o *AioCheckOutInvoiceOption) GetTaxType() TaxTypeEnum {
	if o == nil {
		var ret TaxTypeEnum
		return ret
	}

	return o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetTaxTypeOk() (*TaxTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxType, true
}

// SetTaxType sets field value
func (o *AioCheckOutInvoiceOption) SetTaxType(v TaxTypeEnum) {
	o.TaxType = v
}

// GetCarruerType returns the CarruerType field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCarruerType() CarruerTypeEnum {
	if o == nil || o.CarruerType == nil {
		var ret CarruerTypeEnum
		return ret
	}
	return *o.CarruerType
}

// GetCarruerTypeOk returns a tuple with the CarruerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCarruerTypeOk() (*CarruerTypeEnum, bool) {
	if o == nil || o.CarruerType == nil {
		return nil, false
	}
	return o.CarruerType, true
}

// HasCarruerType returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCarruerType() bool {
	if o != nil && o.CarruerType != nil {
		return true
	}

	return false
}

// SetCarruerType gets a reference to the given CarruerTypeEnum and assigns it to the CarruerType field.
func (o *AioCheckOutInvoiceOption) SetCarruerType(v CarruerTypeEnum) {
	o.CarruerType = &v
}

// GetCarruerNum returns the CarruerNum field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetCarruerNum() string {
	if o == nil || o.CarruerNum == nil {
		var ret string
		return ret
	}
	return *o.CarruerNum
}

// GetCarruerNumOk returns a tuple with the CarruerNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetCarruerNumOk() (*string, bool) {
	if o == nil || o.CarruerNum == nil {
		return nil, false
	}
	return o.CarruerNum, true
}

// HasCarruerNum returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasCarruerNum() bool {
	if o != nil && o.CarruerNum != nil {
		return true
	}

	return false
}

// SetCarruerNum gets a reference to the given string and assigns it to the CarruerNum field.
func (o *AioCheckOutInvoiceOption) SetCarruerNum(v string) {
	o.CarruerNum = &v
}

// GetDonation returns the Donation field value
func (o *AioCheckOutInvoiceOption) GetDonation() InvoiceDonationEunm {
	if o == nil {
		var ret InvoiceDonationEunm
		return ret
	}

	return o.Donation
}

// GetDonationOk returns a tuple with the Donation field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetDonationOk() (*InvoiceDonationEunm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Donation, true
}

// SetDonation sets field value
func (o *AioCheckOutInvoiceOption) SetDonation(v InvoiceDonationEunm) {
	o.Donation = v
}

// GetLoveCode returns the LoveCode field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetLoveCode() string {
	if o == nil || o.LoveCode == nil {
		var ret string
		return ret
	}
	return *o.LoveCode
}

// GetLoveCodeOk returns a tuple with the LoveCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetLoveCodeOk() (*string, bool) {
	if o == nil || o.LoveCode == nil {
		return nil, false
	}
	return o.LoveCode, true
}

// HasLoveCode returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasLoveCode() bool {
	if o != nil && o.LoveCode != nil {
		return true
	}

	return false
}

// SetLoveCode gets a reference to the given string and assigns it to the LoveCode field.
func (o *AioCheckOutInvoiceOption) SetLoveCode(v string) {
	o.LoveCode = &v
}

// GetPrint returns the Print field value
func (o *AioCheckOutInvoiceOption) GetPrint() InvoicePrintEnum {
	if o == nil {
		var ret InvoicePrintEnum
		return ret
	}

	return o.Print
}

// GetPrintOk returns a tuple with the Print field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetPrintOk() (*InvoicePrintEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Print, true
}

// SetPrint sets field value
func (o *AioCheckOutInvoiceOption) SetPrint(v InvoicePrintEnum) {
	o.Print = v
}

// GetInvoiceItemName returns the InvoiceItemName field value
func (o *AioCheckOutInvoiceOption) GetInvoiceItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceItemName
}

// GetInvoiceItemNameOk returns a tuple with the InvoiceItemName field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceItemName, true
}

// SetInvoiceItemName sets field value
func (o *AioCheckOutInvoiceOption) SetInvoiceItemName(v string) {
	o.InvoiceItemName = v
}

// GetInvoiceItemCount returns the InvoiceItemCount field value
func (o *AioCheckOutInvoiceOption) GetInvoiceItemCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceItemCount
}

// GetInvoiceItemCountOk returns a tuple with the InvoiceItemCount field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceItemCount, true
}

// SetInvoiceItemCount sets field value
func (o *AioCheckOutInvoiceOption) SetInvoiceItemCount(v string) {
	o.InvoiceItemCount = v
}

// GetInvoiceItemWord returns the InvoiceItemWord field value
func (o *AioCheckOutInvoiceOption) GetInvoiceItemWord() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceItemWord
}

// GetInvoiceItemWordOk returns a tuple with the InvoiceItemWord field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceItemWord, true
}

// SetInvoiceItemWord sets field value
func (o *AioCheckOutInvoiceOption) SetInvoiceItemWord(v string) {
	o.InvoiceItemWord = v
}

// GetInvoiceItemPrice returns the InvoiceItemPrice field value
func (o *AioCheckOutInvoiceOption) GetInvoiceItemPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceItemPrice
}

// GetInvoiceItemPriceOk returns a tuple with the InvoiceItemPrice field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceItemPrice, true
}

// SetInvoiceItemPrice sets field value
func (o *AioCheckOutInvoiceOption) SetInvoiceItemPrice(v string) {
	o.InvoiceItemPrice = v
}

// GetInvoiceItemTaxType returns the InvoiceItemTaxType field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemTaxType() string {
	if o == nil || o.InvoiceItemTaxType == nil {
		var ret string
		return ret
	}
	return *o.InvoiceItemTaxType
}

// GetInvoiceItemTaxTypeOk returns a tuple with the InvoiceItemTaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceItemTaxTypeOk() (*string, bool) {
	if o == nil || o.InvoiceItemTaxType == nil {
		return nil, false
	}
	return o.InvoiceItemTaxType, true
}

// HasInvoiceItemTaxType returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasInvoiceItemTaxType() bool {
	if o != nil && o.InvoiceItemTaxType != nil {
		return true
	}

	return false
}

// SetInvoiceItemTaxType gets a reference to the given string and assigns it to the InvoiceItemTaxType field.
func (o *AioCheckOutInvoiceOption) SetInvoiceItemTaxType(v string) {
	o.InvoiceItemTaxType = &v
}

// GetInvoiceRemark returns the InvoiceRemark field value if set, zero value otherwise.
func (o *AioCheckOutInvoiceOption) GetInvoiceRemark() string {
	if o == nil || o.InvoiceRemark == nil {
		var ret string
		return ret
	}
	return *o.InvoiceRemark
}

// GetInvoiceRemarkOk returns a tuple with the InvoiceRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvoiceRemarkOk() (*string, bool) {
	if o == nil || o.InvoiceRemark == nil {
		return nil, false
	}
	return o.InvoiceRemark, true
}

// HasInvoiceRemark returns a boolean if a field has been set.
func (o *AioCheckOutInvoiceOption) HasInvoiceRemark() bool {
	if o != nil && o.InvoiceRemark != nil {
		return true
	}

	return false
}

// SetInvoiceRemark gets a reference to the given string and assigns it to the InvoiceRemark field.
func (o *AioCheckOutInvoiceOption) SetInvoiceRemark(v string) {
	o.InvoiceRemark = &v
}

// GetDelayDay returns the DelayDay field value
func (o *AioCheckOutInvoiceOption) GetDelayDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DelayDay
}

// GetDelayDayOk returns a tuple with the DelayDay field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetDelayDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelayDay, true
}

// SetDelayDay sets field value
func (o *AioCheckOutInvoiceOption) SetDelayDay(v string) {
	o.DelayDay = v
}

// GetInvType returns the InvType field value
func (o *AioCheckOutInvoiceOption) GetInvType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvType
}

// GetInvTypeOk returns a tuple with the InvType field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutInvoiceOption) GetInvTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvType, true
}

// SetInvType sets field value
func (o *AioCheckOutInvoiceOption) SetInvType(v string) {
	o.InvType = v
}

func (o AioCheckOutInvoiceOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RelateNumber"] = o.RelateNumber
	}
	if o.CustomerID != nil {
		toSerialize["CustomerID"] = o.CustomerID
	}
	if o.CustomerIdentifier != nil {
		toSerialize["CustomerIdentifier"] = o.CustomerIdentifier
	}
	if o.CustomerName != nil {
		toSerialize["CustomerName"] = o.CustomerName
	}
	if o.CustomerAddr != nil {
		toSerialize["CustomerAddr"] = o.CustomerAddr
	}
	if o.CustomerPhone != nil {
		toSerialize["CustomerPhone"] = o.CustomerPhone
	}
	if o.CustomerEmail != nil {
		toSerialize["CustomerEmail"] = o.CustomerEmail
	}
	if o.ClearanceMark != nil {
		toSerialize["ClearanceMark"] = o.ClearanceMark
	}
	if true {
		toSerialize["TaxType"] = o.TaxType
	}
	if o.CarruerType != nil {
		toSerialize["CarruerType"] = o.CarruerType
	}
	if o.CarruerNum != nil {
		toSerialize["CarruerNum"] = o.CarruerNum
	}
	if true {
		toSerialize["Donation"] = o.Donation
	}
	if o.LoveCode != nil {
		toSerialize["LoveCode"] = o.LoveCode
	}
	if true {
		toSerialize["Print"] = o.Print
	}
	if true {
		toSerialize["InvoiceItemName"] = o.InvoiceItemName
	}
	if true {
		toSerialize["InvoiceItemCount"] = o.InvoiceItemCount
	}
	if true {
		toSerialize["InvoiceItemWord"] = o.InvoiceItemWord
	}
	if true {
		toSerialize["InvoiceItemPrice"] = o.InvoiceItemPrice
	}
	if o.InvoiceItemTaxType != nil {
		toSerialize["InvoiceItemTaxType"] = o.InvoiceItemTaxType
	}
	if o.InvoiceRemark != nil {
		toSerialize["InvoiceRemark"] = o.InvoiceRemark
	}
	if true {
		toSerialize["DelayDay"] = o.DelayDay
	}
	if true {
		toSerialize["InvType"] = o.InvType
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutInvoiceOption struct {
	value *AioCheckOutInvoiceOption
	isSet bool
}

func (v NullableAioCheckOutInvoiceOption) Get() *AioCheckOutInvoiceOption {
	return v.value
}

func (v *NullableAioCheckOutInvoiceOption) Set(val *AioCheckOutInvoiceOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutInvoiceOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutInvoiceOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutInvoiceOption(val *AioCheckOutInvoiceOption) *NullableAioCheckOutInvoiceOption {
	return &NullableAioCheckOutInvoiceOption{value: val, isSet: true}
}

func (v NullableAioCheckOutInvoiceOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutInvoiceOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
