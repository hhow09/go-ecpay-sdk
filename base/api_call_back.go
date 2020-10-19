/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.21
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"bytes"
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

type apiCallbackPeriodReturnURLPostRequest struct {
	ctx               _context.Context
	apiService        *CallBackApiService
	merchantID        *string
	merchantTradeNo   *string
	storeID           *string
	rtnCode           *int
	rtnMsg            *string
	periodType        *CreditPeriodTypeEnum
	frequency         *int
	execTimes         *int
	amount            *int
	gwsr              *int
	processDate       *ECPayDateTime
	authCode          *string
	firstAuthAmount   *int
	totalSuccessTimes *int
	checkMacValue     *string
	simulatePaid      *SimulatePaidEnum
	customField1      *string
	customField2      *string
	customField3      *string
	customField4      *string
}

func (r apiCallbackPeriodReturnURLPostRequest) MerchantID(merchantID string) apiCallbackPeriodReturnURLPostRequest {
	r.merchantID = &merchantID
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) MerchantTradeNo(merchantTradeNo string) apiCallbackPeriodReturnURLPostRequest {
	r.merchantTradeNo = &merchantTradeNo
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) StoreID(storeID string) apiCallbackPeriodReturnURLPostRequest {
	r.storeID = &storeID
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) RtnCode(rtnCode int) apiCallbackPeriodReturnURLPostRequest {
	r.rtnCode = &rtnCode
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) RtnMsg(rtnMsg string) apiCallbackPeriodReturnURLPostRequest {
	r.rtnMsg = &rtnMsg
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) PeriodType(periodType CreditPeriodTypeEnum) apiCallbackPeriodReturnURLPostRequest {
	r.periodType = &periodType
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) Frequency(frequency int) apiCallbackPeriodReturnURLPostRequest {
	r.frequency = &frequency
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) ExecTimes(execTimes int) apiCallbackPeriodReturnURLPostRequest {
	r.execTimes = &execTimes
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) Amount(amount int) apiCallbackPeriodReturnURLPostRequest {
	r.amount = &amount
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) Gwsr(gwsr int) apiCallbackPeriodReturnURLPostRequest {
	r.gwsr = &gwsr
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) ProcessDate(processDate ECPayDateTime) apiCallbackPeriodReturnURLPostRequest {
	r.processDate = &processDate
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) AuthCode(authCode string) apiCallbackPeriodReturnURLPostRequest {
	r.authCode = &authCode
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) FirstAuthAmount(firstAuthAmount int) apiCallbackPeriodReturnURLPostRequest {
	r.firstAuthAmount = &firstAuthAmount
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) TotalSuccessTimes(totalSuccessTimes int) apiCallbackPeriodReturnURLPostRequest {
	r.totalSuccessTimes = &totalSuccessTimes
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) CheckMacValue(checkMacValue string) apiCallbackPeriodReturnURLPostRequest {
	r.checkMacValue = &checkMacValue
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) SimulatePaid(simulatePaid SimulatePaidEnum) apiCallbackPeriodReturnURLPostRequest {
	r.simulatePaid = &simulatePaid
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) CustomField1(customField1 string) apiCallbackPeriodReturnURLPostRequest {
	r.customField1 = &customField1
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) CustomField2(customField2 string) apiCallbackPeriodReturnURLPostRequest {
	r.customField2 = &customField2
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) CustomField3(customField3 string) apiCallbackPeriodReturnURLPostRequest {
	r.customField3 = &customField3
	return r
}

func (r apiCallbackPeriodReturnURLPostRequest) CustomField4(customField4 string) apiCallbackPeriodReturnURLPostRequest {
	r.customField4 = &customField4
	return r
}

