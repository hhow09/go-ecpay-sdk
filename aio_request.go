package ecpay

import (
	"github.com/Laysi/go-ecpay-sdk/base"
	"strconv"
	"strings"
	"time"
)

type AioOrderRequest struct {
	base.AioCheckOutGeneralOption
	*base.AioCheckOutAtmOption
	*base.AioCheckOutCvsBarcodeOption
	*base.AioCheckOutCreditOption
	*base.AioCheckOutCreditOnetimeOption
	*base.AioCheckOutCreditInstallmentOption
	*base.AioCheckOutCreditPeriodOption
	*base.AioCheckOutInvoiceOption
}

type AioOrderRequestWithClient struct {
	*AioOrderRequest
	client ECPayClient
}

func (e ECPayClient) CreateOrder(tradeNo string, tradeDate time.Time, amount int, description string, itemName string) *AioOrderRequestWithClient {
	return &AioOrderRequestWithClient{
		AioOrderRequest: &AioOrderRequest{
			AioCheckOutGeneralOption: base.AioCheckOutGeneralOption{
				MerchantID:        e.MerchantID,
				EncryptType:       base.ENCRYPTTYPEENUM_SHA256,
				MerchantTradeNo:   tradeNo,
				MerchantTradeDate: base.ECPayDateTime(tradeDate),
				PaymentType:       base.AIOCHECKPAYMENTTYPEENUM_AIO,
				TotalAmount:       amount,
				TradeDesc:         description,
				ItemName:          itemName,
				ReturnURL:         e.ReturnURL,
				ChoosePayment:     "",
			},
		},
		client: e,
	}
}

func (r *AioOrderRequestWithClient) WithOptional(optional AioCheckOutGeneralOptional) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.ChooseSubPayment = optional.ChooseSubPayment
	r.AioCheckOutGeneralOption.ClientBackURL = optional.ClientBackURL
	r.AioCheckOutGeneralOption.CustomField1 = optional.CustomField1
	r.AioCheckOutGeneralOption.CustomField2 = optional.CustomField2
	r.AioCheckOutGeneralOption.CustomField3 = optional.CustomField3
	r.AioCheckOutGeneralOption.CustomField4 = optional.CustomField4
	r.AioCheckOutGeneralOption.IgnorePayment = optional.IgnorePayment
	r.AioCheckOutGeneralOption.ItemURL = optional.ItemURL
	r.AioCheckOutGeneralOption.OrderResultURL = optional.OrderResultURL
	r.AioCheckOutGeneralOption.PlatformID = optional.PlatformID
	r.AioCheckOutGeneralOption.Remark = optional.Remark
	r.AioCheckOutGeneralOption.StoreID = optional.StoreID
	r.AioCheckOutGeneralOption.Language = optional.Language
	r.AioCheckOutGeneralOption.NeedExtraPaidInfo = optional.NeedExtraPaidInfo
	return r
}

func (r *AioOrderRequestWithClient) SetAtmPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_ATM
	r.AioCheckOutAtmOption = &base.AioCheckOutAtmOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithAtmOptional(option AioCheckOutAtmOptional) *AioOrderRequestWithClient {
	r.AioCheckOutAtmOption.ExpireDate = option.ExpireDate
	r.AioCheckOutAtmOption.ClientRedirectURL = option.ClientRedirectURL
	r.AioCheckOutAtmOption.PaymentInfoURL = option.PaymentInfoURL
	return r
}

func (r *AioOrderRequestWithClient) SetCvsPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_CVS
	r.AioCheckOutCvsBarcodeOption = &base.AioCheckOutCvsBarcodeOption{}
	return r
}

func (r *AioOrderRequestWithClient) SetBarcodePayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_BARCODE
	r.AioCheckOutCvsBarcodeOption = &base.AioCheckOutCvsBarcodeOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithCvsBarcodeOptional(option AioCheckOutCvsBarcodeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCvsBarcodeOption.PaymentInfoURL = option.PaymentInfoURL
	r.AioCheckOutCvsBarcodeOption.ClientRedirectURL = option.ClientRedirectURL
	r.AioCheckOutCvsBarcodeOption.Desc1 = option.Desc1
	r.AioCheckOutCvsBarcodeOption.Desc2 = option.Desc2
	r.AioCheckOutCvsBarcodeOption.Desc3 = option.Desc3
	r.AioCheckOutCvsBarcodeOption.Desc4 = option.Desc4
	r.AioCheckOutCvsBarcodeOption.StoreExpireDate = option.StoreExpireDate
	return r
}

