/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	_context "context"
	"github.com/antihax/optional"
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

// CallbackReturnURLPostOpts Optional parameters for the method 'CallbackReturnURLPost'
type CallbackReturnURLPostOpts struct {
	CustomField1 optional.String
	CustomField2 optional.String
	CustomField3 optional.String
	CustomField4 optional.String
	CustomField5 optional.String
}

/*
CallbackReturnURLPost Method for CallbackReturnURLPost
付款結果通知
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param merchantID 特店編號
 * @param merchantTradeNo 特店交易編號
 * @param storeID 特店旗下店舖代號
 * @param rtnCode 交易狀態
 * @param rtnMsg 交易訊息
 * @param tradeNo 綠界的交易編號
 * @param tradeAmt 交易金額
 * @param paymentDate 付款時間
 * @param paymentType 特店選擇的付款方式
 * @param paymentTypeChargeFee 通路費
 * @param tradeDate 訂單成立時間
 * @param checkMacValue 檢查碼
 * @param simulatePaid 是否為模擬付款
 * @param optional nil or *CallbackReturnURLPostOpts - Optional Parameters:
 * @param "CustomField1" (optional.String) -  自訂名稱欄位 1
 * @param "CustomField2" (optional.String) -  自訂名稱欄位 2
 * @param "CustomField3" (optional.String) -  自訂名稱欄位 3
 * @param "CustomField4" (optional.String) -  自訂名稱欄位 4
 * @param "CustomField5" (optional.String) -  自訂名稱欄位 5
@return string
*/
func (a *CallBackApiService) CallbackReturnURLPost(ctx _context.Context, merchantID string, merchantTradeNo string, storeID string, rtnCode float32, rtnMsg string, tradeNo string, tradeAmt int32, paymentDate string, paymentType PaymentType, paymentTypeChargeFee int32, tradeDate string, checkMacValue string, simulatePaid EcPaySimulateEnum, localVarOptionals *CallbackReturnURLPostOpts) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/callback/ReturnURL"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if tradeAmt < 0 {
		return localVarReturnValue, nil, reportError("tradeAmt must be greater than 0")
	}
	if paymentTypeChargeFee < 0 {
		return localVarReturnValue, nil, reportError("paymentTypeChargeFee must be greater than 0")
	}

	localVarQueryParams.Add("MerchantID", parameterToString(merchantID, ""))
	localVarQueryParams.Add("MerchantTradeNo", parameterToString(merchantTradeNo, ""))
	localVarQueryParams.Add("StoreID", parameterToString(storeID, ""))
	localVarQueryParams.Add("RtnCode", parameterToString(rtnCode, ""))
	localVarQueryParams.Add("RtnMsg", parameterToString(rtnMsg, ""))
	localVarQueryParams.Add("TradeNo", parameterToString(tradeNo, ""))
	localVarQueryParams.Add("TradeAmt", parameterToString(tradeAmt, ""))
	localVarQueryParams.Add("PaymentDate", parameterToString(paymentDate, ""))
	localVarQueryParams.Add("PaymentType", parameterToString(paymentType, ""))
	localVarQueryParams.Add("PaymentTypeChargeFee", parameterToString(paymentTypeChargeFee, ""))
	localVarQueryParams.Add("TradeDate", parameterToString(tradeDate, ""))
	localVarQueryParams.Add("CheckMacValue", parameterToString(checkMacValue, ""))
	localVarQueryParams.Add("SimulatePaid", parameterToString(simulatePaid, ""))
	if localVarOptionals != nil && localVarOptionals.CustomField1.IsSet() {
		localVarQueryParams.Add("CustomField1", parameterToString(localVarOptionals.CustomField1.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomField2.IsSet() {
		localVarQueryParams.Add("CustomField2", parameterToString(localVarOptionals.CustomField2.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomField3.IsSet() {
		localVarQueryParams.Add("CustomField3", parameterToString(localVarOptionals.CustomField3.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomField4.IsSet() {
		localVarQueryParams.Add("CustomField4", parameterToString(localVarOptionals.CustomField4.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CustomField5.IsSet() {
		localVarQueryParams.Add("CustomField5", parameterToString(localVarOptionals.CustomField5.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
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

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
