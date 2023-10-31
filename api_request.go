package ecpay

import (
	"context"
	"net/http"
	"time"

	ecpayBase "github.com/hhow09/go-ecpay-sdk/base"
)

func defaultContext(contexts ...context.Context) context.Context {
	if len(contexts) == 0 {
		return context.Background()
	}
	return contexts[0]
}

func (c Client) QueryCreditCardPeriodInfo(merchantTradeNo string, timeStamp time.Time, cs ...context.Context) (ecpayBase.CreditCardPeriodInfo, *http.Response, error) {
	params := ecpayBase.QueryCreditCardPeriodInfoRequest{
		MerchantID:      c.merchantID,
		MerchantTradeNo: merchantTradeNo,
		TimeStamp:       int(timeStamp.Unix()),
	}
	params.CheckMacValue = c.GenerateCheckMacValue(StructToParamsMap(params))

	return c.apiClient.ECPayApi.
		QueryCreditCardPeriodInfo(c.WithContext(defaultContext(cs...))).
		MerchantID(params.MerchantID).
		MerchantTradeNo(params.MerchantTradeNo).
		TimeStamp(params.TimeStamp).
		CheckMacValue(params.CheckMacValue).
		Execute()
}

func (c Client) QueryTradeInfo(merchantTradeNo string, timeStamp time.Time, cs ...context.Context) (ecpayBase.TradeInfo, *http.Response, error) {
	params := ecpayBase.QueryTradeInfoRequest{
		MerchantID:      c.merchantID,
		MerchantTradeNo: merchantTradeNo,
		TimeStamp:       int(timeStamp.Unix()),
		PlatformID:      c.platformID,
		CheckMacValue:   "",
	}

	params.CheckMacValue = c.GenerateCheckMacValue(StructToParamsMap(params))

	return c.apiClient.ECPayApi.
		QueryTradeInfo(c.WithContext(defaultContext(cs...))).
		MerchantID(params.MerchantID).
		MerchantTradeNo(params.MerchantTradeNo).
		TimeStamp(params.TimeStamp).
		CheckMacValue(params.CheckMacValue).
		Execute()
}
