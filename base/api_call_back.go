/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// CallBackApiService CallBackApi service
type CallBackApiService service

type apiCallbackReturnURLPostRequest struct {
	ctx                  _context.Context
	apiService           *CallBackApiService
	merchantID           *string
	merchantTradeNo      *string
	storeID              *string
	rtnCode              *float32
	rtnMsg               *string
	tradeNo              *string
	tradeAmt             *int32
	paymentDate          *string
	paymentType          *PaymentTypeEnum
	paymentTypeChargeFee *int32
	tradeDate            *string
	checkMacValue        *string
	simulatePaid         *ECPaySimulateEnum
	customField1         *string
	customField2         *string
	customField3         *string
	customField4         *string
	customField5         *string
}

func (r apiCallbackReturnURLPostRequest) MerchantID(merchantID string) apiCallbackReturnURLPostRequest {
	r.merchantID = &merchantID
	return r
}

func (r apiCallbackReturnURLPostRequest) MerchantTradeNo(merchantTradeNo string) apiCallbackReturnURLPostRequest {
	r.merchantTradeNo = &merchantTradeNo
	return r
}

func (r apiCallbackReturnURLPostRequest) StoreID(storeID string) apiCallbackReturnURLPostRequest {
	r.storeID = &storeID
	return r
}

func (r apiCallbackReturnURLPostRequest) RtnCode(rtnCode float32) apiCallbackReturnURLPostRequest {
	r.rtnCode = &rtnCode
	return r
}

func (r apiCallbackReturnURLPostRequest) RtnMsg(rtnMsg string) apiCallbackReturnURLPostRequest {
	r.rtnMsg = &rtnMsg
	return r
}

func (r apiCallbackReturnURLPostRequest) TradeNo(tradeNo string) apiCallbackReturnURLPostRequest {
	r.tradeNo = &tradeNo
	return r
}

