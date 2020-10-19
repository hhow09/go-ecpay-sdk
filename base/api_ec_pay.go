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

// ECPayApiService ECPayApi service
type ECPayApiService service

type apiQueryCreditCardPeriodInfoRequest struct {
	ctx             _context.Context
	apiService      *ECPayApiService
	merchantID      *string
	merchantTradeNo *string
	timeStamp       *int
	checkMacValue   *string
}

func (r apiQueryCreditCardPeriodInfoRequest) MerchantID(merchantID string) apiQueryCreditCardPeriodInfoRequest {
	r.merchantID = &merchantID
	return r
}

func (r apiQueryCreditCardPeriodInfoRequest) MerchantTradeNo(merchantTradeNo string) apiQueryCreditCardPeriodInfoRequest {
	r.merchantTradeNo = &merchantTradeNo
	return r
}

func (r apiQueryCreditCardPeriodInfoRequest) TimeStamp(timeStamp int) apiQueryCreditCardPeriodInfoRequest {
	r.timeStamp = &timeStamp
	return r
}

func (r apiQueryCreditCardPeriodInfoRequest) CheckMacValue(checkMacValue string) apiQueryCreditCardPeriodInfoRequest {
	r.checkMacValue = &checkMacValue
	return r
}

/*
QueryCreditCardPeriodInfo Method for QueryCreditCardPeriodInfo
13.信用卡定期定額訂單查詢
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiQueryCreditCardPeriodInfoRequest
*/
func (a *ECPayApiService) QueryCreditCardPeriodInfo(ctx _context.Context) apiQueryCreditCardPeriodInfoRequest {
	return apiQueryCreditCardPeriodInfoRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return CreditCardPeriodInfo
*/
func (r apiQueryCreditCardPeriodInfoRequest) Execute() (CreditCardPeriodInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreditCardPeriodInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ECPayApiService.QueryCreditCardPeriodInfo")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Cashier/QueryCreditCardPeriodInfo"

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

	if r.timeStamp == nil {
		return localVarReturnValue, nil, reportError("timeStamp is required and must be specified")
	}
	if *r.timeStamp < 0 {
		return localVarReturnValue, nil, reportError("timeStamp must be greater than 0")
	}

	if r.checkMacValue == nil {
		return localVarReturnValue, nil, reportError("checkMacValue is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("MerchantID", parameterToString(*r.merchantID, ""))
	localVarFormParams.Add("MerchantTradeNo", parameterToString(*r.merchantTradeNo, ""))
	localVarFormParams.Add("TimeStamp", parameterToString(*r.timeStamp, ""))
	localVarFormParams.Add("CheckMacValue", parameterToString(*r.checkMacValue, ""))
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

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"), "application/json")
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiQueryTradeInfoRequest struct {
	ctx             _context.Context
	apiService      *ECPayApiService
	merchantID      *string
	merchantTradeNo *string
	timeStamp       *int
	checkMacValue   *string
	platformID      *string
}

func (r apiQueryTradeInfoRequest) MerchantID(merchantID string) apiQueryTradeInfoRequest {
	r.merchantID = &merchantID
	return r
}

func (r apiQueryTradeInfoRequest) MerchantTradeNo(merchantTradeNo string) apiQueryTradeInfoRequest {
	r.merchantTradeNo = &merchantTradeNo
	return r
}

func (r apiQueryTradeInfoRequest) TimeStamp(timeStamp int) apiQueryTradeInfoRequest {
	r.timeStamp = &timeStamp
	return r
}

func (r apiQueryTradeInfoRequest) CheckMacValue(checkMacValue string) apiQueryTradeInfoRequest {
	r.checkMacValue = &checkMacValue
	return r
}

func (r apiQueryTradeInfoRequest) PlatformID(platformID string) apiQueryTradeInfoRequest {
	r.platformID = &platformID
	return r
}

/*
QueryTradeInfo Method for QueryTradeInfo
8. 查詢訂單
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiQueryTradeInfoRequest
*/
func (a *ECPayApiService) QueryTradeInfo(ctx _context.Context) apiQueryTradeInfoRequest {
	return apiQueryTradeInfoRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return TradeInfo
*/
func (r apiQueryTradeInfoRequest) Execute() (TradeInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TradeInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ECPayApiService.QueryTradeInfo")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Cashier/QueryTradeInfo/V5"

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

	if r.timeStamp == nil {
		return localVarReturnValue, nil, reportError("timeStamp is required and must be specified")
	}
	if *r.timeStamp < 0 {
		return localVarReturnValue, nil, reportError("timeStamp must be greater than 0")
	}

	if r.checkMacValue == nil {
		return localVarReturnValue, nil, reportError("checkMacValue is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("MerchantID", parameterToString(*r.merchantID, ""))
	localVarFormParams.Add("MerchantTradeNo", parameterToString(*r.merchantTradeNo, ""))
	localVarFormParams.Add("TimeStamp", parameterToString(*r.timeStamp, ""))
	if r.platformID != nil {
		localVarFormParams.Add("PlatformID", parameterToString(*r.platformID, ""))
	}
	localVarFormParams.Add("CheckMacValue", parameterToString(*r.checkMacValue, ""))
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