func (r *AioOrderRequestWithClient) SetCreditPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_CREDIT
	r.AioCheckOutCreditOption = &base.AioCheckOutCreditOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOptional(option AioCheckOutCreditOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOption.BindingCard = option.BindingCard
	r.AioCheckOutCreditOption.MerchantMemberID = option.MerchantMemberID
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOnetimeOptional(option AioCheckOutCreditOnetimeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOnetimeOption = &base.AioCheckOutCreditOnetimeOption{
		Redeem:   option.Redeem,
		UnionPay: option.UnionPay,
	}
	return r
}

type IntSliceConverter []int

func (i IntSliceConverter) ToStringSlice() []string {
	var result []string
	for _, val := range i {
		result = append(result, strconv.Itoa(val))
	}
	return result
}

func (r *AioOrderRequestWithClient) WithCreditInstallmentOptional(installments []int) *AioOrderRequestWithClient {
	r.AioCheckOutCreditInstallmentOption = &base.AioCheckOutCreditInstallmentOption{
		CreditInstallment: strings.Join(IntSliceConverter(installments).ToStringSlice(), ","),
	}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditPeriodOptional(periodAmount int, periodType base.CreditPeriodTypeEnum, frequency int, execTimes int) *AioOrderRequestWithClient {
	r.AioCheckOutCreditPeriodOption = &base.AioCheckOutCreditPeriodOption{
		PeriodAmount:    periodAmount,
		PeriodType:      periodType,
		Frequency:       frequency,
		ExecTimes:       execTimes,
		PeriodReturnURL: base.PtrString(r.client.PeriodReturnURL),
	}
	return r
}

func (r *AioOrderRequestWithClient) SetAllPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_ALL
	return r
}

type InvoiceItem struct {
	Name    string
	Count   int
	Word    string
	Price   int
	TaxType *base.TaxTypeEnum
}

func (r *AioOrderRequestWithClient) SetInvoice(relateNumber string, taxType base.TaxTypeEnum, donation base.InvoiceDonationEunm, print base.InvoicePrintEnum, items []InvoiceItem, delay int, invType string) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.InvoiceMark = base.INVOICEMARKENUM_Y.Ptr()
	var itemsName []string
	var itemsCount []string
	var itemsWord []string
	var itemsPrice []string
	for _, item := range items {
		itemsName = append(itemsName, item.Name)
		itemsCount = append(itemsCount, strconv.Itoa(item.Count))
		itemsWord = append(itemsWord, item.Word)
		itemsPrice = append(itemsPrice, strconv.Itoa(item.Price))
	}
	r.AioCheckOutInvoiceOption = &base.AioCheckOutInvoiceOption{
		RelateNumber:     relateNumber,
		TaxType:          taxType,
		Donation:         donation,
		Print:            print,
		InvoiceItemName:  strings.Join(itemsName, "|"),
		InvoiceItemCount: strings.Join(itemsCount, "|"),
		InvoiceItemWord:  strings.Join(itemsWord, "|"),
		InvoiceItemPrice: strings.Join(itemsPrice, "|"),
		DelayDay:         delay,
		InvType:          invType,
	}
	if taxType == base.TAXTYPEENUM_MIXED {
		var itemsTaxType []string
		for _, item := range items {
			if taxType == base.TAXTYPEENUM_MIXED {
				itemsTaxType = append(itemsTaxType, string(*item.TaxType))

			}
		}
		r.AioCheckOutInvoiceOption.InvoiceItemTaxType = base.PtrString(strings.Join(itemsTaxType, "|"))

	}
	return r
}

func (r *AioOrderRequestWithClient) WithInvoiceOptional(option AioCheckOutInvoiceOptional) *AioOrderRequestWithClient {
	r.AioCheckOutInvoiceOption.CustomerID = option.CustomerID
	r.AioCheckOutInvoiceOption.CustomerIdentifier = option.CustomerIdentifier
	r.AioCheckOutInvoiceOption.CustomerName = option.CustomerName
	r.AioCheckOutInvoiceOption.CustomerAddr = option.CustomerAddr
	r.AioCheckOutInvoiceOption.CustomerPhone = option.CustomerPhone
	r.AioCheckOutInvoiceOption.CustomerEmail = option.CustomerEmail
	r.AioCheckOutInvoiceOption.ClearanceMark = option.ClearanceMark
	r.AioCheckOutInvoiceOption.CarruerType = option.CarruerType
	r.AioCheckOutInvoiceOption.CarruerNum = option.CarruerNum
	r.AioCheckOutInvoiceOption.LoveCode = option.LoveCode
	//r.AioCheckOutInvoiceOption.InvoiceItemTaxType = option.InvoiceItemTaxType
	r.AioCheckOutInvoiceOption.InvoiceRemark = option.InvoiceRemark
	return r
}

func (r *AioOrderRequestWithClient) GenerateUrlQuery() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return NewECPayValuesFromMap(params).Encode()

}

func (r *AioOrderRequestWithClient) GenerateRequestHtml() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return r.client.GenerateAutoSubmitHtmlForm(params, r.client.GetCurrentServer()+r.client.AioCheckOutPath)

}