func (r apiCallbackReturnURLPostRequest) TradeAmt(tradeAmt int32) apiCallbackReturnURLPostRequest {
	r.tradeAmt = &tradeAmt
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentDate(paymentDate string) apiCallbackReturnURLPostRequest {
	r.paymentDate = &paymentDate
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentType(paymentType PaymentTypeEnum) apiCallbackReturnURLPostRequest {
	r.paymentType = &paymentType
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentTypeChargeFee(paymentTypeChargeFee int32) apiCallbackReturnURLPostRequest {
	r.paymentTypeChargeFee = &paymentTypeChargeFee
	return r
}

func (r apiCallbackReturnURLPostRequest) TradeDate(tradeDate string) apiCallbackReturnURLPostRequest {
	r.tradeDate = &tradeDate
	return r
}

func (r apiCallbackReturnURLPostRequest) CheckMacValue(checkMacValue string) apiCallbackReturnURLPostRequest {
	r.checkMacValue = &checkMacValue
	return r
}

func (r apiCallbackReturnURLPostRequest) SimulatePaid(simulatePaid ECPaySimulateEnum) apiCallbackReturnURLPostRequest {
	r.simulatePaid = &simulatePaid
	return r
}

func (r apiCallbackReturnURLPostRequest) CustomField1(customField1 string) apiCallbackReturnURLPostRequest {
	r.customField1 = &customField1
	return r
}

func (r apiCallbackReturnURLPostRequest) CustomField2(customField2 string) apiCallbackReturnURLPostRequest {
	r.customField2 = &customField2
	return r
}

func (r apiCallbackReturnURLPostRequest) CustomField3(customField3 string) apiCallbackReturnURLPostRequest {
	r.customField3 = &customField3
	return r
}

func (r apiCallbackReturnURLPostRequest) CustomField4(customField4 string) apiCallbackReturnURLPostRequest {
	r.customField4 = &customField4
	return r
}

func (r apiCallbackReturnURLPostRequest) CustomField5(customField5 string) apiCallbackReturnURLPostRequest {
	r.customField5 = &customField5
	return r
}

/*
CallbackReturnURLPost Method for CallbackReturnURLPost
付款結果通知
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiCallbackReturnURLPostRequest
*/
func (a *CallBackApiService) CallbackReturnURLPost(ctx _context.Context) apiCallbackReturnURLPostRequest {
	return apiCallbackReturnURLPostRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return string
*/
func (r apiCallbackReturnURLPostRequest) Execute() (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "CallBackApiService.CallbackReturnURLPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/callback/ReturnURL"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.merchantID == nil {
		return localVarReturnValue, nil, reportError("merchantID is required and must be specified")
	}

	if r.merchantTradeNo == nil {
		return localVarReturnValue, nil, reportError("merchantTradeNo is required and must be specified")
	}

	if r.storeID == nil {
		return localVarReturnValue, nil, reportError("storeID is required and must be specified")
	}

	if r.rtnCode == nil {
		return localVarReturnValue, nil, reportError("rtnCode is required and must be specified")
	}

	if r.rtnMsg == nil {
		return localVarReturnValue, nil, reportError("rtnMsg is required and must be specified")
	}

	if r.tradeNo == nil {
		return localVarReturnValue, nil, reportError("tradeNo is required and must be specified")
	}

	if r.tradeAmt == nil {
		return localVarReturnValue, nil, reportError("tradeAmt is required and must be specified")
	}
	if *r.tradeAmt < 0 {
		return localVarReturnValue, nil, reportError("tradeAmt must be greater than 0")
	}

	if r.paymentDate == nil {
		return localVarReturnValue, nil, reportError("paymentDate is required and must be specified")
	}

	if r.paymentType == nil {
		return localVarReturnValue, nil, reportError("paymentType is required and must be specified")
	}

	if r.paymentTypeChargeFee == nil {
		return localVarReturnValue, nil, reportError("paymentTypeChargeFee is required and must be specified")
	}
	if *r.paymentTypeChargeFee < 0 {
		return localVarReturnValue, nil, reportError("paymentTypeChargeFee must be greater than 0")
	}

	if r.tradeDate == nil {
		return localVarReturnValue, nil, reportError("tradeDate is required and must be specified")
	}

	if r.checkMacValue == nil {
		return localVarReturnValue, nil, reportError("checkMacValue is required and must be specified")
	}

	if r.simulatePaid == nil {
		return localVarReturnValue, nil, reportError("simulatePaid is required and must be specified")
	}

	localVarQueryParams.Add("MerchantID", parameterToString(*r.merchantID, ""))
	localVarQueryParams.Add("MerchantTradeNo", parameterToString(*r.merchantTradeNo, ""))
	localVarQueryParams.Add("StoreID", parameterToString(*r.storeID, ""))
	localVarQueryParams.Add("RtnCode", parameterToString(*r.rtnCode, ""))
	localVarQueryParams.Add("RtnMsg", parameterToString(*r.rtnMsg, ""))
	localVarQueryParams.Add("TradeNo", parameterToString(*r.tradeNo, ""))
	localVarQueryParams.Add("TradeAmt", parameterToString(*r.tradeAmt, ""))
	localVarQueryParams.Add("PaymentDate", parameterToString(*r.paymentDate, ""))
	localVarQueryParams.Add("PaymentType", parameterToString(*r.paymentType, ""))
	localVarQueryParams.Add("PaymentTypeChargeFee", parameterToString(*r.paymentTypeChargeFee, ""))
	localVarQueryParams.Add("TradeDate", parameterToString(*r.tradeDate, ""))
	localVarQueryParams.Add("CheckMacValue", parameterToString(*r.checkMacValue, ""))
	localVarQueryParams.Add("SimulatePaid", parameterToString(*r.simulatePaid, ""))
	if r.customField1 != nil {
		localVarQueryParams.Add("CustomField1", parameterToString(*r.customField1, ""))
	}
	if r.customField2 != nil {
		localVarQueryParams.Add("CustomField2", parameterToString(*r.customField2, ""))
	}
	if r.customField3 != nil {
		localVarQueryParams.Add("CustomField3", parameterToString(*r.customField3, ""))
	}
	if r.customField4 != nil {
		localVarQueryParams.Add("CustomField4", parameterToString(*r.customField4, ""))
	}
	if r.customField5 != nil {
		localVarQueryParams.Add("CustomField5", parameterToString(*r.customField5, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