/*
CallbackPeriodReturnURLPost Method for CallbackPeriodReturnURLPost
付款結果通知
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiCallbackPeriodReturnURLPostRequest
*/
func (a *CallBackApiService) CallbackPeriodReturnURLPost(ctx _context.Context) apiCallbackPeriodReturnURLPostRequest {
	return apiCallbackPeriodReturnURLPostRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return string
*/
func (r apiCallbackPeriodReturnURLPostRequest) Execute() (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "CallBackApiService.CallbackPeriodReturnURLPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/callback/PeriodReturnURL"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.merchantID == nil {
		return localVarReturnValue, nil, reportError("merchantID is required and must be specified")
	}
	if strlen(*r.merchantID) > 10 {
		return localVarReturnValue, nil, reportError("merchantID must have less than 10 elements")
	}

	if r.merchantTradeNo == nil {
		return localVarReturnValue, nil, reportError("merchantTradeNo is required and must be specified")
	}
	if strlen(*r.merchantTradeNo) > 20 {
		return localVarReturnValue, nil, reportError("merchantTradeNo must have less than 20 elements")
	}

	if r.storeID == nil {
		return localVarReturnValue, nil, reportError("storeID is required and must be specified")
	}
	if strlen(*r.storeID) > 20 {
		return localVarReturnValue, nil, reportError("storeID must have less than 20 elements")
	}

	if r.rtnCode == nil {
		return localVarReturnValue, nil, reportError("rtnCode is required and must be specified")
	}

	if r.rtnMsg == nil {
		return localVarReturnValue, nil, reportError("rtnMsg is required and must be specified")
	}
	if strlen(*r.rtnMsg) > 200 {
		return localVarReturnValue, nil, reportError("rtnMsg must have less than 200 elements")
	}

	if r.periodType == nil {
		return localVarReturnValue, nil, reportError("periodType is required and must be specified")
	}

	if r.frequency == nil {
		return localVarReturnValue, nil, reportError("frequency is required and must be specified")
	}
	if *r.frequency < 1 {
		return localVarReturnValue, nil, reportError("frequency must be greater than 1")
	}
	if *r.frequency > 365 {
		return localVarReturnValue, nil, reportError("frequency must be less than 365")
	}

	if r.execTimes == nil {
		return localVarReturnValue, nil, reportError("execTimes is required and must be specified")
	}
	if *r.execTimes < 1 {
		return localVarReturnValue, nil, reportError("execTimes must be greater than 1")
	}
	if *r.execTimes > 999 {
		return localVarReturnValue, nil, reportError("execTimes must be less than 999")
	}

	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}
	if *r.amount < 0 {
		return localVarReturnValue, nil, reportError("amount must be greater than 0")
	}

	if r.gwsr == nil {
		return localVarReturnValue, nil, reportError("gwsr is required and must be specified")
	}

	if r.processDate == nil {
		return localVarReturnValue, nil, reportError("processDate is required and must be specified")
	}

	if r.authCode == nil {
		return localVarReturnValue, nil, reportError("authCode is required and must be specified")
	}
	authCodeTxt, err := atoi(*r.authCode)
	if authCodeTxt > 6 {
		return localVarReturnValue, nil, reportError("authCode must be less than 6")
	}

	if r.firstAuthAmount == nil {
		return localVarReturnValue, nil, reportError("firstAuthAmount is required and must be specified")
	}
	if *r.firstAuthAmount < 0 {
		return localVarReturnValue, nil, reportError("firstAuthAmount must be greater than 0")
	}

	if r.totalSuccessTimes == nil {
		return localVarReturnValue, nil, reportError("totalSuccessTimes is required and must be specified")
	}
	if *r.totalSuccessTimes < 0 {
		return localVarReturnValue, nil, reportError("totalSuccessTimes must be greater than 0")
	}

	if r.checkMacValue == nil {
		return localVarReturnValue, nil, reportError("checkMacValue is required and must be specified")
	}

	if r.simulatePaid == nil {
		return localVarReturnValue, nil, reportError("simulatePaid is required and must be specified")
	}

	if r.customField1 == nil {
		return localVarReturnValue, nil, reportError("customField1 is required and must be specified")
	}

	if r.customField2 == nil {
		return localVarReturnValue, nil, reportError("customField2 is required and must be specified")
	}

	if r.customField3 == nil {
		return localVarReturnValue, nil, reportError("customField3 is required and must be specified")
	}

	if r.customField4 == nil {
		return localVarReturnValue, nil, reportError("customField4 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("MerchantID", parameterToString(*r.merchantID, ""))
	localVarFormParams.Add("MerchantTradeNo", parameterToString(*r.merchantTradeNo, ""))
	localVarFormParams.Add("StoreID", parameterToString(*r.storeID, ""))
	localVarFormParams.Add("RtnCode", parameterToString(*r.rtnCode, ""))
	localVarFormParams.Add("RtnMsg", parameterToString(*r.rtnMsg, ""))
	localVarFormParams.Add("PeriodType", parameterToString(*r.periodType, ""))
	localVarFormParams.Add("Frequency", parameterToString(*r.frequency, ""))
	localVarFormParams.Add("ExecTimes", parameterToString(*r.execTimes, ""))
	localVarFormParams.Add("Amount", parameterToString(*r.amount, ""))
	localVarFormParams.Add("Gwsr", parameterToString(*r.gwsr, ""))
	localVarFormParams.Add("ProcessDate", parameterToString(*r.processDate, ""))
	localVarFormParams.Add("AuthCode", parameterToString(*r.authCode, ""))
	localVarFormParams.Add("FirstAuthAmount", parameterToString(*r.firstAuthAmount, ""))
	localVarFormParams.Add("TotalSuccessTimes", parameterToString(*r.totalSuccessTimes, ""))
	localVarFormParams.Add("CheckMacValue", parameterToString(*r.checkMacValue, ""))
	localVarFormParams.Add("SimulatePaid", parameterToString(*r.simulatePaid, ""))
	localVarFormParams.Add("CustomField1", parameterToString(*r.customField1, ""))
	localVarFormParams.Add("CustomField2", parameterToString(*r.customField2, ""))
	localVarFormParams.Add("CustomField3", parameterToString(*r.customField3, ""))
	localVarFormParams.Add("CustomField4", parameterToString(*r.customField4, ""))
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"), "")
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiCallbackReturnURLPostRequest struct {
	ctx                  _context.Context
	apiService           *CallBackApiService
	merchantID           *string
	merchantTradeNo      *string
	storeID              *string
	rtnCode              *int
	rtnMsg               *string
	tradeNo              *string
	tradeAmt             *int
	paymentDate          *ECPayDateTime
	paymentType          *ReturnPaymentTypeEnum
	paymentTypeChargeFee *int
	tradeDate            *ECPayDateTime
	checkMacValue        *string
	simulatePaid         *SimulatePaidEnum
	customField1         *string
	customField2         *string
	customField3         *string
	customField4         *string
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

func (r apiCallbackReturnURLPostRequest) RtnCode(rtnCode int) apiCallbackReturnURLPostRequest {
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

func (r apiCallbackReturnURLPostRequest) TradeAmt(tradeAmt int) apiCallbackReturnURLPostRequest {
	r.tradeAmt = &tradeAmt
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentDate(paymentDate ECPayDateTime) apiCallbackReturnURLPostRequest {
	r.paymentDate = &paymentDate
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentType(paymentType ReturnPaymentTypeEnum) apiCallbackReturnURLPostRequest {
	r.paymentType = &paymentType
	return r
}

func (r apiCallbackReturnURLPostRequest) PaymentTypeChargeFee(paymentTypeChargeFee int) apiCallbackReturnURLPostRequest {
	r.paymentTypeChargeFee = &paymentTypeChargeFee
	return r
}

func (r apiCallbackReturnURLPostRequest) TradeDate(tradeDate ECPayDateTime) apiCallbackReturnURLPostRequest {
	r.tradeDate = &tradeDate
	return r
}

func (r apiCallbackReturnURLPostRequest) CheckMacValue(checkMacValue string) apiCallbackReturnURLPostRequest {
	r.checkMacValue = &checkMacValue
	return r
}

func (r apiCallbackReturnURLPostRequest) SimulatePaid(simulatePaid SimulatePaidEnum) apiCallbackReturnURLPostRequest {
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
	if strlen(*r.merchantID) > 10 {
		return localVarReturnValue, nil, reportError("merchantID must have less than 10 elements")
	}

	if r.merchantTradeNo == nil {
		return localVarReturnValue, nil, reportError("merchantTradeNo is required and must be specified")
	}
	if strlen(*r.merchantTradeNo) > 20 {
		return localVarReturnValue, nil, reportError("merchantTradeNo must have less than 20 elements")
	}

	if r.storeID == nil {
		return localVarReturnValue, nil, reportError("storeID is required and must be specified")
	}
	if strlen(*r.storeID) > 20 {
		return localVarReturnValue, nil, reportError("storeID must have less than 20 elements")
	}

	if r.rtnCode == nil {
		return localVarReturnValue, nil, reportError("rtnCode is required and must be specified")
	}

	if r.rtnMsg == nil {
		return localVarReturnValue, nil, reportError("rtnMsg is required and must be specified")
	}
	if strlen(*r.rtnMsg) > 200 {
		return localVarReturnValue, nil, reportError("rtnMsg must have less than 200 elements")
	}

	if r.tradeNo == nil {
		return localVarReturnValue, nil, reportError("tradeNo is required and must be specified")
	}
	if strlen(*r.tradeNo) > 20 {
		return localVarReturnValue, nil, reportError("tradeNo must have less than 20 elements")
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

	if r.customField1 == nil {
		return localVarReturnValue, nil, reportError("customField1 is required and must be specified")
	}

	if r.customField2 == nil {
		return localVarReturnValue, nil, reportError("customField2 is required and must be specified")
	}

	if r.customField3 == nil {
		return localVarReturnValue, nil, reportError("customField3 is required and must be specified")
	}

	if r.customField4 == nil {
		return localVarReturnValue, nil, reportError("customField4 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("MerchantID", parameterToString(*r.merchantID, ""))
	localVarFormParams.Add("MerchantTradeNo", parameterToString(*r.merchantTradeNo, ""))
	localVarFormParams.Add("StoreID", parameterToString(*r.storeID, ""))
	localVarFormParams.Add("RtnCode", parameterToString(*r.rtnCode, ""))
	localVarFormParams.Add("RtnMsg", parameterToString(*r.rtnMsg, ""))
	localVarFormParams.Add("TradeNo", parameterToString(*r.tradeNo, ""))
	localVarFormParams.Add("TradeAmt", parameterToString(*r.tradeAmt, ""))
	localVarFormParams.Add("PaymentDate", parameterToString(*r.paymentDate, ""))
	localVarFormParams.Add("PaymentType", parameterToString(*r.paymentType, ""))
	localVarFormParams.Add("PaymentTypeChargeFee", parameterToString(*r.paymentTypeChargeFee, ""))
	localVarFormParams.Add("TradeDate", parameterToString(*r.tradeDate, ""))
	localVarFormParams.Add("CheckMacValue", parameterToString(*r.checkMacValue, ""))
	localVarFormParams.Add("SimulatePaid", parameterToString(*r.simulatePaid, ""))
	localVarFormParams.Add("CustomField1", parameterToString(*r.customField1, ""))
	localVarFormParams.Add("CustomField2", parameterToString(*r.customField2, ""))
	localVarFormParams.Add("CustomField3", parameterToString(*r.customField3, ""))
	localVarFormParams.Add("CustomField4", parameterToString(*r.customField4, ""))
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"), "")
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
